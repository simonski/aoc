package main

import (
	"strings"
	"testing"
)

const TEST_STRING_04 = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func Test_AOC2020_04_PassportControl(t *testing.T) {

	passports := NewPassportControlFromString(strings.Split(TEST_STRING_04, "\n"))

	if passports.Size() != 4 {
		t.Errorf("Should have loaded 4 passports, instead loaded %v.\n", passports.Size())
	}

	p1 := passports.Get(0)
	p2 := passports.Get(1)
	p3 := passports.Get(2)
	p4 := passports.Get(3)

	p1.Debug()
	p2.Debug()
	p3.Debug()
	p4.Debug()

	verifyPassportField(p1, "ecl", "gry", t)
	verifyPassportField(p1, "pid", "860033327", t)
	verifyPassportField(p1, "eyr", "2020", t)
	verifyPassportField(p1, "hcl", "#fffffd", t)
	verifyPassportField(p1, "byr", "1937", t)
	verifyPassportField(p1, "iyr", "2017", t)
	verifyPassportField(p1, "cid", "147", t)
	verifyPassportField(p1, "hgt", "183cm", t)

	if !p1.IsValid() {
		t.Errorf("Passport1 should be valid.")
	}

	if p2.IsValid() {
		t.Errorf("Passport2 is not valid.")
	}

	if !p3.IsValid() {
		t.Errorf("Passport3 should be valid.")
	}

	if p4.IsValid() {
		t.Errorf("Passport4 is not valid.")
	}

	if p1.Size() != 8 {
		t.Errorf("Passport1 key length is invalid should be 8, is %v\n", p1.Size())
	}
	if p2.Size() != 7 {
		t.Errorf("Passport2 key length is invalid should be 7, is %v\n", p2.Size())
	}
	if p3.Size() != 7 {
		t.Errorf("Passport3 key length is invalid should be 7, is %v\n", p3.Size())
	}
	if p4.Size() != 6 {
		t.Errorf("Passport4 key length is invalid should be 6, is %v\n", p4.Size())
	}

}

func Test_AOC2020_04_IsValidHGT(t *testing.T) {

	h1 := NewPassport()
	h1.Add("hgt:34m")
	if h1.IsValidHGT() {
		t.Errorf("%v is not a valid height.", h1)
	}

	h2 := NewPassport()
	h2.Add("hgt:34cm")
	if h2.IsValidHGT() {
		t.Errorf("%v is not a valid height.", h2)
	}

	h3 := NewPassport()
	h3.Add("hgt:150cm")
	if !h3.IsValidHGT() {
		t.Errorf("%v is  a valid height.", h3)
	}

	h4 := NewPassport()
	h4.Add("hgt:149cm")
	if h4.IsValidHGT() {
		t.Errorf("%v is not a valid height.", h4)
	}

	h5 := NewPassport()
	h5.Add("hgt:193cm")
	if !h5.IsValidHGT() {
		t.Errorf("%v is a valid height.", h5)
	}

	h6 := NewPassport()
	h6.Add("hgt:194cm")
	if h6.IsValidHGT() {
		t.Errorf("%v is not a valid height.", h6)
	}

}

func verifyPassportField(passport *Passport, key string, expected string, t *testing.T) {
	value, _ := passport.Get(key)
	if expected != value {
		t.Errorf("%v=%v != %v", key, expected, value)
	}
}

func Test_AOC2020_04_IsValidHCL(t *testing.T) {

	h1 := NewPassport()
	h1.Add("hcl:sdfdfdf")
	if h1.IsValidHCL() {
		t.Errorf("%v is not a valid hcl.", h1)
	}

	h2 := NewPassport()
	h2.Add("hcl:123123")
	if h2.IsValidHCL() {
		t.Errorf("%v is not a valid hcl.", h2)
	}

	h3 := NewPassport()
	h3.Add("hcl:#123123")
	if !h3.IsValidHCL() {
		t.Errorf("%v is valid hcl.", h3)
	}

	h4 := NewPassport()
	h4.Add("hcl:fffaaa")
	if h4.IsValidHCL() {
		t.Errorf("%v is not a valid hcl.", h4)
	}

	h5 := NewPassport()
	h5.Add("hcl:#fffaaa")
	if !h5.IsValidHCL() {
		t.Errorf("%v is valid hcl.", h5)
	}

	h6 := NewPassport()
	h6.Add("hcl:#fffaax")
	if h6.IsValidHCL() {
		t.Errorf("%v is not a valid hcl.", h6)
	}

}
