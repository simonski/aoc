package aoc2021

import (
	"fmt"

	utils "github.com/simonski/aoc/utils"
	"github.com/simonski/goutils"
)

/*
--- Day 16: Packet Decoder ---
As you leave the cave and reach open waters, you receive a transmission from the Elves back on the ship.

The transmission was sent using the Buoyancy Interchange Transmission System (BITS), a method of packing numeric expressions into a binary sequence. Your submarine's computer has saved the transmission in hexadecimal (your puzzle input).

The first step of decoding the message is to convert the hexadecimal representation into binary. Each character of hexadecimal corresponds to four bits of binary data:

0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
A = 1010
B = 1011
C = 1100
D = 1101
E = 1110
F = 1111

The BITS transmission contains a single packet at its outermost layer which itself contains many other packets. The hexadecimal representation of this packet might encode a few extra 0 bits at the end; these are not part of the transmission and should be ignored.

Every packet begins with a standard header: the first three bits encode the packet version, and the next three bits encode the packet type ID. These two values are numbers; all numbers encoded in any packet are represented as binary with the most significant bit first. For example, a version encoded as the binary sequence 100 represents the number 4.

Packets with type ID 4 represent a literal value. Literal value packets encode a single binary number. To do this, the binary number is padded with leading zeroes until its length is a multiple of four bits, and then it is broken into groups of four bits. Each group is prefixed by a 1 bit except the last group, which is prefixed by a 0 bit. These groups of five bits immediately follow the packet header. For example, the hexadecimal string D2FE28 becomes:

4xxx[1xxx][1xxx][1xxx][1xxx][0xxx]

D2FE28
110100101111111000101000    24
VVVTTTAAAAABBBBBCCCCC       21

D2DE28
110|100|10111|11110|00101000    24
         g1  | g2  |  g |      a = 0111 b 1110 c 0101  discard final 3,
		 011111100101 = 2021
[version][typeID]




[110|100]
110 version 6 = VVV
100 type 4 TTT

110100 vesion and type
110100 add padding to multiple of 4
0011, 0100 break to groups
10011, 00100 prefix the groups

AAAAA




Below each bit is a label indicating its purpose:

The three bits labeled V (110) are the packet version, 6.
The three bits labeled T (100) are the packet type ID, 4, which means the packet is a literal value.
The five bits labeled A (10111) start with a 1 (not the last group, keep reading) and contain the first four bits of the number, 0111.
The five bits labeled B (11110) start with a 1 (not the last group, keep reading) and contain four more bits of the number, 1110.
The five bits labeled C (00101) start with a 0 (last group, end of packet) and contain the last four bits of the number, 0101.
The three unlabeled 0 bits at the end are extra due to the hexadecimal representation and should be ignored.
So, this packet represents a literal value with binary representation 011111100101, which is 2021 in decimal.


Every other type of packet (any packet with a type ID other than 4) represent an operator that performs some calculation on one or more sub-packets contained within. Right now, the specific operations aren't important; focus on parsing the hierarchy of sub-packets.

An operator packet contains one or more packets. To indicate which subsequent binary data represents its sub-packets, an operator packet can use one of two modes indicated by the bit immediately after the packet header; this is called the length type ID:

If the length type ID is 0, then the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet.
If the length type ID is 1, then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
Finally, after the length type ID bit and the 15-bit or 11-bit field, the sub-packets appear.

For example, here is an operator packet (hexadecimal string 38006F45291200) with length type ID 0 that contains two sub-packets:

001110 0 000000000011011 110100010100101001000100100 0000000
VVVTTT I LLLLLLLLLLLLLLL AAAAAAAAAAABBBBBBBBBBBBBBBB
The three bits labeled V (001) are the packet version, 1.
The three bits labeled T (110) are the packet type ID, 6, which means the packet is an operator.
The bit labeled I (0) is the length type ID, which indicates that the length is a 15-bit number representing the number of bits in the sub-packets.
The 15 bits labeled L (000000000011011) contain the
 the sub-packets in bits, 27.
The 11 bits labeled A contain the first sub-packet, a literal value representing the number 10.
The 16 bits labeled B contain the second sub-packet, a literal value representing the number 20.
After reading 11 and 16 bits of sub-packet data, the total length indicated in L (27) is reached, and so parsing of this packet stops.

As another example, here is an operator packet (hexadecimal string EE00D40C823060) with length type ID 1 that contains three sub-packets:

111011100000000011 01010000001 10010000010 00110000011 00000
VVVTTTILLLLLLLLLLL AAAAAAAAAAA BBBBBBBBBBB CCCCCCCCCCC
The three bits labeled V (111) are the packet version, 7.
The three bits labeled T (011) are the packet type ID, 3, which means the packet is an operator.
The bit labeled I (1) is the length type ID, which indicates that the length is a 11-bit number representing the number of sub-packets.
The 11 bits labeled L (00000000011) contain the number of sub-packets, 3.
The 11 bits labeled A contain the first sub-packet, a literal value representing the number 1.
The 11 bits labeled B contain the second sub-packet, a literal value representing the number 2.
The 11 bits labeled C contain the third sub-packet, a literal value representing the number 3.
After reading 3 complete sub-packets, the number of sub-packets indicated in L (3) is reached, and so parsing of this packet stops.

For now, parse the hierarchy of the packets throughout the transmission and add up all of the version numbers.

Here are a few more examples of hexadecimal-encoded transmissions:

8A004A801A8002F478 represents an operator packet (version 4) which contains an operator packet (version 1) which contains an operator packet (version 5) which contains a literal value (version 6); this packet has a version sum of 16.
620080001611562C8802118E34 represents an operator packet (version 3) which contains two sub-packets; each sub-packet is an operator packet that contains two literal values. This packet has a version sum of 12.
C0015000016115A2E0802F182340 has the same structure as the previous example, but the outermost packet uses a different length type ID. This packet has a version sum of 23.
A0016C880162017C3686B18A3D4780 is an operator packet that contains an operator packet that contains an operator packet that contains five literal values; it has a version sum of 31.
Decode the structure of your hexadecimal-encoded BITS transmission; what do you get if you add up the version numbers in all packets?
*/

func (app *Application) Y2021D16_Summary() *utils.Summary {
	s := utils.NewSummary(2021, 16)
	s.Name = "Packet Decoder"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

// rename this to the year and day in question
func (app *Application) Y2021D16P1() {

	DEBUG := false

	RUN_1 := false
	RUN_2 := false
	RUN_3 := false
	RUN_4 := false
	RUN_5 := false
	RUN_6 := false
	RUN_7 := false
	RUN_PART_ONE := true

	line := goutils.Repeatstring("*", 120)

	if RUN_1 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse(DAY_2021_16_TEST_DATA_1, true, DEBUG, c, 0)
		RParse(DAY_2021_16_TEST_DATA_1, c)
		total := c.Root.CalculateVersionTotal()
		fmt.Printf("Part1-1: Version total is %v\n", total)
		c.Root.Debug()
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_2 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		RParse("38006F45291200", c)
		total := c.Root.CalculateVersionTotal()
		fmt.Printf("Part1-2: Version total is %v\n", total)
		c.Root.Debug()
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_3 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse("EE00D40C823060", true, DEBUG, c, 0)
		RParse("EE00D40C823060", c)
		total := c.Root.CalculateVersionTotal()
		fmt.Printf("Part1-3: Version total is %v\n", total)
		c.Root.Debug()
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_4 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse("8A004A801A8002F478", true, DEBUG, c, 0)
		RParse("8A004A801A8002F478", c)
		fmt.Printf("Part1-4: Version total should be 16, is %v\n", c.Root.CalculateVersionTotal())
		c.Root.Debug()
		c.Root.Tree(0)
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_5 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse("620080001611562C8802118E34", true, DEBUG, c, 0)
		RParse("620080001611562C8802118E34", c)
		fmt.Printf("Part1-5: Version total should be 12, is %v\n", c.Root.CalculateVersionTotal())
		c.Root.Debug()
		c.Root.Tree(0)
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_6 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse("C0015000016115A2E0802F182340", true, DEBUG, c, 0)
		RParse("C0015000016115A2E0802F182340", c)
		fmt.Printf("Part1-6: Version total should be 23, is %v\n", c.Root.CalculateVersionTotal())
		c.Root.Debug()
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_7 {
		if DEBUG {
			fmt.Println(line)
			fmt.Println(line)
		}
		c := NewContextD16()
		c.DEBUG = true
		// Parse("A0016C880162017C3686B18A3D4780", true, DEBUG, c, 0)
		RParse("A0016C880162017C3686B18A3D4780", c)
		fmt.Printf("Part1-7: Version total should be 31, is %v\n", c.Root.CalculateVersionTotal())
		c.Root.Debug()
		c.Root.Tree(0)
		if DEBUG {
			fmt.Println(line)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	if RUN_PART_ONE {
		c := NewContextD16()
		c.DEBUG = true
		// Parse(DAY_2021_16_DATA, true, DEBUG, c, 0)
		RParse(DAY_2021_16_DATA, c)
		total := c.Root.CalculateVersionTotal()
		c.Root.Debug()
		c.Root.Tree(0)
		fmt.Printf("Version Total should be 920, is %v\n", total)
	}

}

// // rename this to the year and day in question
func (app *Application) Y2021D16P2() {

	if true {
		c := NewContextD16()
		c.DEBUG = true
		RParse("C200B40A82", c)
		p := c.Root
		fmt.Printf("Part2: C200B40A82 value should be 3, is %v\n", p.GetValue())

		c = NewContextD16()
		c.DEBUG = true
		RParse("04005AC33890", c)
		p = c.Root
		fmt.Printf("Part2: 04005AC33890 : Value should be 54, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("880086C3E88112", c)
		p = c.Root
		fmt.Printf("Part2: 880086C3E88112 Value should be 7, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("CE00C43D881120", c)
		p = c.Root
		fmt.Printf("Part2: CE00C43D881120 Value should be 9, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("D8005AC2A8F0", c)
		p = c.Root
		fmt.Printf("Part2: D8005AC2A8F0 Value should be 1, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("F600BC2D8F", c)
		p = c.Root
		fmt.Printf("Part2: F600BC2D8F Value should be 0, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("9C005AC2F8F0", c)
		p = c.Root
		fmt.Printf("Part2: 9C005AC2F8F0 Value should be 0, is %v\n", p.GetValue())
		c.Root.Debug()

		c = NewContextD16()
		c.DEBUG = true
		RParse("9C0141080250320F1802104A08", c)
		p = c.Root
		fmt.Printf("Part2: 9C0141080250320F1802104A08 Value should be 1, is %v\n", p.GetValue())
		c.Root.Debug()
	}

	fmt.Println()
	fmt.Println()
	// fmt.Println()
	// fmt.Println()
	c2 := NewContextD16()
	c2.DEBUG = true
	RParse(DAY_2021_16_DATA, c2)
	// fmt.Printf("DATA Value is %v\n", p2.GetValue())
	c2.Root.Debug()
	v := c2.Root.GetValue()
	fmt.Printf("DATA Value is %v\n", v)

}

// rename and uncomment this to the year and day in question once complete for a gold star}
// func (app *Application) Y20XXDXXP1Render() {
// }

// rename and uncomment this to the year and day in question once complete for a gold star!
// func (app *Application) Y20XXDXXP2Render() {
// }

// this is what we will reflect and call - so both parts with run. It's up to you to make it print nicely etc.
// The app reference has a CLI for logging.
// func (app *Application) Y2021D16() {
// 	app.Y2021D16P1()
// 	app.Y2021D16P2()
// }
