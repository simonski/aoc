package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"math"
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

// make_map_of_inty_list helper makes a map[int]int of a []int to give me
// whatever go's maps key algorithm performance is, at the cost of the memory
func make_map_of_inty_list(data []int) map[int]int {
	m := make(map[int]int)
	for index := 0; index < len(data); index++ {
		value := data[index]
		m[value] = value
	}
	return m
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

func decimal_to_binary(value int64) string {
	b := NewBitSet(int64(value))
	return b.ToBinaryString(36)
}

func binary_to_decimal(decimalValue string) int64 {
	total := int64(0)
	for index := 0; index < len(decimalValue); index++ {
		value := decimalValue[index : index+1]
		power := len(decimalValue) - index - 1
		if value == "1" {
			total += int64(math.Pow(2, float64(power)))
		}
	}
	return total
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
