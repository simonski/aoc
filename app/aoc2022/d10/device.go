package d10

import "strings"

type Device struct {
	CPU          *CPU
	CRT          *CRT
	Instructions []string
	DEBUG        bool
}

func NewDevice(input string, rows int, cols int) *Device {
	instructions := strings.Split(input, "\n")
	device := &Device{Instructions: instructions}
	device.CPU = NewCPU(device)
	device.CRT = NewCRT(device, rows, cols)
	return device
}

func (d *Device) ProcessInstructions() {
	for _, instruction := range d.Instructions {
		d.CPU.ProcessInstruction(instruction)
	}
}

func (d *Device) ProcessInstruction(instruction string) {
	d.CPU.ProcessInstruction(instruction)
}
