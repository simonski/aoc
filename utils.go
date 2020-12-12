package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
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

func load_file_to_ints(filename string) []int {
	file, err := os.Open(filename)
	results := make([]int, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(line)
		results = append(results, value)
	}
	return results

}

func load_file_to_strings(filename string) []string {
	file, err := os.Open(filename)
	results := make([]string, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
	}
	return results

}

func convert_strings_to_ints(input []string) []int {
	output := make([]int, 0)
	for _, value := range input {
		ivalue, _ := strconv.Atoi(value)
		output = append(output, ivalue)
	}
	return output
}

func Min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

type IntMap struct {
	data map[int]int
}

func NewIntMap() *IntMap {
	data := make(map[int]int)
	m := IntMap{data: data}
	return &m
}

func (m *IntMap) Get(key int, defaultValue int) int {
	value, exists := m.data[key]
	if exists {
		return value
	} else {
		return defaultValue
	}
}

func (m *IntMap) Put(key int, value int) {
	m.data[key] = value
}

func (m *IntMap) Increment(key int) int {
	value := m.Get(key, 0)
	value++
	m.Put(key, value)
	return value
}
