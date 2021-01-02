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
	AOC_2020_20_part2_attempt1(cli)
}

func AOC_2020_20_part1_attempt1(cli *goutils.CLI) {
	image := NewImageFromString(DAY_20_DATA)
	fmt.Printf("FindCorners\n")
	data := image.Arrange(false)
	DebugMatrix(data, true)

	size := int(math.Sqrt(float64(len(data))))
	key1 := fmt.Sprintf("%v,%v", 0, 0)
	key2 := fmt.Sprintf("%v,%v", size-1, 0)
	key3 := fmt.Sprintf("%v,%v", 0, size-1)
	key4 := fmt.Sprintf("%v,%v", size-1, size-1)
	corner1 := data[key1]
	corner2 := data[key2]
	corner3 := data[key3]
	corner4 := data[key4]
	fmt.Printf("Day 20.1: Corners are %v * %v * %v * %v = ?\n", corner1.TileId, corner2.TileId, corner3.TileId, corner4.TileId)
}

func AOC_2020_20_part2_attempt1(cli *goutils.CLI) {
	image := NewImageFromString(DAY_20_DATA)
	image.Arrange(false)
	hideEdges := false
	showBorders := true
	showCoordinates := false
	matrix := image.Debug(hideEdges, showBorders, showCoordinates)
	fmt.Printf(matrix)

	// put all tiles together
	// strip the edges of each
	// create a single image
	// apply the regex
	//

	// image := NewImageFromString(DAY_20_DATA)
	// corners := image.FindCorners()
	// image.ArrangeUsingCorners(corners)
	// 3221
	// 1447
	// 1873
	// 2029

	// fmt.Printf("FindCorners: found %v\n", len(corners))
	// for _, tile := range corners {
	// 	fmt.Printf("%v\n", tile.TileId)
	// }
}

// const TILE_WIDTH = 10

func NormaliseTileMap(source map[string]*Tile) map[string]*Tile {
	min_col := 1000
	min_row := 1000
	max_col := -1000
	max_row := -1000
	// 90 cw rotation = (x,y) -> (-y, x)
	for key, _ := range source {
		splits := strings.Split(key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		min_col = Min(min_col, col)
		min_row = Min(min_row, row)
		max_col = Max(max_col, col)
		max_row = Max(max_row, row)
	}

	// now normalise our x,y so we are at 0,0 again
	matrix_normalised := make(map[string]*Tile)
	for key, value := range source {
		splits := strings.Split(key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := col + int(math.Abs(float64(min_col)))
		new_row := row + int(math.Abs(float64(min_row)))
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		// fmt.Printf("Rotate: (%v) -> (%v)\n", key, new_key)
		matrix_normalised[new_key] = value
	}
	return matrix_normalised
}

// Rotates the tile 90 degrees clockwise, removing any relationships this
// tile has
func RotateMatrix(input map[string]*Tile) map[string]*Tile {
	output := make(map[string]*Tile)
	min_col := 1000
	min_row := 1000
	max_col := -1000
	max_row := -1000
	// 90 cw rotation = (x,y) -> (-y, x)
	for key, tile := range input {
		col, _ := strconv.Atoi(strings.Split(key, ",")[0])
		row, _ := strconv.Atoi(strings.Split(key, ",")[1])
		new_col := -row
		new_row := col
		min_col = Min(min_col, new_col)
		min_row = Min(min_row, new_row)
		max_col = Max(max_col, new_col)
		max_row = Max(max_row, new_row)
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		output[new_key] = tile
	}

	// now normalise our x,y so we are at 0,0 again
	normalised := make(map[string]*Tile)
	// fmt.Printf("\n\n")
	for key, value := range output {
		splits := strings.Split(key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := col + int(math.Abs(float64(min_col)))
		new_row := row + int(math.Abs(float64(min_row)))
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		// fmt.Printf("Rotate: (%v) -> (%v)\n", key, new_key)
		normalised[new_key] = value
	}

	return normalised
}

// Rotates the tile 90 degrees clockwise, removing any relationships this
// tile has
func FlipMatrixVertical(input map[string]*Tile) map[string]*Tile {
	output := make(map[string]*Tile)
	min_col := 1000
	min_row := 1000
	max_col := -1000
	max_row := -1000
	// vertical flip is x, y -> x, -y and normalise
	for key, tile := range input {
		col, _ := strconv.Atoi(strings.Split(key, ",")[0])
		row, _ := strconv.Atoi(strings.Split(key, ",")[1])
		new_row := -row
		new_col := col
		min_col = Min(min_col, new_col)
		min_row = Min(min_row, new_row)
		max_col = Max(max_col, new_col)
		max_row = Max(max_row, new_row)
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		output[new_key] = tile
		tile.SetCanRotateOrFlip(true)
		tile.FlipVertical()
		tile.SetCanRotateOrFlip(false)
	}

	// now normalise our x,y so we are at 0,0 again
	normalised := make(map[string]*Tile)
	// fmt.Printf("\n\n")
	for key, value := range output {
		splits := strings.Split(key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := col + int(math.Abs(float64(min_col)))
		new_row := row + int(math.Abs(float64(min_row)))
		new_key := fmt.Sprintf("%v,%v", new_col, new_row)
		// fmt.Printf("Rotate: (%v) -> (%v)\n", key, new_key)
		normalised[new_key] = value
	}

	return normalised
}

func DebugMatrix(matrix map[string]*Tile, showPositions bool) string {
	size := int(math.Sqrt(float64(len(matrix))))
	line := ""
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			key := fmt.Sprintf("%v,%v", col, row)
			tile := matrix[key]
			if showPositions {
				line += "(" + key + ") "
			}
			line += tile.TileId
			line += " "
		}
		line += "\n\n"
	}
	return fmt.Sprintf(line)
}

func ValidateMatrix(matrix map[string]*Tile) (string, bool) {
	size := int(math.Sqrt(float64(len(matrix))))
	line := ""
	errorsExist := false
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			key := fmt.Sprintf("%v,%v", col, row)
			tile := matrix[key]
			northKey := fmt.Sprintf("%v,%v", col, row+1)
			southKey := fmt.Sprintf("%v,%v", col, row-1)
			eastKey := fmt.Sprintf("%v,%v", col+1, row)
			westKey := fmt.Sprintf("%v,%v", col-1, row)

			northRequired := row < size
			southRequired := row > 0
			eastRequired := col < size
			westRequired := col > 0

			// check all tiles around this tile to ensure their keys line up
			northTile, northExists := matrix[northKey]
			southTile, southExists := matrix[southKey]
			eastTile, eastExists := matrix[eastKey]
			westTile, westExists := matrix[westKey]

			errors := ""
			if northRequired != northExists {
				errorsExist = true
				errors += fmt.Sprintf("North required %v, north exists %v.\n", northRequired, northExists)
			} else if northExists {
				// verify alignment
				if tile.NorthEdge != northTile.SouthEdge {
					errorsExist = true
					errors += fmt.Sprintf("North tile %v (%v) edge does not match %v != %v.\n", northTile.TileId, northKey, northTile.SouthEdge, tile.NorthEdge)
					errors += fmt.Sprintf("Tile (required matching north edge)\n%v\n\n", tile.Debug(false))
					errors += fmt.Sprintf("\nNorth Tile\n%v\n\n", northTile.Debug(false))
				}
			}
			if southRequired != southExists {
				errorsExist = true
				errors += fmt.Sprintf("South required %v, south exists %v.\n", southRequired, southExists)
			} else if southExists {
				// verify alignment
				if tile.SouthEdge != southTile.NorthEdge {
					errorsExist = true
					errors += fmt.Sprintf("South tile %v (%v) edge does not match %v != %v.\n", southTile.TileId, southKey, southTile.NorthEdge, tile.SouthEdge)
					errors += fmt.Sprintf("Tile (required matching south edge)\n%v\n\n", tile.Debug(false))
					errors += fmt.Sprintf("\nSouth Tile\n%v\n\n", southTile.Debug(false))
				}
			}

			if eastRequired != eastExists {
				errorsExist = true
				errors += fmt.Sprintf("East required %v, east exists %v.\n", eastRequired, eastExists)
			} else if eastExists {
				// verify alignment
				if tile.EastEdge != eastTile.WestEdge {
					errorsExist = true
					errors += fmt.Sprintf("East tile %v (%v) edge does not match %v != %v.\n", eastTile.TileId, eastKey, eastTile.WestEdge, tile.EastEdge)
					errors += fmt.Sprintf("Tile (required matching east edge)\n%v\n\n", tile.Debug(false))
					errors += fmt.Sprintf("\nEast Tile\n%v\n\n", eastTile.Debug(false))
				}
			}

			if westRequired != westExists {
				errorsExist = true
				errors += fmt.Sprintf("West required %v, west exists %v.\n", westRequired, westExists)
			} else if westExists {
				// verify alignment
				if tile.WestEdge != westTile.EastEdge {
					errorsExist = true
					errors += fmt.Sprintf("West tile %v (%v) edge does not match %v != %v.\n", westTile.TileId, westKey, westTile.EastEdge, tile.WestEdge)
					errors += fmt.Sprintf("Tile (requires matching west edge)\n%v\n\n", tile.Debug(false))
					errors += fmt.Sprintf("\nWest Tile\n%v\n\n", westTile.Debug(false))
				}
			}

			if errorsExist {
				line += fmt.Sprintf("\n%v (%v)\n", tile.TileId, key)
				line += errors
			} else {
				line += fmt.Sprintf("\n%v (%v) OK\n", tile.TileId, key)
			}
		}
		line += "\n\n"
	}
	return fmt.Sprintf(line), errorsExist
}
