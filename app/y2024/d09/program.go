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

func NewDiskP1(data string) *Disk {
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

func NewDiskP2(data string) *Disk {
	id := 0
	blocks := make([]*Block, 0)
	for index := 0; index < len(data); index++ {
		value_str := data[index : index+1]
		value, _ := strconv.Atoi(value_str)
		if index%2 == 0 {
			block := NewBlock(id)
			block.size = value
			blocks = append(blocks, block)
			id += 1
		} else {
			block := NewBlock(-1)
			block.size = value
			blocks = append(blocks, block)
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

func (disk *Disk) DefragP2(debug bool) {

	if debug {
		disk.debugP2()
	}

	max_file_id := 0
	for _, block := range disk.blocks {
		max_file_id = utils.MaxInt(max_file_id, block.id)
	}
	for file_id := max_file_id; file_id >= 0; file_id-- {
		disk.stepP2(file_id)
		if debug {
			disk.debugP2()
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

func (disk *Disk) checksumP2() int {
	result := 0
	index := 0
	for block_index := 0; block_index < len(disk.blocks); block_index++ {
		block := disk.blocks[block_index]
		if block.id > -1 {
			for i := 0; i < block.size; i++ {
				value := index * block.id
				result += value
				index += 1
			}
		} else {
			for i := 0; i < block.size; i++ {
				index += 1
			}

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

func (disk *Disk) get_first_available_space(min_size int) (int, *Block) {
	for index := 0; index < len(disk.blocks); index++ {
		block := disk.blocks[index]
		if block.id == -1 && block.size >= min_size {
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

func (disk *Disk) get_block(file_id int) (int, *Block) {
	for block_index, block := range disk.blocks {
		if block.id == file_id {
			return block_index, block
		}
	}
	return -1, nil
}

func (disk *Disk) stepP2(file_id int) bool {
	// fmt.Printf("stepP2(file_id=%v)\n", file_id)
	block_index, block := disk.get_block(file_id)
	space_index, space := disk.get_first_available_space(block.size)
	if space_index == -1 {
		return false
	} else if space_index > block_index {
		// fmt.Printf("step(file_id=%v, block_index=%v, block_size=%v, no space available)\n", file_id, block_index, block.size)
		return false
	}
	// fmt.Printf("step(file_id=%v, block_index=%v, block_size=%v, space_index=%v, space_size=%v)\n", file_id, block_index, block.size, space_index, space.size)
	if space.size == block.size {
		// matches exactly
		// fmt.Printf("stepP2(%v) perfect match()\n", file_id)
		disk.swap(block_index, space_index)
		return true
	} else if space.size < block.size {
		// too big, move on
		// fmt.Printf("stepP2(%v) no spaces()\n", file_id)
		return false
	} else {
		// the space is "too big" so we need to
		//   split the space into two
		//   swap the first bit with the block
		// fmt.Printf("stepP2(%v) space too big (size=%v), split it()\n", file_id, space.size)
		// fmt.Printf("pre-split, block count is %v\n", len(disk.blocks))
		disk.split_space(space_index, block.size)
		// fmt.Printf("post-split, block count is %v\n", len(disk.blocks))
		block_index, block := disk.get_block(file_id)
		space_index, _ := disk.get_first_available_space(block.size)
		// fmt.Printf("stepP2(%v) swapping space_index=%v for block_index=%v\n", file_id, space_index, block_index+1)
		// fmt.Println("pre-swap:")
		// disk.debugP2()
		disk.swap(space_index, block_index)
		// fmt.Println("post-swap:")
		// disk.debugP2()
		fmt.Println()

	}
	return false
}

// takes a contiguous space blocksand splits into two blocks, the first of size 'size'
func (disk *Disk) split_space(index int, size int) {
	block := disk.blocks[index]
	lblock := NewBlock(-1)
	lblock.size = size
	rblock := NewBlock(-1)
	rblock.size = block.size - size
	// fmt.Printf("split_space(index=%v, size=%v), lblock.size=%v\n", index, size, lblock.size)
	// fmt.Printf("split_space(index=%v, size=%v), rblock.size=%v\n", index, size, rblock.size)

	blocks := make([]*Block, 0)
	for i := 0; i < len(disk.blocks); i++ {
		if i != index {
			blocks = append(blocks, disk.blocks[i])
		} else {
			blocks = append(blocks, lblock)
			blocks = append(blocks, rblock)
		}
	}
	disk.blocks = blocks
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

func (disk *Disk) debugP2() {
	for _, b := range disk.blocks {
		id_str := fmt.Sprintf("%v", b.id)
		if b.id > -1 {
			output := strings.Repeat(id_str, b.size)
			fmt.Print(output)
		} else {
			output := strings.Repeat(".", b.size)
			fmt.Print(output)
		}
	}
	fmt.Println()
}
func (puzzle *Puzzle) p1(data string, debug bool) int {
	fmt.Println()
	d := NewDiskP1(data)
	d.Defrag(debug)
	return d.checksum()
}

func (puzzle *Puzzle) p2(data string, debug bool) int {
	fmt.Println()
	d := NewDiskP2(data)
	d.DefragP2(debug)
	return d.checksumP2()
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
	answer1 := puzzle.p2(TEST_DATA_1, true)
	answer2 := puzzle.p2(TEST_DATA_2, true)
	answer3 := puzzle.p2(REAL_DATA, false)
	fmt.Printf("a1: %v\na2: %v\na3: %v\n", answer1, answer2, answer3)
}

func (puzzle *Puzzle) Run() {
	// puzzle.Part1()
	puzzle.Part2()
}
