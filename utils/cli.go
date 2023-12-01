package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
CLI is my helper utility to work out what the user did type
*/
type CLI struct {
	Args          []string
	IsInteractive bool
	IsVerbose     bool
	IsExitOnError bool
}

/*
NewCLI create instance of a cli
*/
func NewCLI(args []string) *CLI {
	c := CLI{Args: args}
	c.IsVerbose = c.Contains("-v")
	return &c
}

func NewCLIFromString(line string) *CLI {
	splits := strings.Split(line, " ")
	cli := NewCLI(splits)
	return cli
}

func (c *CLI) GetCommand() string {
	if len(c.Args) > 0 {
		return c.Args[0]
	} else {
		return ""
	}
}

func (c *CLI) Shift() {
	if len(c.Args) > 0 {
		c.Args = c.Args[1:]
	}
}

/*
IndexOf find the position (or -1 if not present) in the args
*/
func (c CLI) IndexOf(key string) int {
	for index := 0; index < len(c.Args); index++ {
		if c.Args[index] == key {
			return index
		}
	}
	return -1
}

// Contains indicates if a key exists or
func (c CLI) Contains(key string) bool {
	return c.IndexOf(key) > -1
}

/*
SplitStringToInts splits the cols string based on the delimiter, converting the results in an []int
*/
func (c CLI) SplitStringToInts(cols string, delim string) []int {
	columns := strings.Split(cols, delim)
	result := make([]int, len(columns))
	for index := 0; index < len(columns); index++ {
		strValue := columns[index]
		intValue, _ := strconv.Atoi(strValue)
		result[index] = intValue
	}
	return result
}

/*
SplitStringToFloats splits the cols string based on the delimiter, converting the results in an []float64
*/
func (c CLI) SplitStringToFloats(cols string, delim string) []float64 {
	columns := strings.Split(cols, delim)
	result := make([]float64, len(columns))
	for index := 0; index < len(columns); index++ {
		strValue := columns[index]
		intValue, _ := strconv.ParseFloat(strValue, 64)
		result[index] = intValue
	}
	return result
}

/*
GetStringOrDie requires the key exist in the CLI arguments or os.Exit(1)
*/
func (c CLI) GetStringOrDie(key string) string {
	index := c.IndexOf(key)
	if index == -1 {
		fmt.Printf("Fatal: '%s' is required.\n", key)
		os.Exit(1)
		return ""
	} else {
		if index+1 < len(c.Args) {
			testValue := c.Args[index+1]
			if testValue[0:1] == "-" {
				// Then there is no value - the key is specified without a value.
				// In this case if the user wants to know if a key ahs been
				// reuqested they would call .ContainsKey or .IndexOf
				fmt.Printf("Fatal: '%s' requires a value.\n", key)
				os.Exit(1)
				return ""
			}
			return testValue
		}
		return ""
	}
}

/*
GetStringOrDefault returns the value associated with the key or the defaultValue if not present
*/
func (c CLI) GetStringOrDefault(key string, defaultValue string) string {
	index := c.IndexOf(key)
	if index == -1 {
		return defaultValue
	}

	if index+1 < len(c.Args) {
		testValue := c.Args[index+1]
		if testValue[0:1] == "-" {
			// then there is no value
			return defaultValue
		}
		return testValue
	}
	return defaultValue
}

/*
GetIntOrDie returns the int value associated with the passed key or fails and os.Exit(1)
*/
func (c CLI) GetIntOrDie(key string) int {
	value := c.GetStringOrDie(key)
	v, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Fatal: '%s' should be an integer.\n", key)
		os.Exit(1)
		return -1
	}
	return v
}

/*
GetIntOrDefault returns the int asssociated with the key or the defaultValue
*/
func (c CLI) GetIntOrDefault(key string, defaultValue int) int {
	strDefaultValue := strconv.Itoa(defaultValue)
	value := c.GetStringOrDefault(key, strDefaultValue)
	v, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Fatal: '%s' should be an integer.\n", key)
		os.Exit(1)
		return -1
	}
	return v
}

func (c CLI) GetStringOrEnvOrDefault(key string, envKey string, defaultValue string) string {
	envValue := os.Getenv(envKey)
	value := defaultValue
	if envValue != "" {
		value = envValue
	}
	return c.GetStringOrDefault(key, value)
}

func (c CLI) GetStringOrEnvOrDie(key string, envKey string) string {
	value := os.Getenv(envKey)
	if value == "" {
		value = c.GetStringOrDefault(key, "")
	}
	if value == "" {
		fmt.Printf("Error, '%v' or '$%v' are required.\n", key, envKey)
		os.Exit(1)
	}
	return value
}

func (c CLI) GetIntOrEnvOrDefault(key string, envKey string, defaultValue int) int {
	svalue := fmt.Sprintf("%v", defaultValue)
	value := c.GetStringOrEnvOrDefault(key, envKey, svalue)
	ivalue, err := strconv.Atoi(value)
	if err == nil {
		return ivalue
	} else {
		return defaultValue
	}
}

/*
GetFileExistsOrDie returns the name of the file provided if it exists or failes and os.Exit(1)
*/
func (c CLI) GetFileExistsOrDie(key string) string {
	message := fmt.Sprintf("Fatal: '%s' does not have a value.\n", key)
	return c.GetFileExistsOrDieWithMessage(key, message)
}

func (c CLI) GetFileExistsOrDieWithMessage(key string, message string) string {
	filename := c.GetStringOrDie(key)
	if filename == "" {
		fmt.Println(message)
		os.Exit(1)
		return ""
	}

	if c.FileExists(filename) {
		return filename
	} else {
		fmt.Printf("Fatal: '%s' does not exist.\n", filename)
		os.Exit(1)
		return ""
	}
}

/*
GetFileExistsOrDefault returns the filename associated with the key or returns the defaultValue if the file does not exist
*/
func (c CLI) GetFileExistsOrDefault(key string, defaultValue string) string {
	filename := c.GetStringOrDefault(key, defaultValue)
	if filename == "" {
		return defaultValue
	}
	return filename
}

/*
FileExists asserts the filename exists and is a file
*/
func (c CLI) FileExists(filename string) bool {
	result, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !result.IsDir()
}

// GetEnvOrDefault returns an os.Getenv value or the defaultValue
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	} else {
		return defaultValue
	}
}

func (c CLI) GetStringFromSetOrDefault(key string, defaultValue string, permitted []string) string {
	value := c.GetStringOrDefault(key, defaultValue)
	for _, entry := range permitted {
		if value == entry {
			return value
		}
	}
	return defaultValue
}

func (c CLI) GetStringFromSetOrDie(key string, permitted []string) string {
	value := c.GetStringOrDefault(key, "")
	for _, entry := range permitted {
		if value == entry {
			return value
		}
	}
	fmt.Printf("Fatal: '%s' does not exist in set %v .\n", key, permitted)
	os.Exit(1)
	return ""
}

func (c CLI) Flatten() string {
	output := ""
	for _, value := range c.Args {
		output += value
		output += " "
	}
	return strings.TrimSpace(output)
}
