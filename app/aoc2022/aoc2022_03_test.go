package aoc2022

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func Test_AOC2022_03_Part1(t *testing.T) {
	data := DAY_2022_03_TEST_DATA
	splits := strings.Split(data, "\n")
	rs := NewRucksack(splits[0])
	fmt.Printf("SumCommon: %v\n", rs.SumCommon())
	log.Fatal("fooo")
}

func Test_AOC2022_03_Part11(t *testing.T) {
	data := DAY_2022_03_TEST_DATA
	splits := strings.Split(data, "\n")
	total := 0
	for _, split := range splits {
		rs := NewRucksack(split)
		rs.DebugCommon()
		total += rs.SumCommon()
		// fmt.Printf("SumCommon: %v\n", rs.SumCommon())
		// log.Fatal("fooo")
		// rs.Debug()
	}
	fmt.Printf("Sum is %v\n", total)
	log.Fatal("fooo")
}

func Test_AOC2022_03(t *testing.T) {
	data := DAY_2022_03_DATA
	splits := strings.Split(data, "\n")
	total := 0
	for _, split := range splits {
		rs := NewRucksack(split)
		rs.DebugCommon()
		total += rs.SumCommon()
	}
	fmt.Printf("Sum is %v\n", total)
	log.Fatal("fooo")
}

func Test_AOC2022_03_Part2_Test(t *testing.T) {
	data := DAY_2022_03_TEST_DATA
	splits := strings.Split(data, "\n")
	total := 0
	for index := 0; index < len(splits); index += 3 {
		rs1 := NewRucksack(splits[index])
		rs2 := NewRucksack(splits[index+1])
		rs3 := NewRucksack(splits[index+2])

		f1 := rs1.Frequency()
		f2 := rs2.Frequency()
		f3 := rs3.Frequency()

		fmt.Printf("f1: %v\n", f1)
		fmt.Printf("f2: %v\n", f2)
		fmt.Printf("f3: %v\n", f3)

		for i := 0; i < len(f1); i++ {
			if f1[i] > 0 && f2[i] > 0 && f3[i] > 0 {
				value := (i + 1)
				total += value
				fmt.Printf("group[%v] shares %v, int value %v, total=%v\n", index/3, AZ[i:i+1], value, total)
			}
		}
	}
	fmt.Printf("total is %v\n", total)
	log.Fatal("fooo")
}

func Test_AOC2022_03_Part2_TestX(t *testing.T) {
	data := DAY_2022_03_DATA
	splits := strings.Split(data, "\n")
	total := 0
	for index := 0; index < len(splits); index += 3 {
		rs1 := NewRucksack(splits[index])
		rs2 := NewRucksack(splits[index+1])
		rs3 := NewRucksack(splits[index+2])

		f1 := rs1.Frequency()
		f2 := rs2.Frequency()
		f3 := rs3.Frequency()

		fmt.Printf("f1: %v\n", f1)
		fmt.Printf("f2: %v\n", f2)
		fmt.Printf("f3: %v\n", f3)

		for i := 0; i < len(f1); i++ {
			if f1[i] > 0 && f2[i] > 0 && f3[i] > 0 {
				value := (i + 1)
				total += value
				fmt.Printf("group[%v] shares %v, int value %v, total=%v\n", index/3, AZ[i:i+1], value, total)
			}
		}
	}
	fmt.Printf("total is %v\n", total)
	log.Fatal("fooo")
}
