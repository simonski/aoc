package d10

import (
	"fmt"
	"strconv"
	"strings"
)

type CPU struct {
	RegisterX       int
	SignalStrengths []int
	Device          *Device
	TickNum         int
}

func NewCPU(device *Device) *CPU {
	strengths := make([]int, 0)
	cpu := &CPU{RegisterX: 1, SignalStrengths: strengths, Device: device, TickNum: 0}
	return cpu
}

func (c *CPU) ProcessInstruction(instruction string) {
	splits := strings.Split(instruction, " ")
	if splits[0] == "noop" {
		c.Device.CRT.Tick(c.TickNum)
		c.Tick()
	} else if splits[0] == "addx" {
		// 1. begin adding (tick1)
		// 2. CRT draw pixel
		c.Device.CRT.Tick(c.TickNum)
		c.Tick()
		c.Device.CRT.Tick(c.TickNum)
		c.Tick()
		value, _ := strconv.Atoi(splits[1])
		c.RegisterX += value
	}

}

func (c *CPU) Tick() {
	c.TickNum += 1
	if c.TickNum%20 == 0 {
		c.SignalStrength()
	}
}

func (c *CPU) SignalStrength() {
	value := c.Device.CPU.TickNum * c.RegisterX
	fmt.Printf("SignalStrength() Tick=%v, RegisterX=%v, strength=%v\n", c.Device.CPU.TickNum, c.RegisterX, value)
	c.SignalStrengths = append(c.SignalStrengths, value)
}

func (c *CPU) Debug() {
	total := 0
	for index, value := range c.SignalStrengths {
		total += value
		fmt.Printf("[%v] = [%v], total=%v\n", index, value, total)
	}
}
