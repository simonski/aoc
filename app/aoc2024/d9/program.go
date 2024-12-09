package d9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/aoc/utils"
)

/*
Day 9: Disk Fragmenter
*/

type Puzzle struct {
	title string
	year  int
	day   int
	input string
	lines []string
}

func (puzzle *Puzzle) GetSummary() utils.Summary {
	s := utils.Summary{Day: puzzle.day, Year: puzzle.year, Name: puzzle.title}
	s.ProgressP1 = utils.NotStarted
	s.ProgressP2 = utils.NotStarted
	s.DateStarted = "2024-12-09 13:28:57"
	return s
}

func NewPuzzleWithData(input string) *Puzzle {
	iyear, _ := strconv.Atoi("2024")
	iday, _ := strconv.Atoi("9")
	p := Puzzle{year: iyear, day: iday, title: "Day 9: Disk Fragmenter"}
	p.Load(input)
	return &p
}

type Block struct {
	id   int
	size int
}

func NewBlock(id int) *Block {
	return &Block{id: id}
}

type Disk struct {
	blocks    []*Block
	last_swap int
}

func NewDisk(data string) *Disk {
	id := 0
	blocks := make([]*Block, 0)
	for index := 0; index < len(data); index++ {
		value_str := data[index : index+1]
		value, _ := strconv.Atoi(value_str)
		if index%2 == 0 {
			for count := 0; count < value; count++ {
				block := NewBlock(id)
				blocks = append(blocks, block)
			}
			id += 1
		} else {
			for count := 0; count < value; count++ {
				block := NewBlock(-1)
				blocks = append(blocks, block)
			}
		}
	}
	disk := Disk{blocks: blocks, last_swap: -999}
	return &disk
}

func (disk *Disk) Defrag(debug bool) {
	if debug {
		disk.debug()
	}
	for {
		result := disk.step()
		if debug && !result {
			disk.debug()
		}
		if result {
			break
		}
	}
}

func (disk *Disk) checksum() int {
	result := 0
	for index := 0; index < len(disk.blocks); index++ {
		block := disk.blocks[index]
		if block.id > -1 {
			value := index * block.id
			result += value
		}
	}
	return result
}

func (disk *Disk) get_last_block() (int, *Block) {
	for index := len(disk.blocks) - 1; index >= 0; index-- {
		block := disk.blocks[index]
		if block.id > -1 {
			return index, block
		}
	}
	return -1, nil
}

func (disk *Disk) get_first_space() (int, *Block) {
	for index := 0; index < len(disk.blocks); index++ {
		block := disk.blocks[index]
		if block.id == -1 {
			return index, block
		}
	}
	return -1, nil
}

func (disk *Disk) swap(index1 int, index2 int) {
	b1 := disk.blocks[index1]
	b2 := disk.blocks[index2]
	id := b1.id
	size := b1.size

	b1.id = b2.id
	b1.size = b2.size
	b2.id = id
	b2.size = size
}

func (disk *Disk) step() bool {
	block_index, _ := disk.get_last_block()
	space_index, _ := disk.get_first_space()
	if block_index < space_index {
		return true
	}
	if space_index == disk.last_swap {
		return true
	}
	disk.last_swap = space_index
	disk.swap(block_index, space_index)
	return false
}

func (disk *Disk) debug() {
	for _, b := range disk.blocks {
		id_str := fmt.Sprintf("%v", b.id)
		if b.id > -1 {
			fmt.Print(id_str)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (puzzle *Puzzle) p1(data string, debug bool) int {
	fmt.Println()
	d := NewDisk(data)
	d.Defrag(debug)
	return d.checksum()
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
}

func (puzzle *Puzzle) Part1() {
	answer1 := puzzle.p1(TEST_DATA_1, true)
	answer2 := puzzle.p1(TEST_DATA_2, true)
	answer3 := puzzle.p1(REAL_DATA, false)

	fmt.Printf("a1: %v\na2: %v\na3: %v\n", answer1, answer2, answer3)
	// puzzle.p1(REAL_DATA)
}

func (puzzle *Puzzle) Part2() {
	// puzzle.p2(TEST_DATA)
	// puzzle.p2(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
