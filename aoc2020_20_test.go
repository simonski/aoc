package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_AOC2020_20_TestX(t *testing.T) {
	image := NewImageFromString(DAY_20_DATA)
	if image.Size() != 144 {
		t.Errorf("Day 20 Part 1: Test data should have 144 tiles, was %v.\n", image.Size())
	}
}

func Test_AOC2020_20_Test1(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	if image.Size() != 9 {
		t.Errorf("Day 20 Part 1: Test data should have 6 tiles, was %v.\n", image.Size())
	}

	tile2311 := image.GetTile("2311")
	fmt.Printf("\nRotate(0)\n")
	fmt.Printf("%v\n", tile2311.Debug(false))

	tile2311_cw90 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(90)\n")
	fmt.Printf("%v\n", tile2311_cw90.Debug(false))

	tile2311_cw180 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(180)\n")
	fmt.Printf("%v\n", tile2311_cw180.Debug(false))

	tile2311_cw270 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(270)\n")
	fmt.Printf("%v\n", tile2311_cw270.Debug(false))

	tile2311_cw360 := image.GetTile("2311").Rotate()
	fmt.Printf("\nRotate(360)\n")
	fmt.Printf("%v\n", tile2311_cw360.Debug(false))

}

func Test_AOC2020_20_TestFlipVertical(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	tile2311 := image.GetTile("2311")

	expected_original := `..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`

	expected_flipped := `..###..###
###...#.#.
..#....#..
.#.#.#..##
##...#.###
##.##.###.
####.#...#
#...##..#.
##..#.....
..##.#..#.`

	// 	expected_flipped := `..###..###
	// ###...#.#.
	// ..#....#..
	// .#.#.#..##
	// ##...#.###
	// ##.##.###.
	// ####.#...#
	// #...##..#.
	// ##..#.....
	// ..##.#..#.`

	actual_original := tile2311.Debug(false)
	if actual_original != expected_original {
		t.Errorf("FlipVertical: originals do not match\n.")
		t.Errorf("Expected Original\n\n`%v`\n\nActual Original\n\n`%v`\n", expected_original, actual_original)
	}

	tile2311.FlipVertical()
	actual_flipped := tile2311.Debug(false)
	if actual_flipped != expected_flipped {
		t.Errorf("FlipVertical: flipped do not match\n.")
		t.Errorf("Pre-flip\n\n%v\n\n", actual_original)
		t.Errorf("Expected\n\n%v\n\n", expected_flipped)
		t.Errorf("Actual\n\n%v\n\n", actual_flipped)
	}

}

func Test_AOC2020_20_TestFlipHorizontal(t *testing.T) {
	image := NewImageFromString(DAY_20_TEST_DATA)
	tile2311 := image.GetTile("2311")

	expected_original := `..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`

	expected_flipped := `.#..#.##..
.....#..##
.#..##...#
#...#.####
.###.##.##
###.#...##
##..#.#.#.
..#....#..
.#.#...###
###..###..`

	actual_original := tile2311.Debug(false)
	if actual_original != expected_original {
		t.Errorf("FlipHorizontal: originals do not match\n.")
	}

	tile2311.FlipHorizontal()
	actual_flipped := tile2311.Debug(false)
	if actual_flipped != expected_flipped {
		t.Errorf("FlipHorizontal: flipped do not match\n.")
		t.Errorf("Pre-flip\n\n%v\n\n", actual_original)
		t.Errorf("Expected\n\n%v\n\n", expected_flipped)
		t.Errorf("Actual\n\n%v\n\n", actual_flipped)
	}

}

func Test_AOC2020_20_RotateAndFlip(t *testing.T) {
	data := `Tile 1319:
#.########
##.#.#.#..
.....#..##
.#.....#..
.........#
#.........
#..##...#.
#.#.#...##
.#....#...
....##..#.
`

	expectedNorth := "#.########"
	expectedSouth := "....##..#."
	expectedEast := "#.#.#..#.."
	expectedWest := "##...###.."

	image := NewImageFromString(data)
	tile := image.GetTile("1319")
	if tile == nil {
		t.Errorf("There should be a single tile, 1319\n")
	}

	if tile.NorthEdge != expectedNorth {
		t.Errorf("NorthEdge should be %v, was %v\n", expectedNorth, tile.NorthEdge)
	}

	if tile.SouthEdge != expectedSouth {
		t.Errorf("SouthEdge should be %v, was %v\n", expectedSouth, tile.SouthEdge)
	}

	if tile.EastEdge != expectedEast {
		t.Errorf("EastEdge should be %v, was %v\n", expectedEast, tile.EastEdge)
	}

	if tile.WestEdge != expectedWest {
		t.Errorf("WestEdge should be %v, was %v\n", expectedWest, tile.WestEdge)
	}

	tile.FlipHorizontal()
	expectedNorth = "########.#"
	expectedSouth = ".#..##...."
	expectedEast = "##...###.."
	expectedWest = "#.#.#..#.."
	if tile.NorthEdge != expectedNorth {
		t.Errorf("HorizontallyFlipped NorthEdge should be %v, was %v\n", expectedNorth, tile.NorthEdge)
	}

	if tile.SouthEdge != expectedSouth {
		t.Errorf("HorizontallyFlipped SouthEdge should be %v, was %v\n", expectedSouth, tile.SouthEdge)
	}

	if tile.EastEdge != expectedEast {
		t.Errorf("HorizontallyFlipped EastEdge should be %v, was %v\n", expectedEast, tile.EastEdge)
	}

	if tile.WestEdge != expectedWest {
		t.Errorf("HorizontallyFlipped WestEdge should be %v, was %v\n", expectedWest, tile.WestEdge)
	}

	tile.FlipHorizontal()
	expectedNorth = "#.########"
	expectedSouth = "....##..#."
	expectedEast = "#.#.#..#.."
	expectedWest = "##...###.."
	if tile.NorthEdge != expectedNorth {
		t.Errorf("HorizontallyFlipped NorthEdge should be %v, was %v\n", expectedNorth, tile.NorthEdge)
	}

	if tile.SouthEdge != expectedSouth {
		t.Errorf("HorizontallyFlipped SouthEdge should be %v, was %v\n", expectedSouth, tile.SouthEdge)
	}

	if tile.EastEdge != expectedEast {
		t.Errorf("HorizontallyFlipped EastEdge should be %v, was %v\n", expectedEast, tile.EastEdge)
	}

	if tile.WestEdge != expectedWest {
		t.Errorf("HorizontallyFlipped WestEdge should be %v, was %v\n", expectedWest, tile.WestEdge)
	}

	/*
	   #.########
	   ##.#.#.#..
	   .....#..##
	   .#.....#..
	   .........#
	   #.........
	   #..##...#.
	   #.#.#...##
	   .#....#...
	   ....##..#.
	*/

	tile.Rotate()
	expectedNorth = "..###...##"
	expectedSouth = "..#..#.#.#"
	expectedEast = "#.########"
	expectedWest = "....##..#."
	if tile.NorthEdge != expectedNorth {
		t.Errorf("R90 NorthEdge should be %v, was %v\n", expectedNorth, tile.NorthEdge)
	}

	if tile.SouthEdge != expectedSouth {
		t.Errorf("R90 SouthEdge should be %v, was %v\n", expectedSouth, tile.SouthEdge)
	}

	if tile.EastEdge != expectedEast {
		t.Errorf("R90 EastEdge should be %v, was %v\n", expectedEast, tile.EastEdge)
	}

	if tile.WestEdge != expectedWest {
		t.Errorf("R90 WestEdge should be %v, was %v\n", expectedWest, tile.WestEdge)
	}

}

/*
1951    2311    3079
2729    1427    2473
2971    1489    1171
*/

func Test_AOC2020_20_TestArrangeX(t *testing.T) {
	testImage := NewImageFromString(DAY_20_TEST_DATA)
	testImage.Arrange(false)
	hideEdges := false
	showBorders := true
	showCoordinates := false
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf(testImage.Debug(hideEdges, showBorders, showCoordinates))
	fmt.Printf("\n")
	fmt.Printf("\n")

	hideEdges = true
	showBorders = true
	showCoordinates = false
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf(testImage.Debug(hideEdges, showBorders, showCoordinates))
	fmt.Printf("\n")
	fmt.Printf("\n")

	hideEdges = true
	showBorders = false
	showCoordinates = false
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf(testImage.Debug(hideEdges, showBorders, showCoordinates))
	fmt.Printf("\n")
	fmt.Printf("\n")

	// result, errors := ValidateMatrix(testDataFlippedVertically)
	// if errors {
	// 	t.Errorf(result)
	// }

	// testDataFlippedVertically := FlipMatrixVertical(testDataOriginal)

	expectedWithSpacesAndBorder := `#...##.#.. ..###..### #.#.#####.
..#.#..#.# ###...#.#. .#..######
.###....#. ..#....#.. ..#.......
###.##.##. .#.#.#..## ######....
.###.##### ##...#.### ####.#..#.
.##.#....# ##.##.###. .#...#.##.
#...###### ####.#...# #.#####.##
.....#..## #...##..#. ..#.###...
#.####...# ##..#..... ..#.......
#.##...##. ..##.#..#. ..#.###...

#.##...##. ..##.#..#. ..#.###...
##..#.##.. ..#..###.# ##.##....#
##.####... .#.####.#. ..#.###..#
####.#.#.. ...#.##### ###.#..###
.#.####... ...##..##. .######.##
.##..##.#. ....#...## #.#.#.#...
....#..#.# #.#.#.##.# #.###.###.
..#.#..... .#.##.#..# #.###.##..
####.#.... .#..#.##.. .######...
...#.#.#.# ###.##.#.. .##...####

...#.#.#.# ###.##.#.. .##...####
..#.#.###. ..##.##.## #..#.##..#
..####.### ##.#...##. .#.#..#.##
#..#.#..#. ...#.#.#.. .####.###.
.#..####.# #..#.#.#.# ####.###..
.#####..## #####...#. .##....##.
##.##..#.. ..#...#... .####...#.
#.#.###... .##..##... .####.##.#
#...###... ..##...#.. ...#..####
..#.#....# ##.#.#.... ...##.....`

	// actualWithSpacesAndBorder := DebugMatrixFully(testDataFlippedVertically, false, " ")
	hideEdges = false
	showBorders = true
	showCoordinates = false
	actualWithSpacesAndBorder := testImage.Debug(hideEdges, showBorders, showCoordinates)
	if strings.TrimSpace(actualWithSpacesAndBorder) != strings.TrimSpace(expectedWithSpacesAndBorder) {
		t.Errorf("Full matrix including edges and including spaces does not match.\nExpected:\n`%v`\n\nActual\n\n`%v`\n", expectedWithSpacesAndBorder, actualWithSpacesAndBorder)
		splitsActual := strings.Split(actualWithSpacesAndBorder, "\n")
		splitsExpected := strings.Split(expectedWithSpacesAndBorder, "\n")
		for index := 0; index < len(splitsActual); index++ {
			lineActual := splitsActual[0]
			lineExpected := splitsExpected[0]
			if strings.TrimSpace(lineActual) != strings.TrimSpace(lineExpected) {
				t.Errorf("stripped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else if lineActual != lineExpected {
				t.Errorf("non-striped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else {
				fmt.Printf("Line %v matches stripped and non-stripped ( '%v' == '%v').\n", index, lineExpected, lineActual)

			}
		}

	} else {
		fmt.Printf("Full matrix including edges and including spaces matches.\n")
	}

	expectedWithSpaces := `.#.#..#. ##...#.# #..#####
###....# .#....#. .#......
##.##.## #.#.#..# #####...
###.#### #...#.## ###.#..#
##.#.... #.##.### #...#.##
...##### ###.#... .#####.#
....#..# ...##..# .#.###..
.####... #..#.... .#......

#..#.##. .#..###. #.##....
#.####.. #.####.# .#.###..
###.#.#. ..#.#### ##.#..##
#.####.. ..##..## ######.#
##..##.# ...#...# .#.#.#..
...#..#. .#.#.##. .###.###
.#.#.... #.##.#.. .###.##.
###.#... #..#.##. ######..

.#.#.### .##.##.# ..#.##..
.####.## #.#...## #.#..#.#
..#.#..# ..#.#.#. ####.###
#..####. ..#.#.#. ###.###.
#####..# ####...# ##....##
#.##..#. .#...#.. ####...#
.#.###.. ##..##.. ####.##.
...###.. .##...#. ..#..###`

	hideEdges = true
	showBorders = true
	showCoordinates = false
	actualWithSpaces := testImage.Debug(hideEdges, showBorders, showCoordinates)
	// if actualWithSpaces != expectedWithSpaces {
	// 	t.Errorf("Stripped content is not the same as expected.\nExpected:\n%v\n\nActual\n\n%v\n", expectedWithSpaces, actualWithSpaces)
	// }

	if strings.TrimSpace(actualWithSpaces) != strings.TrimSpace(expectedWithSpaces) {
		t.Errorf("Full matrix excluding edges and including borders does not match.\nExpected:\n`%v`\n\nActual\n\n`%v`\n", expectedWithSpaces, actualWithSpaces)
		splitsActual := strings.Split(actualWithSpaces, "\n")
		splitsExpected := strings.Split(expectedWithSpaces, "\n")
		for index := 0; index < len(splitsActual); index++ {
			lineActual := splitsActual[0]
			lineExpected := splitsExpected[0]
			if strings.TrimSpace(lineActual) != strings.TrimSpace(lineExpected) {
				t.Errorf("stripped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else if lineActual != lineExpected {
				t.Errorf("non-striped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else {
				fmt.Printf("Line %v matches stripped and non-stripped ( '%v' == '%v').\n", index, lineExpected, lineActual)

			}
		}

	} else {
		fmt.Printf("Full matrix excluding edges and including borders matches.\n")
	}

	expectedWithoutBordersOrSpaces := `.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`

	hideEdges = true
	showBorders = false
	showCoordinates = false
	actualWithoutBordersOrSpaces := testImage.Debug(hideEdges, showBorders, showCoordinates)

	if strings.TrimSpace(actualWithoutBordersOrSpaces) != strings.TrimSpace(expectedWithoutBordersOrSpaces) {
		t.Errorf("Full matrix excluding borders and excluding spaces does not match.\nExpected:\n`%v`\n\nActual\n\n`%v`\n", expectedWithoutBordersOrSpaces, actualWithoutBordersOrSpaces)
		splitsActual := strings.Split(actualWithoutBordersOrSpaces, "\n")
		splitsExpected := strings.Split(expectedWithoutBordersOrSpaces, "\n")
		for index := 0; index < len(splitsActual); index++ {
			lineActual := splitsActual[0]
			lineExpected := splitsExpected[0]
			if strings.TrimSpace(lineActual) != strings.TrimSpace(lineExpected) {
				t.Errorf("stripped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else if lineActual != lineExpected {
				t.Errorf("non-striped Line %v does not match\nExpected `%v`\nActual   `%v`\n\n", index, lineExpected, lineActual)
			} else {
				fmt.Printf("Line %v matches stripped and non-stripped ( '%v' == '%v').\n", index, lineExpected, lineActual)

			}
		}

	} else {
		fmt.Printf("Full matrix excluding borders and spaces matches.\n")
	}

}

func Test_AOC2020_20_TestTileOutputMatchesImageOuput(t *testing.T) {
	testImage := NewImageFromString(DAY_20_TEST_DATA)
	testImage.Arrange(false)
	hideEdges := true
	showBorders := false
	showCoordinates := false

	imageOutput := testImage.Debug(hideEdges, showBorders, showCoordinates)

	tile := NewTileFromString(imageOutput)
	tileOutput := strings.TrimSpace(tile.Debug(false))
	if imageOutput != tileOutput {
		t.Errorf("Tile output must match Image output\nImage\n%v\n\nTile\n%v\n\n", imageOutput, tileOutput)
	} else {
		fmt.Printf("Tile output matches Image output.\n")
	}

	fmt.Printf("R0\n%v\n", tileOutput)

	tile.Rotate()
	fmt.Printf("R90\n%v\n", tile.Debug(false))

	tile.Rotate()
	fmt.Printf("R180\n%v\n", tile.Debug(false))

	tile.Rotate()
	fmt.Printf("R270\n%v\n", tile.Debug(false))

}

// func Test_AOC2020_20_2_Arrange_RealData(t *testing.T) {
// 	fmt.Printf("\nTest_AOC2020_20_2_Arrange_RealData\n\n")

// 	image := NewImageFromString(DAY_20_DATA)
// 	image.Arrange(false)
// 	hideEdges := true
// 	showBorders := false
// 	showCoordinates := false
// 	output := image.Debug(hideEdges, showBorders, showCoordinates)

// 	fmt.Printf("Full Data Data\n")
// 	fmt.Printf(output)
// 	fmt.Printf("\n\n")

// }

func Test_AOC2020_20_TestNessieRegex(t *testing.T) {

	// regex, err := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	// if err != nil {
	// 	t.Errorf("Problem with regex compiling.")
	// }

	// expression, _ := regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	// match := expression.FindStringSubmatch(ecl)
	// // fmt.Printf("ecl: %v, %v entries, %v\n", ecl, len(match), match)
	// if len(match) != 2 {
	// 	return false
	// }

	// 	                 #
	// #    ##    ##    ###
	//  #  #  #  #  #  #
	// line1 := "	                 # "
	// line2 := " #    ##    ##    ###"
	// line3 := "  #  #  #  #  #  #"

	pattern := " # "
	value := " # "
	expression, _ := regexp.Compile(pattern)
	if !expression.MatchString(value) {
		t.Errorf("Regex '%v' fails on value '%v'\n.", pattern, value)

	}

	pattern = "(.)# "
	expression, _ = regexp.Compile(pattern)
	if !expression.MatchString(value) {
		t.Errorf("Regex '%v' fails on value '%v'\n.", pattern, value)

	}

	pattern = "(..)# "
	expression, _ = regexp.Compile(pattern)
	if expression.MatchString(value) {
		t.Errorf("Regex '%v' fails on value '%v'\n.", pattern, value)
	}

	line1Expr, err := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	if err != nil {
		t.Errorf("Problem with regex compiling.")
	}
	if line1Expr.MatchString(value) {
		t.Errorf("Should nt match.")
	}

	value1 := "                  # "
	if !line1Expr.MatchString(value1) {
		t.Errorf("Should match.")
	}
	value1 = "#  #   x         a#b"
	if !line1Expr.MatchString(value1) {
		t.Errorf("Should match.")
	}

	match := line1Expr.FindStringSubmatch(value1)
	if len(match) != 2 {
		t.Errorf("regex is    '%v'", DAY_20_NESSIE_REGEX_LINE1)
		t.Errorf("line is     '%v'", value1)
		t.Errorf("match[0] is '%v'", match[0])
		t.Errorf("match[1] is '%v'", match[1])
		t.Errorf("match[2] is '%v'", match[2])
	}

	// expression, _ := regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	// match := expression.FindStringSubmatch(ecl)
	// // fmt.Printf("ecl: %v, %v entries, %v\n", ecl, len(match), match)
	// if len(match) != 2 {
	// 	return false
	// }

	// 	                 #
	// #    ##    ##    ###
	//  #  #  #  #  #  #
	// line1 := "	                 # "
	// line2 := " #    ##    ##    ###"
	// line3 := "  #  #  #  #  #  #"

	// if !regex.MatchString(line1) {
	// 	t.Errorf("'%v' should match regex '%v'\n", line1, DAY_20_NESSIE_REGEX_LINE1)
	// }
	// if regex.MatchString(line2) {
	// 	t.Errorf("'%v' incorrectly matches regex '%v'\n", line1, DAY_20_NESSIE_REGEX_LINE1)
	// }
	// if regex.MatchString(line3) {
	// 	t.Errorf("'%v' incorrectly matches regex '%v'\n", line1, DAY_20_NESSIE_REGEX_LINE1)
	// }

	// if !regex.MatchString(line1) {
	// 	t.Errorf("%v does not match regex %v\n", line2, DAY_20_NESSIE_REGEX_LINE2)
	// }

	// if !regex.MatchString(line1) {
	// 	t.Errorf("%v does not match regex %v\n", line3, DAY_20_NESSIE_REGEX_LINE3)
	// }

	// testImage := NewImageFromString(DAY_20_TEST_DATA)
	// testImage.Arrange(false)
	// hideEdges := true
	// showBorders := false
	// showCoordinates := false

	// imageOutput := testImage.Debug(hideEdges, showBorders, showCoordinates)

	// tile := NewTileFromString(imageOutput)
	// tileOutput := strings.TrimSpace(tile.Debug(false))
	// if imageOutput != tileOutput {
	// 	t.Errorf("Tile output must match Image output\nImage\n%v\n\nTile\n%v\n\n", imageOutput, tileOutput)
	// } else {
	// 	fmt.Printf("Tile output matches Image output.\n")
	// }

	// fmt.Printf("R0\n%v\n", tileOutput)

	// tile.Rotate()
	// fmt.Printf("R90\n%v\n", tile.Debug(false))

	// tile.Rotate()
	// fmt.Printf("R180\n%v\n", tile.Debug(false))

	// tile.Rotate()
	// fmt.Printf("R270\n%v\n", tile.Debug(false))

}

func Test_AOC2020_20_TestTileCanRotate(t *testing.T) {
	// 7 monsters, 2368 choppy - too high
	// 20 monsters, FlipVertical only - 2173

	// testImage := NewImageFromString(DAY_20_TEST_DATA)
	testImage := NewImageFromString(DAY_20_DATA)
	testImage.Arrange(false)
	hideEdges := true
	showBorders := false
	showCoordinates := false
	imageOutput := testImage.Debug(hideEdges, showBorders, showCoordinates)
	fmt.Printf("%v\n", imageOutput)

	// hideEdges = true
	// showBorders = false
	// showCoordinates = false
	// imageOutputX := testImage.Debug(hideEdges, showBorders, showCoordinates)
	// fmt.Printf("\n\n%v\n\n", imageOutputX)

	tile := NewTileFromString(imageOutput)
	tile.FlipVertical()
	// tile.Rotate()
	// tile.Rotate()
	// tile.Rotate()

	// 2 rotates for real data
	// if tile.RotateUntilAllRegexMatch() {
	// 	t.Errorf("Rotated successfully.")
	// } else {
	// 	t.Errorf("Failed to rotate successfully.")
	// 	os.Exit(1)
	// }

	// tile.Rotate()
	// tile.Rotate()
	// tile.Rotate()

	// if tile.RotateUntilAllRegexMatch() {
	// 	t.Errorf("Rotated successfully.")
	// } else {
	// 	t.Errorf("Failed to rotate successfully.")
	// 	os.Exit(1)
	// }

	monsters := tile.FindSeaMonsters()

	// imageOutputY := tile.ShowSeaMonsters()
	// fmt.Printf("\n\n%v\n\n", imageOutputY)

	// t.Errorf("Line1 regex '%v'\n", DAY_20_NESSIE_REGEX_LINE1)
	// t.Errorf("Line2 regex '%v'\n", DAY_20_NESSIE_REGEX_LINE2)
	// t.Errorf("Line3 regex '%v'\n\n", DAY_20_NESSIE_REGEX_LINE3)

	// for monsterKey, monsterLines := range monsters {
	// 	t.Errorf("monster (%v)\n", monsterKey)
	// 	for index, line := range monsterLines {
	// 		for matchIndex, l := range line {
	// 			t.Errorf("Line[%v] regexMatch[%v] for monster: %v\n", index+1, matchIndex+1, l)
	// 		}
	// 	}
	// 	t.Errorf("\n")
	// }

	imageOutput = strings.ReplaceAll(imageOutput, "\n", "")
	imageOutput = strings.ReplaceAll(imageOutput, " ", "")
	totalLength := len(imageOutput)
	lengthLessHashes := len(strings.ReplaceAll(imageOutput, "#", ""))
	totalHashes := totalLength - lengthLessHashes
	totalChoppyWater := totalHashes - (len(monsters) * 15)
	t.Errorf("Total length %v, monsters %v, total hashes %v, total choppy water %v\n", totalLength, len(monsters), totalHashes, totalChoppyWater)

}

func Test_AOC2020_20_TestFindNessieWithRegex(t *testing.T) {

	testImage := NewImageFromString(DAY_20_TEST_DATA)
	testImage.Arrange(false)
	hideEdges := true
	showBorders := false
	showCoordinates := false
	imageOutput := testImage.Debug(hideEdges, showBorders, showCoordinates)
	tile := NewTileFromString(imageOutput)
	tileOutputR0 := strings.TrimSpace(tile.Debug(false))
	tileOutputR90 := strings.TrimSpace(tile.Rotate().Debug(false))
	tileOutputR180 := strings.TrimSpace(tile.Rotate().Debug(false))
	tileOutputR270 := strings.TrimSpace(tile.Rotate().Debug(false))

	regexLine1, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	regexLine2, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE2)
	regexLine3, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE3)

	matchL1_R0 := regexLine1.FindStringSubmatch(tileOutputR0)
	matchL1_R90 := regexLine1.FindStringSubmatch(tileOutputR90)
	matchL1_R180 := regexLine1.FindStringSubmatch(tileOutputR180)
	matchL1_R270 := regexLine1.FindStringSubmatch(tileOutputR270)

	matchL2_R0 := regexLine2.FindStringSubmatch(tileOutputR0)
	matchL2_R90 := regexLine2.FindStringSubmatch(tileOutputR90)
	matchL2_R180 := regexLine2.FindStringSubmatch(tileOutputR180)
	matchL2_R270 := regexLine2.FindStringSubmatch(tileOutputR270)

	matchL3_R0 := regexLine3.FindStringSubmatch(tileOutputR90)
	matchL3_R90 := regexLine3.FindStringSubmatch(tileOutputR180)
	matchL3_R180 := regexLine3.FindStringSubmatch(tileOutputR180)
	matchL3_R270 := regexLine3.FindStringSubmatch(tileOutputR270)

	t.Errorf("\n")
	t.Errorf("R0 match=%v,%v,%v\n", len(matchL1_R0), len(matchL2_R0), len(matchL3_R0))
	t.Errorf("R90 match=%v,%v,%v\n", len(matchL1_R90), len(matchL2_R90), len(matchL3_R90))
	t.Errorf("R180 match=%v,%v,%v\n", len(matchL1_R180), len(matchL2_R180), len(matchL3_R180))
	t.Errorf("R270 match=%v,%v,%v\n", len(matchL1_R270), len(matchL2_R270), len(matchL3_R270))

	// flip the tile round until it matches

	// if len(match)
	// 	t.Errorf("regex is    '%v'", DAY_20_NESSIE_REGEX_LINE1)
	// 	t.Errorf("line is     '%v'", value1)
	// 	t.Errorf("match[0] is '%v'", match[0])
	// 	t.Errorf("match[1] is '%v'", match[1])
	// 	t.Errorf("match[2] is '%v'", match[2])
	// }

	// regex, err := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	// if err != nil {
	// 	t.Errorf("Problem with regex compiling.")
	// }
}
