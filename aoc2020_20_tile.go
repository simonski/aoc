package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/simonski/aoc/utils"
)

type Tile struct {
	TileId string
	Matrix map[string]string

	Size int

	NorthEdge string // the string "edge" of teh tile e.g. #....##.##
	SouthEdge string
	EastEdge  string
	WestEdge  string

	NorthTile *Tile // when joined to another tile, this is the relationship
	SouthTile *Tile
	EastTile  *Tile
	WestTile  *Tile

	Rotation int

	IsEvaluated bool

	canRotateOrFlip bool
}

// Detach removes all connections between the passed Tile and the NSEW tiles it has
func (tile *Tile) Detach() {
	if tile.NorthTile != nil {
		tile.NorthTile.SouthTile = nil
		tile.NorthTile = nil
	}
	if tile.SouthTile != nil {
		tile.SouthTile.NorthTile = nil
		tile.SouthTile = nil
	}
	if tile.EastTile != nil {
		tile.EastTile.WestTile = nil
		tile.EastTile = nil
	}
	if tile.WestTile != nil {
		tile.WestTile.EastTile = nil
		tile.WestTile = nil
	}
}

// Keys builds every key combination
func (tile *Tile) Keys() []string {
	r1 := tile.Copy().Rotate()
	r2 := r1.Copy().Rotate()
	r3 := r2.Copy().Rotate()
	r3.Rotate() // back to original
	r4 := r3.Copy().FlipVertical()
	r5 := r4.Copy().Rotate()
	r6 := r5.Copy().Rotate()
	r7 := r6.Copy().Rotate()

	keys := utils.NewCounter()
	keys.Increment(r1.NorthEdge)
	keys.Increment(r2.NorthEdge)
	keys.Increment(r3.NorthEdge)
	keys.Increment(r4.NorthEdge)
	keys.Increment(r5.NorthEdge)
	keys.Increment(r6.NorthEdge)
	keys.Increment(r7.NorthEdge)

	keys.Increment(r1.SouthEdge)
	keys.Increment(r2.SouthEdge)
	keys.Increment(r3.SouthEdge)
	keys.Increment(r4.SouthEdge)
	keys.Increment(r5.SouthEdge)
	keys.Increment(r6.SouthEdge)
	keys.Increment(r7.SouthEdge)

	keys.Increment(r1.EastEdge)
	keys.Increment(r2.EastEdge)
	keys.Increment(r3.EastEdge)
	keys.Increment(r4.EastEdge)
	keys.Increment(r5.EastEdge)
	keys.Increment(r6.EastEdge)
	keys.Increment(r7.EastEdge)

	keys.Increment(r1.WestEdge)
	keys.Increment(r2.WestEdge)
	keys.Increment(r3.WestEdge)
	keys.Increment(r4.WestEdge)
	keys.Increment(r5.WestEdge)
	keys.Increment(r6.WestEdge)
	keys.Increment(r7.WestEdge)

	klist := keys.Keys()

	return klist

}

func (tile *Tile) KeyValue(key string) int {
	value := 0
	for index := len(key) - 1; index >= 0; index-- {
		if key[index:index+1] == "#" {
			value += (index * 2)
		}
	}
	return value
}

func (toCopy *Tile) Copy() *Tile {
	t := Tile{}
	// t.Image = toCopy.Image
	t.TileId = toCopy.TileId
	t.NorthEdge = toCopy.NorthEdge
	t.SouthEdge = toCopy.SouthEdge
	t.EastEdge = toCopy.EastEdge
	t.WestEdge = toCopy.WestEdge

	t.NorthTile = toCopy.NorthTile
	t.SouthTile = toCopy.SouthTile
	t.EastTile = toCopy.EastTile
	t.WestTile = toCopy.WestTile
	t.Rotation = toCopy.Rotation

	t.Size = toCopy.Size
	return &t
}

func NewTileFromString(data string) *Tile {
	index := 0
	matrix := make(map[string]string)
	splits := strings.Split(data, "\n")
	size := len(splits)
	for row := 0; row < size; row++ {
		rowData := splits[row]
		for col := 0; col < size; col++ {
			value := rowData[col : col+1]
			key := fmt.Sprintf("%v,%v", col, row)
			matrix[key] = value
			// fmt.Printf("%v = %v\n", key, value)
			index++
		}
	}

	t := Tile{Matrix: matrix, canRotateOrFlip: true}
	t.Size = size
	return &t

}

func NewTile(id string, data []string) *Tile {
	index := 0
	matrix := make(map[string]string)

	WIDTH := len(data)

	for row := 0; row < WIDTH; row++ {
		rowData := data[row]
		for col := 0; col < WIDTH; col++ {
			value := rowData[col : col+1]
			key := fmt.Sprintf("%v,%v", col, row)
			matrix[key] = value
			index++
		}
	}

	// build north key
	t := Tile{TileId: id, Matrix: matrix, canRotateOrFlip: true}
	t.Size = WIDTH
	t.Rekey()
	return &t
}

/// Rekey rebuilds the keys for each direction
func (t *Tile) Rekey() *Tile {
	north := ""
	south := ""
	east := ""
	west := ""
	for index := 0; index < t.Size; index++ {
		north += t.Get(index, 0)
		south += t.Get(index, t.Size-1)
		east += t.Get(t.Size-1, index)
		west += t.Get(0, index)
	}
	t.NorthEdge = north
	t.SouthEdge = south
	t.WestEdge = west
	t.EastEdge = east
	return t
}

// Get returns the char value at col,row   e.g. # or .
func (t *Tile) Get(col int, row int) string {
	key := fmt.Sprintf("%v,%v", col, row)
	value, _ := t.Matrix[key]
	return value
}

// Debug returns (for STDOUT normally) the tile in character form with newlines
// func (t *Tile) Debug() string {
// 	data := ""
// 	for row := 0; row < t.Size; row++ {
// 		for col := 0; col < t.Size; col++ {
// 			value := t.Get(col, row)
// 			data += value
// 		}
// 		data += "\n"
// 	}
// 	return strings.TrimSpace(data)
// }

func (t *Tile) Debug(hideEdges bool) string {
	data := ""
	offset := 0
	maxSize := t.Size
	if hideEdges {
		offset = 1
		maxSize = t.Size - 1
	}
	for row := offset; row < maxSize; row++ {
		for col := offset; col < maxSize; col++ {
			value := t.Get(col, row)
			data += value
		}
		data += "\n"
	}
	return strings.TrimSpace(data)
}

// Rotates the tile 90 degrees clockwise, removing any relationships this
// tile has
func (t *Tile) Rotate() *Tile {
	if !t.CanRotateOrFlip() {
		return t
	}
	matrix := make(map[string]string)
	min_col := 1000
	min_row := 1000
	max_col := -1000
	max_row := -1000
	// 90 cw rotation = (x,y) -> (-y, x)
	for col := 0; col < t.Size; col++ {
		for row := 0; row < t.Size; row++ {
			value := t.Get(col, row)
			new_col := -row
			new_row := col
			min_col = utils.Min(min_col, new_col)
			min_row = utils.Min(min_row, new_row)
			max_col = utils.Max(max_col, new_col)
			max_row = utils.Max(max_row, new_row)
			// original_key := fmt.Sprintf("%v,%v", col, row)
			new_key := fmt.Sprintf("%v,%v", new_col, new_row)
			// fmt.Printf("(%v) cw90 -> (%v)\n", original_key, new_key)
			matrix[new_key] = value
		}
	}

	// now normalise our x,y so we are at 0,0 again
	matrix_normalised := make(map[string]string)
	// fmt.Printf("\n\n")
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
	// fmt.Printf("\n\n")

	t.Matrix = matrix_normalised
	t.Rekey()
	t.Detach()

	t.Rotation += 90
	if t.Rotation == 360 {
		t.Rotation = 0
	}
	return t
}

// FlipVertical Performs a Vertical 'flip' of the Tile
func (t *Tile) FlipVertical() *Tile {
	if !t.CanRotateOrFlip() {
		return t
	}
	matrix_flipped := make(map[string]string)
	for old_key, value := range t.Matrix {
		splits := strings.Split(old_key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_row := t.Size - row - 1
		new_key := fmt.Sprintf("%v,%v", col, new_row)
		// fmt.Printf("Tile.FlipVertical() %v -> %v\n", old_key, new_key)
		matrix_flipped[new_key] = value
	}
	t.Matrix = matrix_flipped
	t.Rekey()
	t.Detach()
	return t
}

// FlipHorizontal Performs a Horizontal 'flip' of the Tile
func (t *Tile) FlipHorizontal() *Tile {
	if !t.CanRotateOrFlip() {
		return t
	}
	matrix_flipped := make(map[string]string)
	for old_key, value := range t.Matrix {
		splits := strings.Split(old_key, ",")
		col, _ := strconv.Atoi(splits[0])
		row, _ := strconv.Atoi(splits[1])
		new_col := t.Size - col - 1
		new_key := fmt.Sprintf("%v,%v", new_col, row)
		matrix_flipped[new_key] = value
	}
	t.Matrix = matrix_flipped
	t.Rekey()
	return t
}

// IsNorthOf indicates if this can attach to the north of the candidate tile
func (t *Tile) IsNorthOf(candidate *Tile) bool {
	return t.SouthEdge == candidate.NorthEdge
}

// IsSouthOf indicates if this can attach to the south of the candidate tile
func (t *Tile) IsSouthOf(candidate *Tile) bool {
	return t.NorthEdge == candidate.SouthEdge
}

// IsWestOf indicates if this can attach to the west of the candidate tile
func (t *Tile) IsWestOf(candidate *Tile) bool {
	return t.EastEdge == candidate.WestEdge
}

// IsWestOf indicates if this can attach to the west of the candidate tile
func (t *Tile) IsEastOf(candidate *Tile) bool {
	return t.WestEdge == candidate.EastEdge
}

// TotalConnections - returns the nnuber of connections this tile has
func (t *Tile) TotalConnections() int {
	count := 0
	if t.NorthTile != nil {
		count++
	}
	if t.SouthTile != nil {
		count++
	}
	if t.EastTile != nil {
		count++
	}
	if t.WestTile != nil {
		count++
	}
	return count
}

func (t *Tile) IsTopLeftCorner() bool {
	return t.SouthTile != nil && t.EastTile != nil && t.NorthTile == nil && t.WestTile == nil
}

func (t *Tile) IsTopRightCorner() bool {
	return t.SouthTile != nil && t.EastTile == nil && t.NorthTile == nil && t.WestTile != nil
}

func (t *Tile) IsBottomLeftCorner() bool {
	return t.SouthTile == nil && t.EastTile != nil && t.NorthTile != nil && t.WestTile == nil
}

func (t *Tile) IsBottomRightCorner() bool {
	return t.SouthTile == nil && t.EastTile == nil && t.NorthTile != nil && t.WestTile != nil
}

// AttachAnyConnection attaches to the candidate either NSEW, bool indicates if it attached
// does NOT rotate any tile, will just attach if it can, will not attach if already attached
func (t *Tile) AttachAnyConnection(candidate *Tile) int {
	if t.IsNorthOf(candidate) {
		candidate.AttachNorth(t)
		return 1
	}
	if t.IsSouthOf(candidate) {
		candidate.AttachSouth(t)
		return 1
	}
	if t.IsEastOf(candidate) {
		candidate.AttachEast(t)
		return 1
	}
	if t.IsWestOf(candidate) {
		candidate.AttachWest(t)
		return 1
	}
	return 0
}

func (t *Tile) SetCanRotateOrFlip(canRotateOrFlip bool) {
	t.canRotateOrFlip = canRotateOrFlip
}

func (t *Tile) CanRotateOrFlip() bool {
	return t.canRotateOrFlip
}

func (t *Tile) IsComplete() bool {
	return false
	if t.NorthTile != nil && t.SouthTile != nil && t.EastTile != nil && t.WestTile != nil {
		return true
	}
	return false
}

func (t *Tile) GetConnectionName(other *Tile) string {
	if t.NorthTile == other && other.SouthTile == t {
		return "South of"
	} else if t.SouthTile == other && other.NorthTile == t {
		return "North of"
	} else if t.WestTile == other && other.EastTile == t {
		return "East of"
	} else if t.EastTile == other && other.WestTile == t {
		return "West of"
	} else {
		return "No connection"
	}
}

// Attaches to the South of the passed tile
func (t *Tile) AttachSouth(candidate *Tile) {
	// fmt.Printf("AttachSouth(%v, %v)\n", t.TileId, candidate.TileId)
	t.SouthTile = candidate
	candidate.NorthTile = t
	t.SetCanRotateOrFlip(false)
	candidate.SetCanRotateOrFlip(false)
}

func (t *Tile) AttachNorth(candidate *Tile) {
	// fmt.Printf("AttachNorth(%v, %v)\n", t.TileId, candidate.TileId)
	t.NorthTile = candidate
	candidate.SouthTile = t
	t.SetCanRotateOrFlip(false)
	candidate.SetCanRotateOrFlip(false)
}

func (t *Tile) AttachEast(candidate *Tile) {
	// fmt.Printf("AttachEast(%v, %v)\n", t.TileId, candidate.TileId)
	t.WestTile = candidate
	candidate.EastTile = t
	t.SetCanRotateOrFlip(false)
	candidate.SetCanRotateOrFlip(false)
}

func (t *Tile) AttachWest(candidate *Tile) {
	// fmt.Printf("AttachWest(%v, %v)\n", t.TileId, candidate.TileId)
	t.EastTile = candidate
	candidate.WestTile = t
	t.SetCanRotateOrFlip(false)
	candidate.SetCanRotateOrFlip(false)
}

func (tile *Tile) RotateUntilAllRegexMatch() bool {
	regexLine1, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	regexLine2, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE2)
	regexLine3, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE3)
	tileOutputR0 := strings.TrimSpace(tile.Debug(false))
	matchL1_R0 := regexLine1.FindStringSubmatch(tileOutputR0)
	matchL2_R0 := regexLine2.FindStringSubmatch(tileOutputR0)
	matchL3_R0 := regexLine3.FindStringSubmatch(tileOutputR0)
	if len(matchL1_R0) > 0 && len(matchL2_R0) > 0 && len(matchL3_R0) > 0 {
		return true
	}

	tileOutputR90 := strings.TrimSpace(tile.Rotate().Debug(false))
	matchL1_R90 := regexLine1.FindStringSubmatch(tileOutputR90)
	matchL2_R90 := regexLine2.FindStringSubmatch(tileOutputR90)
	matchL3_R90 := regexLine3.FindStringSubmatch(tileOutputR90)
	if len(matchL1_R90) > 0 && len(matchL2_R90) > 0 && len(matchL3_R90) > 0 {
		return true
	}

	tileOutputR180 := strings.TrimSpace(tile.Rotate().Debug(false))
	matchL1_R180 := regexLine1.FindStringSubmatch(tileOutputR180)
	matchL2_R180 := regexLine2.FindStringSubmatch(tileOutputR180)
	matchL3_R180 := regexLine3.FindStringSubmatch(tileOutputR180)
	if len(matchL1_R180) > 0 && len(matchL2_R180) > 0 && len(matchL3_R180) > 0 {
		return true
	}

	tileOutputR270 := strings.TrimSpace(tile.Rotate().Debug(false))
	matchL1_R270 := regexLine1.FindStringSubmatch(tileOutputR270)
	matchL2_R270 := regexLine2.FindStringSubmatch(tileOutputR270)
	matchL3_R270 := regexLine3.FindStringSubmatch(tileOutputR270)
	if len(matchL1_R270) > 0 && len(matchL2_R270) > 0 && len(matchL3_R270) > 0 {
		return true
	}

	return false

}

// FindSeaMonsters Once the data is rotated, window over the whole image to find the monsters
// this returns a [] of []string

func (tile *Tile) FindSeaMonsters() map[string][][]string {
	monsters := make(map[string][][]string)
	content := strings.Split(tile.Debug(false), "\n")
	snapshot_width := 20
	snapshot_height := 3
	max_rows := len(content) - (snapshot_height - 1)
	max_cols := len(content[0]) - (snapshot_width - 1)
	for row := 0; row < max_rows; row++ {
		for col := 0; col < max_cols; col++ {
			snap := tile.Snapshot(col, row, snapshot_width, snapshot_height, content)
			isMonster, line1Match, line2Match, line3Match := tile.IsSeaMonster(snap)

			arr := make([][]string, 0)
			arr = append(arr, line1Match)
			arr = append(arr, line2Match)
			arr = append(arr, line3Match)

			if isMonster {
				key := fmt.Sprintf("%v,%v", col, row)
				monsters[key] = arr
			}
		}
	}
	return monsters
}

func (tile *Tile) Snapshot(x int, y int, width int, height int, content []string) []string {
	snapshot := make([]string, 0)
	for row := y; row < y+height; row++ {
		line := content[row]
		sequence := line[x : x+width]
		snapshot = append(snapshot, sequence)
	}
	return snapshot
}

func (tile *Tile) IsSeaMonster(candidate []string) (bool, []string, []string, []string) {
	regexLine1, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE1)
	regexLine2, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE2)
	regexLine3, _ := regexp.Compile(DAY_20_NESSIE_REGEX_LINE3)
	matchL1 := regexLine1.FindStringSubmatch(candidate[0])
	matchL2 := regexLine2.FindStringSubmatch(candidate[1])
	matchL3 := regexLine3.FindStringSubmatch(candidate[2])
	if len(matchL1) > 0 && len(matchL2) > 0 && len(matchL3) > 0 {
		return true, matchL1, matchL2, matchL3
	}

	// matchL1 = regexLine1.FindStringSubmatch(candidate[2])
	// matchL2 = regexLine2.FindStringSubmatch(candidate[1])
	// matchL3 = regexLine3.FindStringSubmatch(candidate[0])
	// if len(matchL1) > 0 && len(matchL2) > 0 && len(matchL3) > 0 {
	// 	return true, matchL1, matchL2, matchL3
	// }

	return false, nil, nil, nil
}
