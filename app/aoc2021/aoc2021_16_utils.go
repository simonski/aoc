package aoc2021

import (
	"fmt"

	utils "github.com/simonski/aoc/utils"
	"github.com/simonski/goutils"
)

type PacketD16 struct {
	Counter      int
	Version      int
	Header       string
	TypeID       int
	LiteralValue uint64
	LengthTypeID string
	Subpackets   []*PacketD16

	NumberOfSubpacketsRequired            int
	TotalLengthInBitsOfSubpacketsRequired int
	TotalBitsRead                         int
	Line                                  string
	SubpacketData                         string
	Value                                 uint64
	Literals                              []string
}

func (p *PacketD16) CalculateVersionTotal() int {
	total := p.Version
	for _, sp := range p.Subpackets {
		sp_total := sp.CalculateVersionTotal()
		total += sp_total
	}
	return total
}

// func (p *PacketD16) RootCalculateTotalBitsRead() uint64 {
// 	return p.CalculateTotalBitsRead() + p.TotalBitsRead
// }

// func (p *PacketD16) CalculateTotalBitsRead() uint64 {
// 	if p.IsTypeLiteral() { //} len(p.Subpackets) == 0 {
// 		// return uint64(len(p.Line)) // p.TotalBitsRead
// 		return p.TotalBitsRead
// 	} else {
// 		// total := uint64(p.TotalBitsRead)
// 		// total := uint64(len(p.Line)) // p.TotalBitsRead)
// 		total := uint64(0) // p.TotalBitsRead
// 		for _, sp := range p.Subpackets {
// 			total += sp.CalculateTotalBitsRead()
// 			// if sp.LengthTypeID == "0" {
// 			// 	total += sp.CalculateTotalBitsRead()
// 			// } else {
// 			// 	total += sp.CalculateTotalBitsRead()
// 			// }
// 		}
// 		return total + p.TotalBitsRead
// 	}
// }

func (p *PacketD16) CalculateTotalBitsRead() int {
	return p.TotalBitsRead
}

func (p *PacketD16) RootCalculateTotalBitsRead() int {
	return p.TotalBitsRead
}

func (p *PacketD16) NumberOfSubpackets() int {
	return len(p.Subpackets)
}

func (p *PacketD16) Debug() {
	p.DebugRecurse("", 0)
}

func (p *PacketD16) DebugRecurse(preprefix string, depth int) {
	line := ""
	prefix := goutils.Repeatstring(" ", depth)
	if p.IsTypeLiteral() {
		line = fmt.Sprintf("%v%vLiteral(%v), Version %v, TypeID %v, Value: %v, requires %v total bits, itself is size %v, has total size of %v bits, requires %v packets, contains %v subpackets:", prefix, preprefix, p.GetOperatorTypeString(), p.Version, p.TypeID, p.GetValue(), p.TotalLengthInBitsOfSubpacketsRequired, p.TotalBitsRead, p.CalculateTotalBitsRead(), p.NumberOfSubpacketsRequired, len(p.Subpackets))
		fmt.Println(line)
	} else {
		_, reason := p.IsSubpacketRequired()
		if p.TotalLengthInBitsOfSubpacketsRequired > 0 {
			line = fmt.Sprintf("%v%vOperator(%v), Version %v, TypeID %v, Value: %v, requires %v total bits, itself+children=%v, itself is size %v, has total size of %v bits, requires %v packets, contains %v subpackets:",
				prefix, preprefix, p.GetOperatorTypeString(), p.Version, p.TypeID, p.GetValue(), p.TotalLengthInBitsOfSubpacketsRequired, p.RootCalculateTotalBitsRead(), p.TotalBitsRead, p.CalculateTotalBitsRead(), p.NumberOfSubpacketsRequired, len(p.Subpackets))
			required := p.TotalLengthInBitsOfSubpacketsRequired
			actual := p.CalculateTotalBitsRead()
			difference := int(required) - int(actual)
			if difference != 0 {
				bitsRequiredInfo := fmt.Sprintf("\n\n\n%v !!!! %v this packet requires bits : %v read %v, remaining %v !!! \n\n\n", preprefix, reason, required, actual, difference)
				fmt.Printf("%v\n", bitsRequiredInfo)
			}
		} else if p.NumberOfSubpacketsRequired > 0 {
			line = fmt.Sprintf("%v%vOperator(%v), Version %v, TypeID %v, Value: %v, requires %v total bits, itself+children=%v, itself is size %v, has total size of %v bits, requires %v packets, contains %v subpackets:", prefix, preprefix, p.GetOperatorTypeString(), p.Version, p.TypeID, p.GetValue(), p.TotalLengthInBitsOfSubpacketsRequired, p.RootCalculateTotalBitsRead(), p.TotalBitsRead, p.CalculateTotalBitsRead(), p.NumberOfSubpacketsRequired, len(p.Subpackets))
			if len(p.Subpackets) != p.NumberOfSubpacketsRequired {
				line += fmt.Sprintf("\n\n\n%v !!!! %v This packet expects %v packets but received %v packets !!!\n\n\n", preprefix, reason, p.NumberOfSubpacketsRequired, len(p.Subpackets))
			}
		}

		fmt.Println(line)
		spDepth := depth + 1
		for index, sp := range p.Subpackets {
			preprefix = fmt.Sprintf("%v/%v ", index+1, len(p.Subpackets))
			sp.DebugRecurse(preprefix, spDepth)
			// line = fmt.Sprintf("%v\n%v", line, sp_line)
			// fmt.Println(line)
		}
	}
	// return line
}

// indicates if this packet "needs" another subpacket
func (p *PacketD16) IsSubpacketRequired() (bool, string) {
	if p.IsTypeLiteral() {
		return false, "Does NOT require a subpacket as it is a literal."

	} else if p.IsTypeOperatorSubpacketsRequired() {
		if p.NumberOfSubpacketsRequired != len(p.Subpackets) {
			return true, fmt.Sprintf("Requires subpackets as it is a '"+p.GetOperatorTypeString()+"' and it does not have enough yet. (%v/%v)", len(p.Subpackets), p.NumberOfSubpacketsRequired)
		}
		return false, fmt.Sprintf("Subpacket not required, matches on number of subpackets (%v).", p.NumberOfSubpacketsRequired)

	} else if p.IsTypeOperatorBitsRequired() {
		if p.TotalLengthInBitsOfSubpacketsRequired != p.CalculateTotalBitsRead() {
			return true, fmt.Sprintf("Requires packets (it is of type '"+p.GetOperatorTypeString()+"') and the bits rerquired (%v) are not satisfied yet (%v).", p.TotalLengthInBitsOfSubpacketsRequired, p.CalculateTotalBitsRead())
		}
		return false, fmt.Sprintf("Subpacket not required, matches on bits (%v).", p.TotalLengthInBitsOfSubpacketsRequired)
	}
	return true, "??? not sure what this case is for IsSubpacketRequired"
}

func (p *PacketD16) Add(subpacket *PacketD16) {
	p.Subpackets = append(p.Subpackets, subpacket)
	// requiresMore := p.RequiresSubpacket()
	// fmt.Printf(">>>>>>>> Packet.Add(type=%v), now contains %v subpackets, needs more %v?\n", p.GetOperatorTypeString(), p.Size(), requiresMore)
}

func (p *PacketD16) Tree(depth int) {
	prefix := goutils.Repeatstring(" ", depth)
	if p.NumberOfSubpacketsRequired == len(p.Subpackets) {
		fmt.Printf("%v(%v)\n", prefix, p.Counter)

	} else if p.NumberOfSubpacketsRequired != 0 {
		missing := p.NumberOfSubpacketsRequired - len(p.Subpackets)
		fmt.Printf("%v(%v)    (missing %v subpackets)\n", prefix, p.Counter, missing)

	} else if p.TotalLengthInBitsOfSubpacketsRequired > 0 && p.TotalLengthInBitsOfSubpacketsRequired != p.CalculateTotalBitsRead() {
		missing := p.TotalLengthInBitsOfSubpacketsRequired - p.CalculateTotalBitsRead()
		fmt.Printf("%v(%v)    (missing %v bits)\n", prefix, p.Counter, missing)

	} else {
		fmt.Printf("%v(%v)\n", prefix, p.Counter)
	}
	for _, sp := range p.Subpackets {
		sp.Tree(depth + 1)
	}
}

func (p *PacketD16) GetValue() uint64 {
	// fmt.Printf("packet(%v).GetValue()...\n", p.Counter)
	if p.IsTypeLiteral() {
		return p.LiteralValue

	} else if p.IsTypeOperatorEq() {
		if len(p.Subpackets) != 2 {
			fmt.Println(">>>>>>>>>>")
			fmt.Println(">>>>>>>>>> EQ is wrong")
			fmt.Println(">>>>>>>>>>")
			return 0
		} else {
			p1 := p.Subpackets[0]
			p2 := p.Subpackets[1]
			if p1.GetValue() == p2.GetValue() {
				return 1
			} else {
				return 0
			}
		}
	} else if p.IsTypeOperatorGt() {
		if len(p.Subpackets) != 2 {
			fmt.Println(">>>>>>>>>>")
			fmt.Println(">>>>>>>>>> GT is wrong")
			fmt.Println(">>>>>>>>>>")
		} else {
			p1 := p.Subpackets[0]
			p2 := p.Subpackets[1]
			if p1.GetValue() > p2.GetValue() {
				return 1
			} else {
				return 0
			}
		}
	} else if p.IsTypeOperatorLt() {
		if len(p.Subpackets) != 2 {
			fmt.Println(">>>>>>>>>>")
			fmt.Println(">>>>>>>>>> LT is wrong")
			fmt.Println(">>>>>>>>>>")
		} else {
			p1 := p.Subpackets[0]
			p2 := p.Subpackets[1]
			if p1.GetValue() < p2.GetValue() {
				return 1
			} else {
				return 0
			}
		}
	} else if p.IsTypeOperatorMax() {
		p1 := p.Subpackets[0]
		max := p1.GetValue()
		for _, sp := range p.Subpackets {
			max = utils.MaxU64(max, sp.GetValue())
		}
		return max
	} else if p.IsTypeOperatorMin() {
		p1 := p.Subpackets[0]
		min := p1.GetValue()
		for _, sp := range p.Subpackets {
			min = utils.MinU64(min, sp.GetValue())
		}
		return min
	} else if p.IsTypeOperatorSum() {
		sum := uint64(0)
		for _, sp := range p.Subpackets {
			sum += sp.GetValue()
		}
		return sum
	} else if p.IsTypeOperatorProduct() {
		total := uint64(1)
		for _, sp := range p.Subpackets {
			total *= sp.GetValue()
		}
		return total
	}
	return uint64(0)
}

func (p *PacketD16) GetOperatorTypeString() string {
	if p.IsTypeOperatorEq() {
		return "eq"
	} else if p.IsTypeOperatorGt() {
		return "gt"
	} else if p.IsTypeOperatorLt() {
		return "lt"
	} else if p.IsTypeOperatorMax() {
		return "max"
	} else if p.IsTypeOperatorMin() {
		return "min"
	} else if p.IsTypeOperatorProduct() {
		return "product"
	} else if p.IsTypeOperatorSum() {
		return "sum"
	} else {
		return "?"
	}
}
func (p *PacketD16) IsTypeLiteral() bool {
	return p.TypeID == 4
}

func (p *PacketD16) IsTypeOperatorSubpacketsRequired() bool {
	return !p.IsTypeLiteral() && p.LengthTypeID == "1"
}

func (p *PacketD16) IsTypeOperatorBitsRequired() bool {
	return !p.IsTypeLiteral() && p.LengthTypeID == "0"
}

// their value is the sum of the sub-packets
func (p *PacketD16) IsTypeOperatorSum() bool {
	return p.TypeID == 0
}

// their value is the product of multiplying together their sub-packets
func (p *PacketD16) IsTypeOperatorProduct() bool {
	return p.TypeID == 1
}

// their value is the min value of their subpackets
func (p *PacketD16) IsTypeOperatorMin() bool {
	return p.TypeID == 2
}

// their value is the max value of their subpackets
func (p *PacketD16) IsTypeOperatorMax() bool {
	return p.TypeID == 3
}

// their value is 1 if the first subpacket is greater, 0 if not
func (p *PacketD16) IsTypeOperatorGt() bool {
	return p.TypeID == 5
}

// their value is 1 if the first subpacket is lesser, 0 if not
func (p *PacketD16) IsTypeOperatorLt() bool {
	return p.TypeID == 6
}

// their value is 1 if the both subpackets are equal, 0 if not
func (p *PacketD16) IsTypeOperatorEq() bool {
	return p.TypeID == 7
}

func take(length int, data string, count int) (string, string, int) {
	left := data[0:length]
	right := data[length:]
	count += length
	return left, right, count
}

func readLiterals(data string) ([]string, uint64, string, int) {

	results := make([]string, 0)
	for index := 0; index < len(data); index += 5 {
		// if index+5 >= len(data) {
		// 	break
		// }
		subvalue := data[index : index+5]
		results = append(results, subvalue)
		if subvalue[0:1] == "0" {
			// then this is the end
			break
		}
	}
	lengthOfLiterals := len(results) * 5

	final_s := ""
	for _, s := range results {
		s := s[1:]
		final_s += s
	}

	intValue := utils.BinaryStringToUInt64(final_s)
	// fmt.Printf("length of literals is %v\n", lengthOfLiterals)
	return results, intValue, data[lengthOfLiterals:], lengthOfLiterals
}

type ContextD16 struct {
	Root    *PacketD16
	Counter int
	Stack   []*PacketD16
	DEBUG   bool
}

func NewContextD16() *ContextD16 {
	ctx := &ContextD16{Stack: make([]*PacketD16, 0)}
	return ctx
}

func (ctx *ContextD16) Debug(msg string) {
	if ctx.DEBUG {
		fmt.Println(msg)
	}
}

func (ctx *ContextD16) Warn(msg string) {
	fmt.Println("WARNING: " + msg)
}

func (ctx *ContextD16) Pop() *PacketD16 {
	p := ctx.Stack[len(ctx.Stack)-1]
	ctx.Stack = ctx.Stack[0 : len(ctx.Stack)-1]
	return p
}

func (ctx *ContextD16) Push(p *PacketD16) {
	ctx.Stack = append(ctx.Stack, p)
	if len(ctx.Stack) == 1 {
		ctx.Root = p
	}
	// prefix := goutils.Repeatstring(" ", ctx.Size())
	// fmt.Printf("%v>>> Context.Push(), size is now %v\n", prefix, ctx.Size())
}

func (ctx *ContextD16) Peek() *PacketD16 {
	if len(ctx.Stack) == 0 {
		return nil
	}
	return ctx.Stack[len(ctx.Stack)-1]
}

func (ctx *ContextD16) Size() int {
	return len(ctx.Stack)
}

func NewPacket(packetVersion int, packetType int, context *ContextD16) *PacketD16 {
	context.Counter += 1
	p := &PacketD16{Version: packetVersion, TypeID: packetType}
	p.Subpackets = make([]*PacketD16, 0)
	p.Counter = context.Counter
	return p
}

func decodeD16(data string) string {
	line := ""
	dm := make(map[string]string)
	dm["0"] = "0000"
	dm["1"] = "0001"
	dm["2"] = "0010"
	dm["3"] = "0011"
	dm["4"] = "0100"
	dm["5"] = "0101"
	dm["6"] = "0110"
	dm["7"] = "0111"
	dm["8"] = "1000"
	dm["9"] = "1001"
	dm["A"] = "1010"
	dm["B"] = "1011"
	dm["C"] = "1100"
	dm["D"] = "1101"
	dm["E"] = "1110"
	dm["F"] = "1111"

	for index := 0; index < len(data); index++ {
		line += dm[data[index:index+1]]
	}
	return line
}
