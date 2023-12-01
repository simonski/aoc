package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := NewCLIFromString("run -file fred.txt c d -e 'hello'")
	fmt.Printf("CLI contains %d\n", c.GetIntOrDefault("-fooo", 4))
}

func Test_NoCommandButOptions(t *testing.T) {
	line := "program -option1 -option2"
	c := NewCLIFromString(line)
	command := c.GetCommand()
	if command != "program" {
		t.Errorf("The line is '%v', the command is '%v' (it should be 'program')", line, command)
	}
}

func Test_NoCommand(t *testing.T) {
	line := "program"
	c := NewCLIFromString(line)
	command := c.GetCommand()
	if command != "program" {
		t.Errorf("The line is '%v', the command is '%v' (it should be 'program')", line, command)
	}
}

func Test_CommandAndOptions(t *testing.T) {
	line := "mycommand -option1 -option2 value1"
	c := NewCLIFromString(line)
	command := c.GetCommand()
	if command != "mycommand" {
		t.Errorf("The commmand should be mycommand, it is '%v'", command)
	}
}

func Test_IndexOf1(t *testing.T) {
	line := "program -option1 -option2 value2"
	c := NewCLIFromString(line)
	if c.IndexOf("-option1") != 1 {
		t.Errorf("line is '%v', indexOf(-option1) should be 0, was %v\n", line, c.IndexOf("-option1"))
	}
	if c.IndexOf("-option2") != 2 {
		t.Errorf("line is '%v', indexOf(-option2) shoudl be 0, was %v\n", line, c.IndexOf("-option2"))
	}

	option1value := c.GetStringOrDefault("-option1", "default")
	option2value := c.GetStringOrDefault("-option2", "default")

	if option1value != "default" {
		t.Errorf("option1 should be empty but instead it is '%v'\n", option1value)
	}

	if option2value != "value2" {
		t.Errorf("option2 should be 'value2' but instead it is '%v'\n", option2value)
	}

}

func Test_GetStringFromSetOrDefault(t *testing.T) {
	line := "program mycommand -foo bar"
	c := NewCLIFromString(line)
	value := c.GetStringOrDefault("-foo", "default")
	if value != "bar" {
		t.Errorf("The value should be 'bar', it is '%v'", value)
	}

	value = c.GetStringFromSetOrDefault("-foo", "default", []string{"a", "b", "c"})
	if value != "default" {
		t.Errorf("The value should be 'bar', it is '%v'", value)
	}

}

func Test_GetStringFromSetOrDie(t *testing.T) {
	line := "program mycommand -foo a"
	c := NewCLIFromString(line)
	c.GetStringFromSetOrDie("-foo", []string{"a", "b", "c"})
}

func Test_GetIntOrEnvOrDefault(t *testing.T) {
	line := "-mykey 10"
	c := NewCLIFromString(line)
	value := c.GetIntOrDefault("-mykey", 6)
	if value != 10 {
		t.Errorf("Should be 10")
	}

	os.Setenv("MY_KEY", "")
	value2 := c.GetIntOrEnvOrDefault("-myotherkey", "MY_KEY", 10)
	if value2 != 10 {
		t.Errorf("Should be 10")
	}

	os.Setenv("MY_KEY", "50")
	value3 := c.GetIntOrEnvOrDefault("-myotherkey", "MY_KEY", 10)
	if value3 != 50 {
		t.Errorf("Should be 50, is %v\n", value3)
	}

}

func Test_Shift(t *testing.T) {
	line := "10 9 8 7 6 5 4 3 2 1"
	c := NewCLIFromString(line)
	if c.GetCommand() != "10" {
		t.Errorf("Command should be ten.")
	}

	if c.Flatten() != "10 9 8 7 6 5 4 3 2 1" {
		t.Errorf("Flatten after 0 shift should be '10 9 8 7 6 5 4 3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	if c.Flatten() != "9 8 7 6 5 4 3 2 1" {
		t.Errorf("Flatten after 1 shift should be '9 8 7 6 5 4 3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	if c.Flatten() != "3 2 1" {
		t.Errorf("Flatten after 7 shifts should be '3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	c.Shift()
	c.Shift()
	if c.Flatten() != "" {
		t.Errorf("Flatten after 10 shifts should be '' but was '%v'", c.Flatten())
	}

	c.Shift()
	if c.Flatten() != "" {
		t.Errorf("Flatten after 10 shifts should be '' but was '%v'", c.Flatten())
	}

}
