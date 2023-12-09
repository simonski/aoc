package d8

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Data  string
	Index int
}

func NewInstruction(input string) *Instruction {
	return &Instruction{Data: input, Index: -1}
}

func (i *Instruction) Reset() {
	i.Index = -1
}

// L or R
func (i *Instruction) Next() string {
	i.Index += 1
	if i.Index >= len(i.Data) {
		i.Index = 0
	}
	return i.Data[i.Index : i.Index+1]
}

type Directions struct {
	Choices map[string]*Choice
}

type Choice struct {
	Name  string
	Left  string
	Right string
	EndA  bool
	EndZ  bool
}

func NewChoice(input string) *Choice {
	splits := strings.Split(input, " = ")
	name := splits[0]
	splits = strings.Split(splits[1], ", ")
	left := strings.ReplaceAll(splits[0], "(", "")
	right := strings.ReplaceAll(splits[1], ")", "")
	c := &Choice{Name: name, Left: left, Right: right}
	if strings.Contains(input, "A =") {
		c.EndA = true
	} else if strings.Contains(input, "Z =") {
		c.EndZ = true
	}
	return c
}

func NewDirections(lines []string) *Directions {
	choices := make(map[string]*Choice)
	for _, line := range lines {
		c := NewChoice(line)
		choices[c.Name] = c
	}
	d := Directions{}
	d.Choices = choices
	return &d
}

func (d *Directions) FindEndsWith(letter string) []*Choice {
	results := make([]*Choice, 0)
	for name, c := range d.Choices {
		if strings.Index(name, letter) == 2 {
			results = append(results, c)
		}
	}
	return results
}

func Load(input string) (*Instruction, *Directions) {
	lines := strings.Split(input, "\n")
	i := NewInstruction(lines[0])
	d := NewDirections(lines[2:])
	return i, d
}

type P2 struct {
	Position     string
	Moves        []string
	Directions   *Directions
	Instructions *Instruction
	KeepMoves    bool
	MoveCount    int
}

func NewP2(position string, input string, keepmoves bool) *P2 {
	p2 := P2{}
	p2.KeepMoves = keepmoves
	i, d := Load(input)
	moves := make([]string, 0)
	p2.Instructions = i
	p2.Directions = d
	p2.Position = position
	p2.Moves = moves
	return &p2

}

func (p *P2) StepToEndingZ() (int, map[string]int) {
	count := 0
	circularMap := make(map[string]int)
	for {
		// if strings.Index(p.Position, "Z") == 2 {
		// 	fmt.Printf("pre-pos is Z '%v'\n", p.Position)
		// }

		p.Step()

		key := p.Position
		value, exists := circularMap[key]
		if exists {
			value += 1
		} else {
			value = 1
		}
		circularMap[key] = value

		count++
		// fmt.Printf(" -> %v", p.Position)
		if strings.Index(p.Position, "Z") == 2 {
			fmt.Println(p.Position)
			// fmt.Print("\n")
			break
		}

	}
	return count, circularMap
}

func (p *P2) Step() string {

	new_position := ""
	position := p.Position
	instruction := p.Instructions.Next()
	if instruction == "L" {
		new_position = p.Directions.Choices[position].Left
	} else {
		new_position = p.Directions.Choices[position].Right
	}
	if p.KeepMoves {
		p.Moves = append(p.Moves, new_position)
	}
	p.MoveCount += 1
	p.Position = new_position
	return new_position

}

func Steps(startPosition string, endPosition string, input string) (int, []string) {

	i, d := Load(input)
	moves := make([]string, 0)

	// position := "AAA"
	position := startPosition
	new_position := ""
	count := 0
	for {
		instruction := i.Next()
		if instruction == "L" {
			new_position = d.Choices[position].Left
		} else {
			new_position = d.Choices[position].Right
		}
		count++
		moves = append(moves, fmt.Sprintf("[%v] %v -> %v", count, position, new_position))
		position = new_position
		if position == endPosition {
			break
		}
	}

	return count, moves

}
