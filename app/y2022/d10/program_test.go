package d10

import (
	"fmt"
	"log"
	"testing"
)

func verifyCPU(c *CPU, instruction string, expectRegisterX int, expectTickNum int, t *testing.T) {
	if c.Device.CPU.TickNum != expectTickNum {
		log.Fatalf("(%v) expect tick %v was %v\n", instruction, expectTickNum, c.Device.CPU.TickNum)
	}
	if c.RegisterX != expectRegisterX {
		log.Fatalf("(%v) expect X %v was %v\n", instruction, expectRegisterX, c.RegisterX)
	}
}

func Test_1(t *testing.T) {
	c := NewDevice(TEST_DATA_01, 6, 40)

	// noop
	i := c.Instructions[0]
	c.ProcessInstruction(i)
	verifyCPU(c.CPU, i, 1, 1, t)

	// addx 3
	i = c.Instructions[1]
	c.ProcessInstruction(i)
	verifyCPU(c.CPU, i, 4, 3, t)

	// addx -5
	i = c.Instructions[2]
	c.ProcessInstruction(i)
	verifyCPU(c.CPU, i, -1, 5, t)

	c.CPU.Debug()
	t.Fatalf("mm")
}

func Test_2(t *testing.T) {
	c := NewDevice(TEST_DATA_02, 6, 40)
	c.ProcessInstructions()
	c.CPU.Debug()

	total := 0
	total += c.CPU.SignalStrengths[20/20-1]
	total += c.CPU.SignalStrengths[60/20-1]
	total += c.CPU.SignalStrengths[100/20-1]
	total += c.CPU.SignalStrengths[140/20-1]
	total += c.CPU.SignalStrengths[180/20-1]
	total += c.CPU.SignalStrengths[220/20-1]

	if total != 13140 {
		log.Fatalf("Expected total %v, got %v\n", 13140, total)
	}
	t.Fatalf("mm")
}

func Test_Part(t *testing.T) {
	c := NewDevice(REAL_DATA, 6, 40)
	c.ProcessInstructions()
	c.CPU.Debug()

	total := 0
	total += c.CPU.SignalStrengths[20/20-1]
	total += c.CPU.SignalStrengths[60/20-1]
	total += c.CPU.SignalStrengths[100/20-1]
	total += c.CPU.SignalStrengths[140/20-1]
	total += c.CPU.SignalStrengths[180/20-1]
	total += c.CPU.SignalStrengths[220/20-1]

	if total != 13140 {
		log.Fatalf("Expected total %v, got %v\n", 13140, total)
	}
	t.Fatalf("mm")
}

func Test_Part2_1(t *testing.T) {
	d := NewDevice(TEST_DATA_02, 6, 40)
	d.DEBUG = true
	// fmt.Printf("Tick=%v, RegisterX=%v, \n", d.CPU.TickNum, d.CPU.RegisterX)
	// fmt.Println("Row: \n" + d.CRT.DrawSprite() + "\n")
	// d.DEBUG = true
	for _, instruction := range d.Instructions {
		d.ProcessInstruction(instruction)
		fmt.Println("Row: \n" + d.CRT.DrawSprite() + "\n")
		fmt.Printf("Tick=%v, X=%v after instruction %v\n", d.CPU.TickNum, d.CPU.RegisterX, instruction)
		fmt.Println(d.CRT.Draw())
	}
}

func Test_Part2_Real(t *testing.T) {
	d := NewDevice(REAL_DATA, 6, 40)
	fmt.Printf("Tick=%v, RegisterX=%v, \n", d.CPU.TickNum, d.CPU.RegisterX)
	fmt.Println("Row: \n" + d.CRT.DrawSprite() + "\n")
	fmt.Println(d.CRT.Draw())
	d.DEBUG = true
	for _, instruction := range d.Instructions {
		d.ProcessInstruction(instruction)
		fmt.Println("Row: \n" + d.CRT.DrawSprite() + "\n")
		fmt.Printf("Tick=%v, X=%v after instruction %v\n", d.CPU.TickNum, d.CPU.RegisterX, instruction)
		fmt.Println(d.CRT.Draw())
	}

	// fmt.Println("")
	// fmt.Println(d.CRT.Draw())
	// d.Tick()
	// fmt.Println(d.CRT.Draw())
	// fmt.Println("")
	// t.Fatal("mm")
}
