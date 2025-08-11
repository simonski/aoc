package d10

import "fmt"

type CRT struct {
	Pixels map[string]string
	Rows   int
	Cols   int
	Device *Device
}

func NewCRT(device *Device, rows int, cols int) *CRT {
	pixels := make(map[string]string)
	c := CRT{Pixels: pixels, Rows: rows, Cols: cols, Device: device}
	return &c
}

func (c *CRT) Tick(tickNum int) {
	row := (tickNum + 1) / c.Cols
	colToWrite := tickNum - (row * c.Cols)
	value := "."
	if c.Device.CPU.RegisterX >= colToWrite-1 && c.Device.CPU.RegisterX <= colToWrite+1 {
		value = "#"
	}
	if c.Device.DEBUG {
		fmt.Printf("CRT: Tick[%v], X=%v, row=%v, col=%v, value=%v\n", tickNum, c.Device.CPU.RegisterX, row, colToWrite, value)
	}

	c.Put(value, row, colToWrite)
}

func (c *CRT) Get(row int, col int) string {
	key := fmt.Sprintf("%v_%v", row, col)
	result := c.Pixels[key]
	if result == "" {
		return "X"
	}
	return result
}

func (c *CRT) Put(value string, row int, col int) {
	key := fmt.Sprintf("%v_%v", row, col)
	c.Pixels[key] = value
}

func (c *CRT) DrawSprite() string {
	left := ""
	for index := 0; index < c.Device.CPU.RegisterX-1; index++ {
		left += "."
	}
	sprite := "###"

	right := ""
	for index := c.Device.CPU.RegisterX - 1 + 3; index < c.Cols; index++ {
		right += "."
	}
	return left + sprite + right
}

func (c *CRT) Draw() string {
	result := ""
	// for row := c.Rows - 1; row >= 0; row-- {
	for row := 0; row < c.Rows; row++ {
		line := ""
		for col := 0; col < c.Cols; col++ {
			value := c.Get(row, col)
			line = fmt.Sprintf("%v%v", line, value)
		}
		line = fmt.Sprintf("%v\n", line)
		result = fmt.Sprintf("%v%v", result, line)
	}
	return result
}
