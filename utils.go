package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func BuildScanner(filename *string) *FileScanner {
	if strings.HasSuffix(*filename, ".gz") {
		file, err := os.OpenFile(*filename, os.O_RDONLY, os.ModePerm)
		Check(err)
		gz, _ := gzip.NewReader(file)
		scanner := bufio.NewScanner(gz)
		return &FileScanner{file, scanner}
	} else {
		file, err := os.OpenFile(*filename, os.O_RDONLY, os.ModePerm)
		Check(err)
		scanner := bufio.NewScanner(file)
		return &FileScanner{file, scanner}
	}

}

func Console(msg ...string) {
	if len(msg) == 2 {
		fmt.Printf("%-30v%s\n", msg[0], msg[1])
	} else {
		fmt.Println(msg[0])
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var letterRunes = []rune(" -_.;:/1234567890)(*&^%$£abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
