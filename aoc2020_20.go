package main

/*

-- Day 20: Jurassic Jigsaw ---
The high-speed train leaves the forest and quickly carries you south. You can even see a desert in the distance! Since you have some spare time, you might as well see if there was anything interesting in the image the Mythical Information Bureau satellite captured.

After decoding the satellite messages, you discover that the data actually contains many small images created by the satellite's camera array. The camera array consists of many cameras; rather than produce a single square image, they produce many smaller square image tiles that need to be reassembled back into a single image.

Each camera in the camera array returns a single monochrome image tile with a random unique ID number. The tiles (your puzzle input) arrived in a random order.

Worse yet, the camera array appears to be malfunctioning: each image tile has been rotated and flipped to a random orientation. Your first task is to reassemble the original image by orienting the tiles so they fit together.

To show how the tiles should be reassembled, each tile's image data includes a border that should line up exactly with its adjacent tiles. All tiles have this border, and the border lines up exactly when the tiles are both oriented correctly. Tiles at the edge of the image also have this border, but the outermost edges won't line up with any other tiles.

For example, suppose you have the following nine tiles:

Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
By rotating, flipping, and rearranging them, you can find a square arrangement that causes all adjacent borders to line up:

#...##.#.. ..###..### #.#.#####.
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
..#.#....# ##.#.#.... ...##.....
For reference, the IDs of the above tiles are:

1951    2311    3079
2729    1427    2473
2971    1489    1171
To check that you've assembled the image correctly, multiply the IDs of the four corner tiles together. If you do this with the assembled tiles from the example above, you get 1951 * 3079 * 2971 * 1171 = 20899048083289.

Assemble the tiles into an image. What do you get if you multiply together the IDs of the four corner tiles?

*/
import (
	"fmt"
	"math"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_20 is the entrypoint
func AOC_2020_20(cli *goutils.CLI) {
	AOC_2020_20_part1_attempt1(cli)
}

func AOC_2020_20_part1_attempt1(cli *goutils.CLI) {
	image := NewImageFromString(DAY_20_DATA)
	fmt.Printf("FindCorners\n")
	corners := image.FindCorners2()
	fmt.Printf("FindCorners: found %v\n", len(corners))
	for _, tile := range corners {
		fmt.Printf("%v\n", tile.TileId)
	}
}

type Image struct {
	Tiles map[string]*Tile
}

func (i *Image) Size() int {
	return len(i.Tiles)
}

func NewImageFromString(data string) *Image {
	lines := strings.Split(data, "\n")
	tiles := make(map[string]*Tile)
	tileData := make([]string, 0)
	tileId := ""
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// end tile
			tile := NewTile(tileId, tileData)
			tiles[tileId] = tile
		} else if strings.Index(line, "Tile") > -1 {
			// start tile
			line = strings.ReplaceAll(line, "Tile", "")
			line = strings.ReplaceAll(line, ":", "")
			line = strings.TrimSpace(line)
			tileId = line
			tileData = make([]string, 0)
		} else {
			// in a tile
			tileData = append(tileData, line)
		}
	}
	i := Image{Tiles: tiles}
	return &i
}

// FindTopLeftCorner looks for a tile which has ONLY 2 relations east and south
func (i *Image) FindCorners2() []*Tile {
	corners := make([]*Tile, 0)
	for _, tile := range i.Tiles {
		count := 0
		for _, candidate := range i.Tiles {
			result := i.AttemptAttach(tile, candidate)
			if result {
				count++
			}
			// two tiles, attempt to attach them
			// walk the candidate all around
			// then rotate the candidate and walkt it all around
			// then flip it one way and walk it all round
			// if it attaches, increment and move on
		}
		if count == 2 {
			corners = append(corners, tile)
		}
	}
	return corners
}

// FindCorners finds any tile which has two relations only.
// func (i *Image) FindCorners() []*Tile {
// 	corners := make([]*Tile, 0)
// 	for _, tile := range i.Tiles {
// 		relations := i.CountRelations(tile)
// 		if relations != 2 {
// 			relations = i.CountRelations(tile.Rotate())
// 		}
// 		if relations != 2 {
// 			relations = i.CountRelations(tile.Rotate())
// 		}
// 		if relations != 2 {
// 			relations = i.CountRelations(tile.Rotate())
// 		}
// 		if relations != 2 {
// 			tile.Rotate()
// 			tile.FlipHorizontal()
// 			relations = i.CountRelations(tile)
// 			if relations != 2 {
// 				relations = i.CountRelations(tile.Rotate())
// 			}
// 			if relations != 2 {
// 				relations = i.CountRelations(tile.Rotate())
// 			}
// 			if relations != 2 {
// 				relations = i.CountRelations(tile.Rotate())
// 			}
// 		}

// 		if relations == 2 {
// 			corners = append(corners, tile)
// 		}

// 	}
// 	return corners
// }

// CountRelations coutns the number different tiles that can attach to this one in this configuration
// func (i *Image) CountRelations(tile *Tile) int {
// 	count := 0
// 	for _, candidate_tile := range i.Tiles {
// 		if candidate_tile == tile {
// 			continue
// 		}
// 		count := candidate_tile.NumberOfConnectionsTo(tile)

// 	}
// 	return count
// }

func (i *Image) GetTile(tileId string) *Tile {
	tile, _ := i.Tiles[tileId]
	return tile
}

const MAX_ROWS = 10
const MAX_COLS = 10

type Tile struct {
	TileId  string
	Matrix  map[string]string
	North   string
	South   string
	East    string
	West    string
	min_col int
	min_row int
}

func NewTile(id string, data []string) *Tile {
	index := 0
	matrix := make(map[string]string)
	for row := 0; row < MAX_ROWS; row++ {
		rowData := data[row]
		for col := 0; col < MAX_COLS; col++ {
			value := rowData[col : col+1]
			key := fmt.Sprintf("%v,%v", col, row)
			matrix[key] = value
			index++
		}
	}

	// build north key
	t := Tile{TileId: id, Matrix: matrix}
	t.Rekey()
	return &t
}

func (t *Tile) Rekey() *Tile {
	north := ""
	south := ""
	east := ""
	west := ""
	for index := 0; index < MAX_COLS; index++ {
		north += t.Get(index, 0)
		south += t.Get(index, MAX_ROWS-1)
		east += t.Get(0, index)
		west += t.Get(MAX_COLS-1, index)
	}
	t.North = north
	t.South = south
	t.West = west
	t.East = east
	return t
}

func (t *Tile) Get(col int, row int) string {
	key := fmt.Sprintf("%v,%v", col, row)
	value, _ := t.Matrix[key]
	return value
}

func (t *Tile) Debug() string {
	data := ""
	for row := 0; row < MAX_ROWS; row++ {
		for col := 0; col < MAX_COLS; col++ {
			value := t.Get(col, row)
			data += value
		}
		data += "\n"
	}
	return strings.TrimSpace(data)
}

func (t *Tile) Rotate() *Tile {
	matrix := make(map[string]string)
	min_col := 1000
	min_row := 1000
	max_col := -1000
	max_row := -1000
	// 90 cw rotation = (x,y) -> (-y, x)
	for col := 0; col < MAX_COLS; col++ {
		for row := 0; row < MAX_ROWS; row++ {
			value := t.Get(col, row)
			new_col := -row
			new_row := col
			min_col = Min(min_col, new_col)
			min_row = Min(min_row, new_row)
			max_col = Max(max_col, new_col)
			max_row = Max(max_row, new_row)
			// original_key := fmt.Sprintf("%v,%v", col, row)
			new_key := fmt.Sprintf("%v,%v", new_col, new_row)
			// fmt.Printf("(%v) cw90 -> (%v)\n", original_key, new_key)
			matrix[new_key] = value
		}
	}

	// now normalise our x,y so we are at 0,0 again
	matrix_normalised := make(map[string]string)
	fmt.Printf("\n\n")
	for key, value := range matrix {
		splits := strings.Split(key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := col + int(math.Abs(float64(min_col)))
		new_row := row + int(math.Abs(float64(min_row)))
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		// fmt.Printf("Rotate: (%v) -> (%v)\n", key, new_key)
		matrix_normalised[new_key] = value
	}
	fmt.Printf("\n\n")

	t.Matrix = matrix_normalised
	t.Rekey()
	return t
}

func (t *Tile) FlipVertical() *Tile {
	matrix_flipped := make(map[string]string)
	for old_key, value := range t.Matrix {
		splits := strings.Split(old_key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_row := MAX_ROWS - row - 1
		new_key := fmt.Sprintf("%v,%v", col, new_row)
		// fmt.Printf("Tile.FlipVertical() %v -> %v\n", old_key, new_key)
		matrix_flipped[new_key] = value
	}
	t.Matrix = matrix_flipped
	t.Rekey()
	return t
}

func (t *Tile) FlipHorizontal() *Tile {
	matrix_flipped := make(map[string]string)
	for old_key, value := range t.Matrix {
		splits := strings.Split(old_key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := MAX_COLS - col - 1
		new_key := fmt.Sprintf("%v,%v", new_col, row)
		matrix_flipped[new_key] = value
	}
	t.Matrix = matrix_flipped
	t.Rekey()
	return t
}

// IsNorthOf indicates if this attaches to the north of the candidate tile
func (t *Tile) IsNorthOf(candidate *Tile) bool {
	return t.South == candidate.North
}

// IsSouthOf indicates if this attaches to the south of the candidate tile
func (t *Tile) IsSouthOf(candidate *Tile) bool {
	return t.North == candidate.South
}

// IsWestOf indicates if this attaches to the west of the candidate tile
func (t *Tile) IsWestOf(candidate *Tile) bool {
	return t.East == candidate.West
}

// IsWestOf indicates if this attaches to the west of the candidate tile
func (t *Tile) IsEastOf(candidate *Tile) bool {
	return t.West == candidate.East
}

// NumberOfConnectionsTo indicates the number of ways this tile can connect to the candidate
func (t *Tile) NumberOfConnections(candidate *Tile) int {
	count := 0
	if t.IsNorthOf(candidate) {
		count++
	}
	if t.IsSouthOf(candidate) {
		count++
	}
	if t.IsEastOf(candidate) {
		count++
	}
	if t.IsWestOf(candidate) {
		count++
	}
	return count
}

/*
 Thoughts

For any tile, the edge as a string is a key.

A connecting tile will have the same key on some side.

A tile shoudl be able to be "Attached" to some side of another tile, rendering it useless.

A tile should be able to count the number of neighbours it *can* have by iterating over all other tiles and manipulating them until it either gets an attachment or does not.

To connect all the tiles together, the total number of connected faces in a 9x9 would be 24


	x x x       2  3  2           = 3x3 = 24

	x x x       3  4  3

	x x x       2  3  2



	x  x  x  x    2  3   3  2     = 4x4  48

	x  x  x  x    3  4   4  3

	x  x  x  x    3  4   4  3

    x  x  x  x    2  3   3  3


	x  x  x  x  x                 = corners 2                 4 corners = 8

	x  x  x  x  x                 = edge pieces 3 per face  12         = 36

	x  x  x  x  x                 all others = 3 edges x3 = 9, 3 each  = 27   = 71

	x  x  x  x  x

	x  x  x  x  x


	pieces
	for length L, total = L^2
	corners = 4
	edges = L - 2 * 4
	blocks/remainder = total - 4(corners) - edges

	relations: corners 2
	edges	: 3
	blocks  : 4




// Number of Edges = 4
// Rotate 90       = 4
// Rotate 180      = 4
// Rotate 270      = 4
// Flip Horiz	   = 2 (the top and bottom)
// Flip Vert       = 2 (left and right)

*/

// AttemptAttach will attempt to attach on any side to a Tile
// The tile will be rotated and flipped in all combinations to attempt to attach on any side
// then it will FlipHorizontal and attempt 4 times, then it will FlipVertical and attempt 4 times
func (i *Image) AttemptAttach(t *Tile, candidate *Tile) bool {
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}

	// back to 0 and FlipVertical
	candidate.Rotate()
	candidate.FlipVertical()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}

	// back to 0 and FlipHorizontal
	candidate.Rotate()
	candidate.FlipHorizontal()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.NumberOfConnections(candidate) > 0 {
		return true
	}
	return false
}
