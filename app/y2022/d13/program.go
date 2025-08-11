package d13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonski/goutils"
)

/*
--- Day 05:  ---

*/

type Puzzle struct {
	title string
	year  string
	day   string
	input string
	lines []string
	Pairs []*Pair
}

type Pair struct {
	Left  *Packet
	Right *Packet
}

type CompareOutcome string

const LEFT CompareOutcome = "left"
const RIGHT CompareOutcome = "right"
const CONTINUE CompareOutcome = "continue"

func (p *Pair) Compare(VERBOSE bool, indent int, left *Packet, right *Packet) CompareOutcome {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix = fmt.Sprintf("%v%v", prefix, " ")
	}
	if VERBOSE {
		fmt.Printf("%vCompare %v vs %v\n", prefix, left.String(), right.String())
	}
	if left.PacketType == INTEGER && right.PacketType == INTEGER {
		if left.Value < right.Value {
			return LEFT
		} else if left.Value > right.Value {
			return RIGHT
		} else {
			// continue
			return CONTINUE
		}
	} else if left.PacketType == LIST && right.PacketType == LIST {
		size := goutils.Max(left.Size(), right.Size())
		for i := 0; i < size; i++ {
			if i >= left.Size() {
				return LEFT
			} else if i >= right.Size() {
				return RIGHT
			}
			lp := left.Entries[i]
			rp := right.Entries[i]
			result := p.Compare(VERBOSE, indent+1, lp, rp)
			if result == LEFT || result == RIGHT {
				return result
			}

		}
		return CONTINUE
	} else {
		// one of them is an integer, one a list
		if left.PacketType == INTEGER {
			leftList := NewPacketList()
			leftList.Add(left)
			return p.Compare(VERBOSE, indent+1, leftList, right)
		} else if right.PacketType == INTEGER {
			rightList := NewPacketList()
			rightList.Add(right)
			return p.Compare(VERBOSE, indent+1, left, rightList)
		}

	}
	return CONTINUE
}

func (p *Pair) IsCorrect() bool {
	VERBOSE := true
	size := goutils.Max(p.Left.Size(), p.Right.Size())
	for index := 0; index < size; index++ {
		result := p.Compare(VERBOSE, 0, p.Left, p.Right)
		if result == LEFT {
			return true
		} else if result == RIGHT {
			return false
		} else if result == CONTINUE {
			// continue
		}
	}
	return false
}

type PacketType string

const INTEGER PacketType = "int"
const LIST PacketType = "list"

type Packet struct {
	data       string
	PacketType PacketType
	Value      int
	Entries    []*Packet
	Parent     *Packet
}

func (p *Packet) String() string {
	result := ""
	if p.PacketType == LIST {
		result = "["
		for index, child := range p.Entries {
			result = fmt.Sprintf("%v%v", result, child.String())
			if index+1 < p.Size() {
				result = fmt.Sprintf("%v,", result)
			}
		}
		result = fmt.Sprintf("%v]", result)
	} else {
		result = fmt.Sprintf("%v", p.Value)
	}
	return result
}

func (p *Packet) Size() int {
	return len(p.Entries)
}

func (p *Packet) Add(child *Packet) {
	p.Entries = append(p.Entries, child)
	child.Parent = p
}

// [[[1,[1,6,0,10,10],[],5,[]],[3,[5,8],2,[3,6,5,7],8],4,[0,[10]],8],[7,7],[[3,[9,5,10],[3],8,10],2,3,9],[[8,[3,1,4,9,5]]],[[[],0],10,10,3]]

func nextEntry(data string) (string, string) {
	if data == "" {
		return "", ""
	}
	entry := data[0:1]
	fmt.Printf("first character is '%v'\n", entry)
	if entry == "[" || entry == "]" || entry == "," {
		remainder := data[1:]
		fmt.Printf("returning characeter '%v', remainder '%v'\n", entry, remainder)
		return entry, remainder
	} else {
		lastgood := ""
		for index := 0; index < len(data); index++ {
			attempt := data[0 : index+1]
			_, err := strconv.Atoi(attempt)
			if err == nil {
				lastgood = attempt
				fmt.Printf("lastgood is %v\n", attempt)
			} else {
				remainder := data[len(lastgood):]
				fmt.Printf("failed at '%v', lastgood is '%v', remainder='%v'\n", attempt, lastgood, remainder)
				return lastgood, remainder
			}
		}
		if lastgood == data {
			return lastgood, ""
		}

	}
	return "", data
}

func BuildPacket(VERBOSE bool, data string) *Packet {
	entry := ""
	root := NewPacketList()
	root.data = data
	new_data := data[1 : len(data)-1] // strip the first and last [ and ]
	if VERBOSE {
		fmt.Printf("data=%v, new_data=%v\n", data, new_data)
	}
	data = new_data
	var current *Packet
	current = root
	iteration := 0
	for {
		iteration++
		entry, data = nextEntry(data)
		if entry == "[" {
			if VERBOSE {
				fmt.Printf("[%v] entry starts a packet", iteration)
			}
			// start a list
			packet := NewPacketList()
			current.Add(packet)
			current = packet
		} else if entry == "]" {
			if VERBOSE {
				fmt.Printf("[%v] entry stops the packet", iteration)
			}
			// end the current list
			current = current.Parent
		} else if entry == "," {
			if VERBOSE {
				fmt.Printf("[%v] entry is a comma", iteration)
			}
		} else if entry == "" {
			if VERBOSE {
				fmt.Printf("[%v] entry is empty, will quit", iteration)
			}
			// at the end
			break
		} else {
			// it is a n integer value
			if VERBOSE {
				fmt.Printf("[%v] entry is a number\n", iteration)
			}
			value := NewPacketValue(entry)
			current.Add(value)
			fmt.Printf("current contains %v, root contains %v\n", current.Size(), root.Size())
		}

	}
	return root
}

func NewPacketList() *Packet {
	p := Packet{}
	p.Entries = make([]*Packet, 0)
	p.PacketType = LIST
	return &p
}

func NewPacketValue(entry string) *Packet {
	p := Packet{}
	p.PacketType = INTEGER
	value, _ := strconv.Atoi(entry)
	p.Value = value
	return &p
}

func NewPuzzleWithData(input string) *Puzzle {
	p := Puzzle{year: "2022", day: "13", title: "Distress Signal"}
	p.Load(input)
	return &p
}

func NewPuzzle() *Puzzle {
	return NewPuzzleWithData(REAL_DATA)
}

func (puzzle *Puzzle) Load(input string) {
	lines := strings.Split(input, "\n")
	puzzle.input = input
	puzzle.lines = lines
	pairs := make([]*Pair, 0)
	VERBOSE := false
	for index := 0; index < len(lines); index += 3 {
		line1 := lines[index]
		line2 := lines[index+1]
		pair := &Pair{Left: BuildPacket(VERBOSE, line1), Right: BuildPacket(VERBOSE, line2)}
		pairs = append(pairs, pair)
	}
	puzzle.Pairs = pairs
}

func (puzzle *Puzzle) Part1() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Part2() {
	puzzle.Load(REAL_DATA)
}

func (puzzle *Puzzle) Run() {
	puzzle.Part1()
	puzzle.Part2()
}
