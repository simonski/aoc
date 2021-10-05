package app

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	goutils "github.com/simonski/goutils"
)

// AOC_2020_20 is the entrypoint

type Image struct {
	TileIds []string
	TileMap map[string]*Tile
	Matrix  map[string]*Tile
}

func (i *Image) Size() int {
	return len(i.TileIds)
}

func NewImageFromString(data string) *Image {
	lines := strings.Split(data, "\n")
	tileMap := make(map[string]*Tile)
	tiles := make([]string, 0)
	tileData := make([]string, 0)
	tileId := ""
	image := &Image{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// end tile
			tile := NewTile(tileId, tileData)
			tileMap[tileId] = tile
			tiles = append(tiles, tileId)
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
	// i := Image{Tiles: tiles}

	sort.Strings(tiles)
	image.TileIds = tiles
	image.TileMap = tileMap
	return image
}

// Arrange in this attempt will take the first tile and assign it 0,0 on my grid
// any connected tile will then have coordinates relative to this image
func (image *Image) Arrange(DEBUG bool) map[string]*Tile {
	// remaining := i.CopyMap()
	// if DEBUG {
	// 	fmt.Printf("Map size %v\n", len(remaining))
	// }

	matrix := make(map[string]*Tile)
	inprogress := make(map[string]*Tile)
	tile := image.GetTile(image.TileIds[0])
	matrix["0,0"] = tile
	requiredSize := len(image.TileIds)
	iteration := 0
	inprogress[tile.TileId] = tile
	// delete(remaining, tile.TileId)

	// iterate until we have all places allocated in our matrix (the NxN image)
	for len(matrix) != requiredSize {
		// for each entry in the matrix work out if there are any relations to attach
		for key, tile := range matrix {
			if tile.IsEvaluated {
				continue
			}
			matrix_x, _ := strconv.Atoi(strings.Split(key, ",")[0])
			matrix_y, _ := strconv.Atoi(strings.Split(key, ",")[1])
			north_key := fmt.Sprintf("%v,%v", matrix_x, matrix_y+1)
			south_key := fmt.Sprintf("%v,%v", matrix_x, matrix_y-1)
			east_key := fmt.Sprintf("%v,%v", matrix_x+1, matrix_y)
			west_key := fmt.Sprintf("%v,%v", matrix_x-1, matrix_y)
			_, north_exists := matrix[north_key]
			_, south_exists := matrix[south_key]
			_, east_exists := matrix[east_key]
			_, west_exists := matrix[west_key]
			if north_exists && south_exists && east_exists && west_exists {
				tile.IsEvaluated = true
				continue
			}
			for _, candidate := range image.TileMap {
				if candidate.IsEvaluated {
					continue
				}
				iteration++

				if image.AttemptAttach(tile, candidate) {
					count := 0
					if tile.NorthTile != nil {
						count++
						matrix[north_key] = tile.NorthTile
						if DEBUG {
							fmt.Printf("Iteration [%v] (%v)=%v\n", iteration, north_key, tile.NorthTile.TileId)
						}
					}
					if tile.SouthTile != nil {
						count++
						matrix[south_key] = tile.SouthTile
						if DEBUG {
							fmt.Printf("Iteration [%v] (%v)=%v\n", iteration, south_key, tile.SouthTile.TileId)
						}
					}
					if tile.EastTile != nil {
						count++
						matrix[east_key] = tile.EastTile
						if DEBUG {
							fmt.Printf("Iteration [%v] (%v)=%v\n", iteration, east_key, tile.EastTile.TileId)
						}
					}
					if tile.WestTile != nil {
						count++
						matrix[west_key] = tile.WestTile
						if DEBUG {
							fmt.Printf("Iteration [%v] (%v)=%v\n", iteration, west_key, tile.WestTile.TileId)
						}
					}

					if count == 4 {
						tile.IsEvaluated = true
					}
				}

				if DEBUG {
					fmt.Printf("Iteration [%v] found %v tiles.\n", iteration, len(matrix))
				}
				if len(matrix) == requiredSize {
					break
				}

			}
			if len(matrix) == requiredSize {
				break
			}

		}
	}

	matrix = NormaliseTileMap(matrix)
	image.Matrix = matrix
	return matrix

}

// Debug prints the arranged image out
func (image *Image) Debug(hideTileEdges bool, showBorderBetweenTiles bool, showCoordinates bool) string {
	rows := int(math.Sqrt(float64(image.Size())))
	cols := rows
	content := ""
	maxDepth := 10
	if hideTileEdges {
		maxDepth = 8
	}

	matrix := image.Matrix

	for matrixRow := rows - 1; matrixRow >= 0; matrixRow-- {
		if showCoordinates {
			tileRow := ""
			for matrixCol := 0; matrixCol < cols; matrixCol++ {
				tileCoordinates := fmt.Sprintf("%v,%v", matrixCol, matrixRow) //  0,1
				tile, tileExists := matrix[tileCoordinates]
				tileId := "No tile"
				if tileExists {
					tileId = tile.TileId
				}
				tileRow += tileCoordinates + "(" + tileId + ")  "
			}
			content += tileRow
			content += "\n"
		}
		for depth := 0; depth < maxDepth; depth++ {
			for matrixCol := 0; matrixCol < cols; matrixCol++ {
				tileCoordinates := fmt.Sprintf("%v,%v", matrixCol, matrixRow) //  0,1
				tile, tileExists := matrix[tileCoordinates]
				tileRowAsString := ""
				if tileExists {
					tileRowAsString = strings.Split(tile.Debug(hideTileEdges), "\n")[depth]
					tileRowAsString = goutils.ReverseString(tileRowAsString)
				} else {
					if hideTileEdges {
						tileRowAsString = "        "
					} else {
						tileRowAsString = "          "
					}
				}
				if showBorderBetweenTiles {
					tileRowAsString += " "
				}

				content += tileRowAsString
			}
			content = strings.TrimSpace(content)
			content += "\n"
		}
		// content += "\n"
		if showBorderBetweenTiles {
			content += "\n"
		}

	}

	// t := matrix["0,0"].Debug(hideTileEdges)
	// content += "\n\n" + t
	return strings.TrimSpace(content)
}

// GetTile returns the tile keyed by its Id
func (i *Image) GetTile(tileId string) *Tile {
	tile, _ := i.TileMap[tileId]
	return tile
}

// AttemptAttach will attempt to attach on any side to a Tile
// The tile will be rotated and flipped in all combinations to attempt to attach on any side
// then it will FlipHorizontal and attempt 4 times, then it will FlipVertical and attempt 4 times
func (i *Image) AttemptAttach(t *Tile, candidate *Tile) bool {
	if t == candidate {
		return false
	}
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	// do not rotate/flip if the tile is clamped
	if !candidate.CanRotateOrFlip() {
		return false
	}

	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}

	// back to 0 and FlipVertical
	candidate.Rotate()
	candidate.FlipVertical()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}

	// back to 0 and FlipHorizontal
	candidate.Rotate()
	candidate.FlipHorizontal()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	candidate.Rotate()
	if t.AttachAnyConnection(candidate) > 0 {
		return true
	}
	return false

}
