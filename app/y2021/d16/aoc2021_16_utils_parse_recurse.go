package aoc2021

import (
	"fmt"

	utils "github.com/simonski/aoc/utils"
)

// reads from the data the next packet (does NOT read further)
func ReadPacketHeader(datax string, context *ContextD16) (*PacketD16, string) {
	bitsRead := 0
	header, remainder, _ := take(6, datax, 0)
	packetVersion := utils.BinaryStringToInt(header[0:3])
	packetType := utils.BinaryStringToInt(header[3:6])
	packet := NewPacket(packetVersion, packetType, context)
	prefix := fmt.Sprintf("(%v) ", packet.Counter)
	packet.Header = header

	// fmt.Println("Enter ReadPacketHeader")
	// fmt.Println(datax)
	// fmt.Println(header)
	// fmt.Println(remainder)
	// fmt.Println()

	if packet.IsTypeLiteral() {
		fmt.Printf("%vIsLiteral\n", prefix)
		length := 0
		literals, value, remainderx, length := readLiterals(remainder)
		bitsRead += length
		packet.LiteralValue = value
		packet.Literals = literals
		packet.TotalBitsRead = bitsRead
		remainder = remainderx
		context.Debug(fmt.Sprintf("%vReadPacketHeader(Literal), value=%v", prefix, packet.LiteralValue))
		fmt.Println(remainder)

	} else {
		lengthTypeID, remainderx, read := take(1, remainder, bitsRead)
		// fmt.Println(remainder)
		remainder = remainderx
		bitsRead += read
		packet.LengthTypeID = lengthTypeID
		if packet.IsTypeOperatorBitsRequired() {
			// fmt.Printf("IsOperatorBitsRequired, remainder=\n%v\n", remainder)
			bits, remainderx, read := take(15, remainder, bitsRead)
			fmt.Println(bits)
			fmt.Println(remainderx)

			remainder = remainderx
			bitsRead += read
			bitsAsInt := utils.BinaryStringToInt(bits)
			fmt.Printf("%vIsOperatorBitsRequired, bits=%v, bitsRequired=%v, remainder=\n%v\n", prefix, bits, bitsAsInt, remainder)

			// this bit might be unhelpful
			if packet.IsTypeOperatorEq() || packet.IsTypeOperatorGt() || packet.IsTypeOperatorLt() {
				packet.NumberOfSubpacketsRequired = 2
			}
			packet.TotalBitsRead = bitsRead

			// extract the subset of data to read
			packet.TotalLengthInBitsOfSubpacketsRequired = bitsAsInt

			// fmt.Printf("\nThe total number of bits to read is %v\n", bitsAsInt)

			bits, remainderx, _ = take(packet.TotalLengthInBitsOfSubpacketsRequired, remainder, 0)
			packet.SubpacketData = bits
			remainder = remainderx

			context.Debug(fmt.Sprintf("%vReadPacketHeader(%v) bitsRequired=%v", prefix, packet.GetOperatorTypeString(), packet.TotalLengthInBitsOfSubpacketsRequired))

		} else if packet.IsTypeOperatorSubpacketsRequired() {
			// otherwise this operator type is one that requires a number of packets
			fmt.Printf("%vIsOperatorSubpackets\n", prefix)
			bits, remainderx, read := take(11, remainder, bitsRead)
			fmt.Println(bits)
			fmt.Println(remainderx)

			remainder = remainderx
			bitsRead += read

			bitsInt := utils.BinaryStringToInt(bits)
			packet.NumberOfSubpacketsRequired = bitsInt
			packet.TotalBitsRead = bitsRead

			context.Debug(fmt.Sprintf("%vReadPacketHeader(%v) packetsRequired=%v", prefix, packet.GetOperatorTypeString(), packet.NumberOfSubpacketsRequired))
		} else {
			context.Warn(fmt.Sprintf("%vReadPacketHeader(??? cannot tell type.", prefix))

		}

	}

	context.Debug(fmt.Sprintf("%vExiting ReadPacketHeader, entered with \n%v\nexited with\n%v\n", prefix, datax, remainder))

	return packet, remainder
}

func RParse(data string, context *ContextD16) *PacketD16 {
	decoded := decodeD16(data)
	context.Debug(fmt.Sprintf("\ndata    : %v", data))
	context.Debug(fmt.Sprintf("decoded : %v\n", decoded))
	root, remainder := ReadPacketHeader(decoded, context)
	context.Root = root
	required, _ := root.IsSubpacketRequired()
	if required {
		RecurseParse(root, remainder, context)
	}
	return root
}

func RecurseParse(packet *PacketD16, data string, context *ContextD16) string {

	prefix := fmt.Sprintf("%v ", packet.Counter)
	context.Debug(fmt.Sprintf("RecurseParse: %v\n", data))
	if packet.IsTypeLiteral() {
		context.Debug(fmt.Sprintf("%vRecurseParse, Literal", prefix))
		return data
	}

	if packet.IsTypeOperatorBitsRequired() {
		// then this packet already has the data in packet.Subpacketdata
		// so we just need to decode all of that
		// in this case we want to use the data that we already
		context.Debug(fmt.Sprintf("%vRecurseParse, OperatorBits", prefix))
		remainder := packet.SubpacketData
		for {
			child, r := ReadPacketHeader(remainder, context)
			remainder = r
			packet.Add(child)
			fmt.Printf("%v now has %v children.", prefix, len(packet.Subpackets))
			req, _ := child.IsSubpacketRequired()
			if req {
				fmt.Printf("%v the new child (%v) is an operator, so we recurse into it.\n", prefix, child.Counter)
				remainder = RecurseParse(child, remainder, context)
			} else {
				fmt.Printf("%v the new child (%v) is NOT an operator, no recusion.\n", prefix, child.Counter)
			}
			// remainder = r
			if len(remainder) < 6 {
				fmt.Printf("%v we have read all the packets we can, no more children for this parent.", prefix)
				break
			}
		}
		// now we have read all teh data, we shoud mark it
		packet.TotalBitsRead = packet.TotalLengthInBitsOfSubpacketsRequired
		return data

	} else if packet.IsTypeOperatorSubpacketsRequired() {
		context.Debug(fmt.Sprintf("%vRecurseParse, OperatorPackets", prefix))
		remainder := data
		for {
			required, reason := packet.IsSubpacketRequired()
			context.Debug(fmt.Sprintf("%v %v", required, reason))
			if required {
				context.Debug(fmt.Sprintf("%vRecurseParse, OperatorPackets, requieres more", prefix))
				context.Debug(fmt.Sprintf("%vRemainder before parsing is : %v", prefix, remainder))

				child, r := ReadPacketHeader(remainder, context)
				packet.Add(child)
				req, _ := child.IsSubpacketRequired()
				if req {
					r = RecurseParse(child, r, context)
				}
				remainder = r
			} else {
				context.Debug(fmt.Sprintf("%vRecurseParse, OperatorPackets no longer required.", prefix))
				break
			}
		}
		return remainder

	}
	return data
}

// func foo() {
// 	remainder := data
// 	if DEBUG {
// 		fmt.Printf("Parse(depth=%v, data=%v, length=%v)\n", depth, data, len(remainder))
// 	}
// 	// original := remainder
// 	var header string
// 	var literals []string
// 	var value uint64
// 	var lengthTypeID string
// 	var bits string
// 	// var line string

// 	versionTotal := 0
// 	var packet *PacketD16

// 	if DEBUG {
// 		fmt.Printf("%v\n%v\n", data, remainder)
// 		fmt.Println()
// 	}

// 	for {
// 		// bitsRead, packet : readPacketHeader(remainder)
// 		bitsRead := 0
// 		// header, remainder, bitsRead = take(6, remainder, bitsRead)
// 		// packetVersion := utils.BinaryStringToInt(header[0:3])
// 		// versionTotal += packetVersion
// 		// packetType := utils.BinaryStringToInt(header[3:6])
// 		// packet := NewPacket(packetVersion, packetType, context)

// 		packet, remainder = readPacketHeader(remainder, context)
// 		line := packet.Header
// 		bitsRead += 6
// 		versionTotal += packet.Version

// 		logPrefix := fmt.Sprintf("(%v) ", packet.Counter)
// 		// add this to current entry in the stack if it still wants more
// 		var parent *PacketD16
// 		if context.Size() == 0 {
// 			if DEBUG {
// 				fmt.Printf("%vContext: is empty, setting root and parent.\n", logPrefix)
// 			}
// 			context.Push(packet)
// 			parent = packet
// 		} else {
// 			parent = context.Peek()

// 			for {
// 				canBreak := true
// 				required, _ := parent.IsSubpacketRequired()
// 				if !required {
// 					if DEBUG {
// 						fmt.Printf("<<<<<<<<<<<<<< POP (subpackets matched)!\n")
// 					}
// 					if context.Size() > 1 {
// 						context.Pop()
// 						canBreak = false
// 					}
// 				}

// 				// if parent.TotalLengthInBitsOfSubpacketsRequired > 0 {
// 				// 	if parent.CalculateTotalBitsRead() == parent.TotalLengthInBitsOfSubpacketsRequired {
// 				// 		if DEBUG {
// 				// 			fmt.Printf("<<<<<<<<<<<<<< POP (bits read)!\n")
// 				// 		}
// 				// 		if context.Size() > 1 {
// 				// 			context.Pop()
// 				// 			canBreak = false
// 				// 		}
// 				// 	}
// 				// }
// 				if canBreak {
// 					break
// 				}
// 			}
// 			parent = context.Peek()

// 			parent.Add(packet)
// 			if DEBUG {
// 				fmt.Printf("%vAdded this (%v) to parent %v, giving tree\n\n", logPrefix, packet.Counter, parent.Counter)
// 				// context.Root.Tree(0)
// 				fmt.Printf("\n\n")
// 			}
// 		}

// 		if packet.IsTypeLiteral() {
// 			if DEBUG {
// 				fmt.Printf("%vPacket(Literal): Version=%v, Type=%v, header=%v\n", logPrefix, packet.Version, packet.TypeID, header)
// 			}
// 			// then we can get all the literals as strings and the remaining text
// 			length := 0
// 			literals, value, remainder, length = readLiterals(remainder)
// 			for _, literal := range literals {
// 				line += literal
// 			}
// 			bitsRead += length
// 			// context.Peek().TotalBitsRead = bitsRead
// 			packet.LiteralValue = value
// 			packet.Literals = literals
// 			packet.TotalBitsRead = bitsRead
// 			if DEBUG {
// 				fmt.Printf("Now let us work out what our parent thinks, parent=%v\n", parent.Counter)
// 				fmt.Printf("The parent (%v) wants %v and thinks %v have been read before we update this literal, which is size %v.\n", parent.Counter, parent.TotalLengthInBitsOfSubpacketsRequired, parent.CalculateTotalBitsRead(), packet.TotalBitsRead)
// 			}
// 			// parent.Add(packet)
// 			if DEBUG {
// 				fmt.Printf("This packet has now been addded to the parent, size in bits added is %v, \n", packet.TotalBitsRead)
// 				fmt.Printf("The parent (%v) wants %v and NOW thinks %v have been read.\n", parent.Counter, parent.TotalLengthInBitsOfSubpacketsRequired, parent.CalculateTotalBitsRead())
// 			}

// 		} else {
// 			if DEBUG {
// 				fmt.Printf("Packet(Operator): Version=%v, Type=%v, header=%v\n", packet.Version, packet.TypeID, header)
// 			}
// 			// it's an operator packet
// 			context.Push(packet)
// 			lengthTypeID, remainder, bitsRead = take(1, remainder, bitsRead)
// 			line += lengthTypeID
// 			packet.LengthTypeID = lengthTypeID
// 			if DEBUG {
// 				opType := packet.GetOperatorTypeString()
// 				fmt.Printf("Operator, TypeID=%v, LengthTypeID=%v, (%v)\n", packet.TypeID, packet.LengthTypeID, opType)
// 			}

// 			if packet.IsTypeOperatorBitsRequired() {
// 				bits, remainder, bitsRead = take(15, remainder, bitsRead)
// 				line += bits
// 				bitsAsInt := utils.BinaryStringToInt(bits)
// 				if DEBUG {
// 					fmt.Printf("15 bits = %v (%v is the length of the subpacket data)\n", bits, bitsAsInt)
// 					// fmt.Printf("%v = remainder before stripping %v bits of subpacket data\n", remainder, bitsAsInt)
// 				}

// 				// OK so now I think I DO need to call Parse again, as I want
// 				// to discard the stuff at the end, then resume what I'm doing.
// 				if packet.IsTypeOperatorEq() || packet.IsTypeOperatorGt() || packet.IsTypeOperatorLt() {
// 					packet.NumberOfSubpacketsRequired = 2
// 				}
// 				packet.TotalBitsRead = bitsRead

// 				// extract the subset of data to read
// 				packet.TotalLengthInBitsOfSubpacketsRequired = bitsAsInt
// 				bits, remainder, bitsRead = take(bitsAsInt, remainder, bitsRead)

// 				// so the bits "bits" are the subset of data to process
// 				// remainder is everything to the right of the data we want to process

// 				// so now "bits" contains the whole set of data that neeeds to be parsed, then "remainder" is everything
// 				// after that
// 				// fmt.Println()
// 				// fmt.Printf("bits      : %v\n", bits)
// 				// fmt.Printf("remainder : %v\n", remainder)
// 				// fmt.Printf("this parent requires %v, it has read for itself %v, the section to parse is length %v\n", packet.TotalLengthInBitsOfSubpacketsRequired, packet.TotalBitsRead, len(bits))
// 				fmt.Println()
// 				subcontext := NewContextD16()
// 				subcontext.Push(packet)
// 				subcontext.Counter = context.Counter
// 				// if DEBUG {
// 				// 	fmt.Print("Calling parse internally\n")
// 				// 	fmt.Printf("%v\n", bits)
// 				// }
// 				depth += 1
// 				// before := len(packet.Subpackets)
// 				Parse(bits, false, DEBUG, subcontext, depth)
// 				if packet.TotalLengthInBitsOfSubpacketsRequired > 0 {
// 					packet.TotalBitsRead = len(bits) //bitsRead
// 				}
// 				// after := len(packet.Subpackets)
// 				// requires, reason := packet.IsSubpacketRequired()
// 				// if requires {
// 				// 	fmt.Printf("\n\n\n\n\n for some reason after sub-parsing, this requires a child!\n\n%v\n\n\n\n\n", reason)
// 				// }
// 				// if before == after {
// 				// 	fmt.Printf("\n\n\n\n\n for some reason after sub-parsing, no children were added!\n\n%v\n\n\n\n\n", reason)
// 				// } else {
// 				// 	fmt.Printf("\n\n\n\n\n good for some reason after sub-parsing, there were %v before and %v after : %v!\n\n\n\n\n\n\n", before, after, reason)
// 				// }

// 				// fmt.Printf("all the data has been read, setting the parent size to the required size (Wants %v, has read %v, parsed %v) \n", packet.TotalLengthInBitsOfSubpacketsRequired, packet.TotalBitsRead, len(bits))

// 				// packet.TotalBitsRead = packet.TotalLengthInBitsOfSubpacketsRequired
// 				// fmt.Println()
// 				// fmt.Println(">>>>")
// 				// fmt.Printf("remainder is %v, '%v'\n", len(remainder), remainder)
// 				// fmt.Println(">>>>")
// 				// fmt.Println()

// 				if DEBUG {
// 					fmt.Print("/Called parse internally\n")
// 					fmt.Printf("this resulted in a counter now being %v\n", subcontext.Counter)
// 					fmt.Print("DEBUGGING the subp:\n\n")
// 					// subcontext.Root.Tree(0)
// 					fmt.Printf("%v\n", subcontext.Root.Counter)
// 					// subcontext.Root.Debug()
// 					fmt.Print("/DEBUGGING the context:\n\n")
// 					// context.Root.Tree(0)
// 					fmt.Print("/DEBUGGING the context:\n\n")
// 				}
// 				context.Counter = subcontext.Counter

// 			} else if packet.IsTypeOperatorSubpacketsRequired() {
// 				// otherwise this operator type is one that requires a number of packets
// 				bits, remainder, bitsRead = take(11, remainder, bitsRead)
// 				line += bits
// 				bitsInt := utils.BinaryStringToInt(bits)
// 				packet.NumberOfSubpacketsRequired = bitsInt
// 				packet.TotalBitsRead = bitsRead
// 				if DEBUG {
// 					fmt.Printf("11 bits = %v (%v packets to read)\n", bits, bitsInt)
// 				}

// 			}

// 		}
// 		packet.Line = line
// 		// if context.Size() > 0 {
// 		// 	context.Peek().TotalBitsRead += bitsRead
// 		// }

// 		// if root == nil {
// 		// 	root = packet
// 		// 	// last = packet
// 		// 	// } else {
// 		// 	// 	last.Child = packet
// 		// 	// 	packet.Parent = last
// 		// 	// 	last = packet
// 		// }

// 		willBreak := false
// 		if len(remainder) < 8 {
// 			// if len(remainder) > 0 {
// 			// 	if DEBUG {
// 			// 		fmt.Printf(">>>>>>>>>>>")
// 			// 		fmt.Printf(">>>>>>>>>>>")
// 			// 		fmt.Printf(">>>>> Adding remainder bits (length %v) %v\n", len(remainder), remainder)
// 			// 		fmt.Printf(">>>>>>>>>>>")
// 			// 		fmt.Printf(">>>>>>>>>>>")
// 			// 	}
// 			// 	packet.TotalBitsRead += uint64(len(remainder))
// 			// }
// 			willBreak = true
// 			// fmt.Printf("Discarding remainder.... '%v'\n", remainder)
// 			// return remainder
// 		}
// 		// fmt.Printf("original  (%v) is '%v'\n", len(original), original)
// 		// fmt.Printf("remainder (%v) is '%v'\n", len(remainder), remainder)
// 		if willBreak {
// 			break
// 		}

// 	}
// 	return remainder
// }
