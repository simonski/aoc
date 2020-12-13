package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_04 is the entrypoint
func AOC_2020_04(cli *goutils.CLI) {
	AOC_2020_04_part1_attempt1(cli)
	AOC_2020_04_part2_attempt1(cli)
}

func AOC_2020_04_part1_attempt1(cli *goutils.CLI) {
	filename := cli.GetStringOrDie("-input")
	pc := NewPassportControlFromFile(filename)
	validCount := 0
	totalPassports := 0
	for index := 0; index < pc.Size(); index++ {
		passport := pc.Get(index)
		if passport.IsValid() {
			validCount++
		}
		totalPassports++
	}

	fmt.Printf("Part1: Out of %v passports, %v are valid.\n", totalPassports, validCount)
}

func AOC_2020_04_part2_attempt1(cli *goutils.CLI) {
	filename := cli.GetStringOrDie("-input")
	pc := NewPassportControlFromFile(filename)
	validCount := 0
	totalPassports := 0
	for index := 0; index < pc.Size(); index++ {
		passport := pc.Get(index)
		if passport.IsValidPart2() {
			validCount++
		}
		totalPassports++
	}

	fmt.Printf("Part2: Out of %v passports, %v are valid.\n", totalPassports, validCount)
}

type Passport struct {
	lines    []string
	keypairs map[string]string
}

func NewPassport() *Passport {
	p := Passport{}
	p.lines = make([]string, 0)
	p.keypairs = make(map[string]string)
	return &p
}

func (p *Passport) Add(line string) {
	p.lines = append(p.lines, line)
	splits := strings.Split(line, " ")
	for index := 0; index < len(splits); index++ {
		split := splits[index]
		keypair := strings.Split(split, ":")
		key := keypair[0]
		value := keypair[1]
		p.keypairs[key] = value
	}
}

func (p *Passport) Get(key string) (string, bool) {
	val, ok := p.keypairs[key]
	return val, ok
}

func (p *Passport) Debug() {
	fmt.Printf("Password, line=%v\n", p.lines)
	fmt.Printf("key/value")
	fmt.Printf("%v\n\n", p.keypairs)
}

func (p *Passport) Size() int {
	return len(p.keypairs)
}

func (p *Passport) IsValid() bool {
	_, byr_exists := p.Get("byr")
	_, iyr_exists := p.Get("iyr")
	_, eyr_exists := p.Get("eyr")
	_, hgt_exists := p.Get("hgt")
	_, hcl_exists := p.Get("hcl")
	_, ecl_exists := p.Get("ecl")
	_, pid_exists := p.Get("pid")
	// _, cid_exists := p.Get("cid")

	return byr_exists && iyr_exists && eyr_exists && hgt_exists && hcl_exists && ecl_exists && pid_exists // && cid_exists
}

func (p *Passport) IsValidPart2() bool {
	return p.isValidBYR() && p.isValidIYR() && p.isValidEYR() && p.IsValidHGT() && p.IsValidHCL() && p.IsValidECL() && p.IsValidPID()
}

func (p *Passport) isValidBYR() bool {
	byr, byr_exists := p.Get("byr")
	if !byr_exists {
		return false
	}
	ibyr, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	return ibyr >= 1920 && ibyr <= 2002
}

func (p *Passport) isValidIYR() bool {
	iyr, iyr_exists := p.Get("iyr")
	if !iyr_exists {
		return false
	}
	iiyr, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	return iiyr >= 2010 && iiyr <= 2020
}

func (p *Passport) isValidEYR() bool {
	eyr, eyr_exists := p.Get("eyr")
	if !eyr_exists {
		return false
	}
	ieyr, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	return ieyr >= 2020 && ieyr <= 2030
}

func (p *Passport) IsValidHGT() bool {
	//  split by digits and words

	hgt, hgt_exists := p.Get("hgt")
	if !hgt_exists {
		return false
	}

	expression, _ := regexp.Compile(`(\d{2,3})(cm|in)`)
	match := expression.FindStringSubmatch(hgt)
	// fmt.Printf("%v, %v entries, %v\n", hgt, len(match), match)
	if len(match) != 3 {
		return false
	}

	size := match[1]
	units := match[2]

	isize, err := strconv.Atoi(size)
	if err != nil {
		return false
	}

	if units == "cm" {
		return isize >= 150 && isize <= 193
	} else if units == "in" {
		return isize >= 59 && isize <= 76
	} else {
		return false
	}

}

func (p *Passport) IsValidHCL() bool {
	//  split by digits and words

	hcl, hcl_exists := p.Get("hcl")
	if !hcl_exists {
		return false
	}

	expression, _ := regexp.Compile(`\#[0-9a-f]{6}`)
	match := expression.FindStringSubmatch(hcl)
	// fmt.Printf("%v, %v entries, %v\n", hcl, len(match), match)
	if len(match) != 1 {
		return false
	}
	return true

}

func (p *Passport) IsValidECL() bool {
	//  split by digits and words

	ecl, ecl_exists := p.Get("ecl")
	if !ecl_exists {
		return false
	}

	expression, _ := regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	match := expression.FindStringSubmatch(ecl)
	// fmt.Printf("ecl: %v, %v entries, %v\n", ecl, len(match), match)
	if len(match) != 2 {
		return false
	}
	return true

}

func (p *Passport) IsValidPID() bool {
	//  split by digits and words

	pid, pid_exists := p.Get("pid")
	if !pid_exists {
		return false
	}

	expression, _ := regexp.Compile(`^(\d{9})$`)
	match := expression.FindStringSubmatch(pid)
	if len(match) != 2 {
		fmt.Printf("PID BAD %v, %v entries, %v\n", pid, len(match), match)
		return false
	}
	fmt.Printf("PID GOOD %v, %v entries, %v\n", pid, len(match), match)
	return true

}

// TobogganMap represents the whole snow field
type PassportControl struct {
	passports []*Passport
}

func (pc *PassportControl) Size() int {
	return len(pc.passports)
}

func (pc *PassportControl) Get(index int) *Passport {
	return pc.passports[index]
}

// NewPassportControlFromString constructs and populates a snowy field from the passed filename
func NewPassportControlFromFile(filename string) *PassportControl {
	lines := load_file_to_strings(filename)
	return NewPassportControlFromString(lines)
}

func NewPassportControlFromString(lines []string) *PassportControl {
	pc := PassportControl{}
	pc.passports = make([]*Passport, 0)
	passport := NewPassport()
	pc.passports = append(pc.passports, passport)
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		if line == "\n" || line == "" {
			passport = NewPassport()
			pc.passports = append(pc.passports, passport)
		} else {
			passport.Add(line)
		}
	}
	return &pc
}
