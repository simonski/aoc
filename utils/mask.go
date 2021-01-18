package utils

import "fmt"

type Mask struct {
	Data string
}

func (b *Mask) Get(index int) string {
	if len(b.Data) > index {
		position := len(b.Data) - index
		return b.Data[position-1 : position]
	} else {
		return "X"
	}
}

// DeriveNewMask applies this mask onto a value resulting in a new mask value
// a new mask and an int[] of the bit indexes changed in the new mask

//
//address: 000000000000000000000000000000101010  (decimal 42)
//mask:    000000000000000000000000000000X1001X
//result:  000000000000000000000000000000X1101X
func (m *Mask) DeriveNewMask(address string) *Mask {
	// changedBits := make([]int, 0)
	newMaskValue := ""
	for index := 0; index < len(m.Data); index++ {
		maskValue := m.Data[index : index+1]
		addressValue := address[index : index+1]
		if maskValue == addressValue || maskValue == "0" {
			// ignore
			newMaskValue += addressValue
		} else {
			// fmt.Printf("Mask.DeriveNewMask: address [%v] value is %v, changing to mask value of %v\n", index, addressValue, maskValue)
			addressValue = maskValue
			// changedBits = append(changedBits, index)
			newMaskValue += maskValue
		}
	}

	newMask := NewMask(newMaskValue)
	return newMask //, changedBits
}

func NewMask(value string) *Mask {
	m := Mask{Data: value}
	return &m
}

// func (b *Mask) BuildSubMasks() []*Mask {

// }

func (m *Mask) GetVariations() []string {
	results := make([]string, 0)
	position := 0
	maskData := m.Data
	results = MaskPermutations(maskData, position, results)
	return results
}

func MaskPermutations(maskData string, position int, results []string) []string {
	maskValue := maskData[position : position+1]
	if maskValue == "X" {
		fmt.Printf("MaskPermutations X [%v]=%v %v\n", position, maskValue, maskData)
		leftPart := maskData[0:position]
		rightPart := maskData[position+1:]

		fmt.Printf("left: %v, middle X, right %v\n", leftPart, rightPart)

		mask0 := fmt.Sprintf("%v0%v", leftPart, rightPart)
		mask1 := fmt.Sprintf("%v1%v", leftPart, rightPart)
		fmt.Printf("\tMask0: %v\n", mask0)
		fmt.Printf("\tMask1: %v\n", mask1)

		// if position == len(maskData)-1 {
		// 	results = append(results, mask0)
		// 	results = append(results, mask1)
		// 	return results
		// } else {
		results = MaskPermutations(mask0, position, results)
		results = MaskPermutations(mask1, position, results)
		return results
		// }
	} else if position < len(maskData)-1 {
		// fmt.Printf("MaskPermutations [%v]=%v %v, walk on +1\n", position, maskValue, maskData)
		return MaskPermutations(maskData, position+1, results)
	} else {
		fmt.Printf("MaskPermutations [%v]=%v %v, complete\n", position, maskValue, maskData)
		results = append(results, maskData)
		return results
	}
}
