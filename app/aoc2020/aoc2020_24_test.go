package aoc2020

import (
	"fmt"
	"testing"
)

const DAY_24_TEST_INPUT = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func Test_AOC2020_24_ParseAddress(t *testing.T) {
	input := "neeswseenwwswnwswswnw"
	grid := NewHexGrid()
	output := grid.ParseAddress(input)
	// "ne e sw se e nw w sw nw sw sw nw"
	if len(output) != 12 {
		t.Errorf("Ouput is incorrect length.")
	}

	if output[0] != "ne" {
		t.Errorf("ouput[0] shoud be ne, '%v'\n", output[0])
	}

	if output[1] != "e" {
		t.Errorf("ouput[1] shoud be e, '%v'\n", output[1])
	}

	if output[2] != "sw" {
		t.Errorf("ouput[2] shoud be sw, '%v'\n", output[2])
	}

	if output[3] != "se" {
		t.Errorf("ouput[3] shoud be se, '%v'\n", output[3])
	}

	if output[4] != "e" {
		t.Errorf("ouput[4] shoud be e, '%v'\n", output[4])
	}

	if output[5] != "nw" {
		t.Errorf("ouput[5] shoud be nw, '%v'\n", output[5])
	}

	if output[6] != "w" {
		t.Errorf("ouput[6] shoud be w, '%v'\n", output[6])
	}

	if output[7] != "sw" {
		t.Errorf("ouput[7] shoud be sw, '%v'\n", output[7])
	}

	if output[8] != "nw" {
		t.Errorf("ouput[9] shoud be nw, '%v'\n", output[8])
	}

	if output[9] != "sw" {
		t.Errorf("ouput[10] shoud be sw, '%v'\n", output[9])
	}

	if output[10] != "sw" {
		t.Errorf("ouput[10] shoud be sw, '%v'\n", output[10])
	}

	if output[11] != "nw" {
		t.Errorf("ouput[11] shoud be nw, '%v'\n", output[11])
	}

}

func Test_AOC2020_24_AddressParsing(t *testing.T) {
	// grid := NewHexGrid()
	// lines := strings.Split(DAY_24_TEST_INPUT, "\n")
	verifyCoordinates("seswneswswsenwwnwse", -1.5, -3, t)

	// neeenesenwnwwswnenewnwwsewnenwseswesw
	// seswneswswsenwwnwse
	// nwnwneseeswswnenewneswwnewseswneseene
	// swweswneswnenwsewnwneneseenw
	// eesenwseswswnenwswnwnwsewwnwsene
	// sewnenenenesenwsewnenwwwse
	// wenwwweseeeweswwwnwwe
	// wsweesenenewnwwnwsenewsenwwsesesenwne
	// neeswseenwwswnwswswnw
	// nenwswwsewswnenenewsenwsenwnesesenew
	// enewnwewneswsewnwswenweswnenwsenwsw
	// sweneswneswneneenwnewenewwneswswnese
	// swwesenesewenwneswnwwneseswwne
	// enesenwswwswneneswsenwnewswseenwsese
	// wnwnesenesenenwwnenwsewesewsesesew
	// nenewswnwewswnenesenwnesewesw
	// eneswnwswnwsenenwnwnwwseeswneewsenese
	// neswnwewnwnwseenwseesewsenwsweewe
	// wseweeenwnesenwwwswnew
}

func verifyCoordinates(address string, expected_x float64, expected_y int, t *testing.T) {
	grid := NewHexGrid()
	actual_x, actual_y := grid.CoordinatesCreateHexesAlongTheWay(address)
	if actual_x != expected_x || actual_y != expected_y {
		t.Errorf("VerifyCoordinates(%v) expect (%v,%v) got (%v,%v)\n", address, expected_x, expected_y, actual_x, actual_y)
		t.Errorf("%v\n", grid.ParseAddress(address))
	}
}

// for _, address := range lines {
// 		// parsed := grid.ParseAddress(address)
// 		x, y := grid.Coordinates(address)

// 	}
// }

func Test_AOC2020_24_Part1_Test(t *testing.T) {
	grid := NewHexGrid()
	grid.PlayPart1(DAY_24_TEST_INPUT)
	// fmt.Printf("BlackCount: %v\n", grid.BlackCount())
	grid.Render(0, "test_day_0.png")
	for day := 1; day <= 100; day++ {
		grid.PlayPart2(day)
		filename := fmt.Sprintf("test_day_%v.png", day)
		grid.Render(day, filename)
	}
}

func Test_AOC2020_24_Part2_Real(t *testing.T) {
	grid := NewHexGrid()
	grid.PlayPart1(DAY_24_INPUT)
	width, height := grid.Dimensions()
	fmt.Printf("BlackCount: %v, width: %v, height :%v\n", grid.BlackCount(), width, height)
	grid.Render(0, "real_day_0.png")
	for day := 1; day <= 100; day++ {
		grid.PlayPart2(day)
		filename := fmt.Sprintf("real_day_%v.png", day)
		grid.Render(day, filename)
	}
}
func Test_AOC2020_24_Part2(t *testing.T) {
	grid := NewHexGrid()
	grid.PlayPart1(DAY_24_TEST_INPUT)
	fmt.Printf("Day 0, Count: %v\n", grid.BlackCount())
	// grid.Render("test.png")
	fmt.Printf("BlackCount: %v\n", grid.BlackCount())
	for day := 1; day <= 3; day++ {
		stayWhite, stayBlack, turnWhite, turnBlack := grid.PlayPart2(day)
		fmt.Printf("Day %v, StayWhite: %v, StayBlack %v, Turn White %v, Turn Black %v, Count Blackt: %v\n", day, stayWhite, stayBlack, turnWhite, turnBlack, grid.BlackCount())
	}
}

func Test_AOC2020_24_Part2_Animation(t *testing.T) {
	grid := NewHexGrid()
	grid.RenderAnimation(DAY_24_INPUT, "animation.gif", 100, 25)
}
