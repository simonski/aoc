package d18

import (
	"fmt"
)

func Key(x int, y int, z int) string {
	return fmt.Sprintf("%v,%v,%v", x, y, z)
}
