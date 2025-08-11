package aoc2020

/*
--- Day 24: Lobby Layout ---
Your raft makes it to the tropical island; it turns out that the small crab was an excellent navigator. You make your way to the resort.

As you enter the lobby, you discover a small problem: the floor is being renovated. You can't even reach the check-in desk until they've finished installing the new tile floor.

The tiles are all hexagonal; they need to be arranged in a hex grid with a very specific color pattern. Not in the mood to wait, you offer to help figure out the pattern.

The tiles are all white on one side and black on the other. They start with the white side facing up. The lobby is large enough to fit whatever pattern might need to appear there.

A member of the renovation crew gives you a list of the tiles that need to be flipped over (your puzzle input). Each line in the list identifies a single tile that needs to be flipped by giving a series of steps starting from a reference tile in the very center of the room. (Every line starts from the same reference tile.)

Because the tiles are hexagonal, every tile has six neighbors: east, southeast, southwest, west, northwest, and northeast. These directions are given in your list, respectively, as e, se, sw, w, nw, and ne. A tile is identified by a series of these directions with no delimiters; for example, esenee identifies the tile you land on if you start at the reference tile and then move one tile east, one tile southeast, one tile northeast, and one tile east.

Each time a tile is identified, it flips from white to black or from black to white. Tiles might be flipped more than once. For example, a line like esew flips a tile immediately adjacent to the reference tile, and a line like nwwswee flips the reference tile itself.

Here is a larger example:

sesenwnenenewseeswwswswwnenewsewsw
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
wseweeenwnesenwwwswnew
In the above example, 10 tiles are flipped once (to black), and 5 more are flipped twice (to black, then back to white). After all of these instructions have been followed, a total of 10 tiles are black.

Go through the renovation crew's list and determine which tiles they need to flip. After all of the instructions have been followed, how many tiles are left with the black side up?

*/

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andybons/gogif"
	"github.com/fogleman/gg"
	"github.com/simonski/aoc/utils"
	cli "github.com/simonski/cli"
	goutils "github.com/simonski/goutils"
)

func (app *Application) Y2020D24_Summary() *utils.Summary {
	s := utils.NewSummary(2020, 24)
	s.Name = "Lobby Layout"
	s.ProgressP1 = utils.Completed
	s.ProgressP2 = utils.Completed
	return s
}

func (app *Application) Y2020D24() {
	app.Y2020D24P1()
	app.Y2020D24P2()
}

func (app *Application) Y2020D24P1() {
	AOC_2020_24_part1_attempt1(app.CLI)
}

func (app *Application) Y2020D24P2() {
	AOC_2020_24_part2_attempt1(app.CLI)
	app.Y2020D24P2Render()
}

func (app *Application) Y2020D24P2Render() {
	grid := NewHexGrid()
	grid.RenderAnimation(DAY_24_INPUT, "animation.gif", 100, 1)
}

func AOC_2020_24_part1_attempt1(cli *cli.CLI) {
	start := time.Now()
	grid := NewHexGrid()
	grid.PlayPart1(DAY_24_INPUT)
	end := time.Now()
	fmt.Printf("Day 24 Part 1 Black Count: %v, %v\n", grid.BlackCount(), end.Sub(start))
}

func AOC_2020_24_part2_attempt1(cli *cli.CLI) {
	start := time.Now()
	grid := NewHexGrid()
	grid.PlayPart1(DAY_24_INPUT)
	grid.PlayPart2(0)

	end := time.Now()
	fmt.Printf("%v\n", end.Sub(start))

}

type HexGrid struct {
	Centre *Hex
	Cache  map[string]*Hex
}

func (grid *HexGrid) BlackCount() int {
	count := 0
	for _, hex := range grid.Cache {
		if !hex.White {
			count++
		}
	}
	return count
}

func (grid *HexGrid) PlayPart1(instructions string) {
	splits := strings.Split(instructions, "\n")
	for _, address := range splits {
		// x, y := grid.Coordinates(address)
		hex := grid.FindByAddress(address)
		hex.Flip()
	}
}

func (grid *HexGrid) PlayPart2(day int) (int, int, int, int) {

	stayWhite := 0
	stayBlack := 0
	turnWhite := 0
	turnBlack := 0

	newCache := make(map[string]*Hex)
	keys := grid.Keys()
	for _, key := range keys {
		hex := grid.FindByCoordinates(key)
		for _, nkey := range hex.Neighbours() {
			grid.FindByCoordinates(nkey)
		}
	}

	// fmt.Printf("There are %v tiles.\n", len(keys))
	keys = grid.Keys()
	for _, key := range keys {
		// fmt.Printf("1: %v\n", key)
		hex := grid.FindByCoordinates(key)
		// fmt.Printf("2")
		neighbours := hex.Neighbours()
		// fmt.Printf("4")
		count := 0
		for _, key := range neighbours {
			// fmt.Printf("5.")
			hex := grid.FindByCoordinates(key)
			if !hex.White {
				count++
			}
		}
		// fmt.Printf("6")
		copy := hex.Copy()
		if !hex.White && (count == 0 || count > 2) {
			// fmt.Printf("[Play] [Day %v] (%v,%v) is Black, has %v neighbours, flipping white.\n", day, copy.x, copy.y, count)
			copy.Flip()
			turnWhite++
		} else if hex.White && count == 2 {
			// fmt.Printf("[Play] [Day %v] (%v,%v) is White, has %v neighbours, flipping black.\n", day, copy.x, copy.y, count)
			copy.Flip()
			turnBlack++
		} else if hex.White {
			// fmt.Printf("[Play] [Day %v] (%v,%v) is White, has %v neighbours, ignoring.\n", day, copy.x, copy.y, count)
			stayWhite++
		} else {
			// fmt.Printf("[Play] [Day %v] (%v,%v) is Black, has %v neighbours, ignoring.\n", day, copy.x, copy.y, count)
			stayBlack++
		}
		// fmt.Printf("7")
		newCache[key] = copy
	}
	grid.Cache = newCache
	return stayWhite, stayBlack, turnWhite, turnBlack

}

func (hex *Hex) Neighbours() []string {
	arr := make([]string, 6)
	arr[0] = hex.NorthEast()
	arr[1] = hex.NorthWest()
	arr[2] = hex.SouthEast()
	arr[3] = hex.SouthWest()
	arr[4] = hex.East()
	arr[5] = hex.West()
	return arr
}

// Coordinates returns the relative x,y of this address from the centre
func (grid *HexGrid) CoordinatesDontCreateHexesAlongTheWay(fullAddress string) (float64, int) {
	x := 0.0
	y := 0
	parsed := grid.ParseAddress(fullAddress)
	for _, address := range parsed {
		if strings.Index(address, "se") == 0 {
			x += 0.5
			y -= 1
			address = address[2:]
		} else if strings.Index(address, "sw") == 0 {
			x -= 0.5
			y -= 1
			address = address[2:]
		} else if strings.Index(address, "ne") == 0 {
			x += 0.5
			y += 1
			address = address[2:]
		} else if strings.Index(address, "nw") == 0 {
			x -= 0.5
			y += 1
			address = address[2:]
		} else if strings.Index(address, "e") == 0 {
			x += 1
			address = address[1:]
		} else if strings.Index(address, "w") == 0 {
			x -= 1
			address = address[1:]
		}
	}
	return x, y

}

func (grid *HexGrid) CoordinatesCreateHexesAlongTheWay(fullAddress string) (float64, int) {
	x := 0.0
	y := 0
	parsed := grid.ParseAddress(fullAddress)
	for _, address := range parsed {
		if strings.Index(address, "se") == 0 {
			x += 0.5
			y -= 1
			address = address[2:]
		} else if strings.Index(address, "sw") == 0 {
			x -= 0.5
			y -= 1
			address = address[2:]
		} else if strings.Index(address, "ne") == 0 {
			x += 0.5
			y += 1
			address = address[2:]
		} else if strings.Index(address, "nw") == 0 {
			x -= 0.5
			y += 1
			address = address[2:]
		} else if strings.Index(address, "e") == 0 {
			x += 1
			address = address[1:]
		} else if strings.Index(address, "w") == 0 {
			x -= 1
			address = address[1:]
		}
		// now create this one as-is
		key := fmt.Sprintf("%v,%v", x, y)
		grid.FindByCoordinates(key)
	}
	return x, y

}

func (grid *HexGrid) FindByAddress(address string) *Hex {
	x, y := grid.CoordinatesCreateHexesAlongTheWay(address)
	key := fmt.Sprintf("%v,%v", x, y)
	hex, exists := grid.Cache[key]
	if exists {
		return hex
	} else {
		// fmt.Printf("FindByAddress creating new Hex(%v,%v)\n", x, y)
		hex := &Hex{Grid: grid, x: x, y: y, White: true}
		grid.Cache[key] = hex
		return hex
	}
}

func ToXY(key string) (float64, int) {
	splits := strings.Split(key, ",")
	x, _ := strconv.ParseFloat(splits[0], 64)
	y, _ := strconv.Atoi(splits[1])
	return x, y
}
func (grid *HexGrid) FindByCoordinates(key string) *Hex {
	hex, exists := grid.Cache[key]
	if exists {
		return hex
	} else {
		x, y := ToXY(key)
		// splits := strings.Split(key, ",")
		// x, _ := strconv.ParseFloat(splits[0], 64)
		// y, _ := strconv.ParseFloat(splits[1], 64)
		// fmt.Printf("key %v, x=%v, y=%v\n", key, x, y)
		hex := &Hex{Grid: grid, x: x, y: y, White: true}
		grid.Cache[key] = hex
		return hex
	}
}

// ParseAddress takes an address e.g.
// neeswseenwwswnwswswnw to [ ne, e, sw, se, e, nw, w, sw, nw, sw, sw, nw ]
func (g *HexGrid) ParseAddress(address string) []string {
	arr := make([]string, 0)
	for {
		if strings.Index(address, "se") == 0 {
			arr = append(arr, "se")
			address = address[2:]
		} else if strings.Index(address, "sw") == 0 {
			arr = append(arr, "sw")
			address = address[2:]
		} else if strings.Index(address, "ne") == 0 {
			arr = append(arr, "ne")
			address = address[2:]
		} else if strings.Index(address, "nw") == 0 {
			arr = append(arr, "nw")
			address = address[2:]
		} else if strings.Index(address, "e") == 0 {
			arr = append(arr, "e")
			address = address[1:]
		} else if strings.Index(address, "w") == 0 {
			arr = append(arr, "w")
			address = address[1:]
		}
		if address == "" {
			break
		}
	}
	return arr

}

func NewHexGrid() *HexGrid {
	grid := &HexGrid{}
	centre := &Hex{Grid: grid, x: 0, y: 0, White: true}
	grid.Cache = make(map[string]*Hex)
	grid.Centre = centre
	grid.Cache["0,0"] = centre
	return grid
}

// Render creates a jpeg of the day and tiles.
func (grid *HexGrid) Render(day int, filename string) {
	// keys := grid.Keys()

	COLOR_PEACH := "#ffccff"
	COLOR_BLACK := "#000000"
	COLOR_WHITE := "#ffffff"

	colorBackground := COLOR_PEACH // peach
	colorLine := COLOR_BLACK       // black
	colorFill := COLOR_BLACK       // black

	drawHex := true
	drawBoxWithBorder := false
	drawBoxWithoutBorder := false
	drawPoints := false
	drawText := false
	drawInfo := true

	overallImageWidth := 1000.0
	overallImageHeight := 1000.0

	border := 25.0 // 1pct of the overall side as a horizontal margin
	viewportW := overallImageWidth * 0.75
	viewportH := overallImageHeight * 0.75
	viewportX := overallImageWidth - viewportW - border
	viewportY := overallImageHeight/2 - (viewportH / 2) // - border

	// imageOffsetX := 0.0
	// imageOffsetY := 0.0
	imageCentreX := viewportX + viewportW/2.0 // imageOffsetX + (imageWidth / 2.0)
	imageCentreY := viewportY + viewportH/2.0 // imageOffsetY + (imageHeight / 2.0)
	tileMargin := 1.0                         // amound of space 'inside' tile box before tile is rendered

	dc := gg.NewContext(int(overallImageWidth), int(overallImageHeight))
	// fill background entirely
	dc.DrawRectangle(0, 0, float64(overallImageWidth), float64(overallImageHeight))
	dc.SetHexColor(colorBackground)
	dc.Fill()

	// now draw all boxes
	numberOfBlackTiles := grid.BlackCount()
	numberOfWhiteTiles := len(grid.Cache) - grid.BlackCount()
	numberTilesHorizontal, numberTilesVertical := grid.Dimensions()

	totalTileHeight := (viewportH - (border * 2.0)) / float64((numberTilesVertical + 1))
	totalTileWidth := (viewportW - (border * 2.0)) / float64((numberTilesHorizontal + 1))

	totalTileHeight = math.Min(totalTileHeight, totalTileWidth)
	totalTileWidth = math.Min(totalTileHeight, totalTileWidth)

	tileWidth := totalTileWidth - (tileMargin * 2.0)
	tileHeight := totalTileHeight - (tileMargin * 2.0)

	// x_offset := (totalTileWidth - (tileMargin * 2)) / 2.0
	// y_offset := (totalTileHeight - (tileMargin * 2)) / 2.0

	// fmt.Printf("%v tiles in total, width %v, height %v, tile dimensions %v,%v\n", len(grid.Keys()), totalTileWidth, totalTileHeight, numberTilesHorizontal, numberTilesVertical)

	halfTileHeight := tileHeight / 2.0
	halfTileWidth := tileWidth / 2.0
	divideBy := 2.25
	p1 := goutils.Point2DF{X: 0.0, Y: halfTileHeight}                                          // north
	p2 := goutils.Point2DF{X: halfTileWidth, Y: halfTileHeight - (halfTileHeight / divideBy)}  // north east
	p3 := goutils.Point2DF{X: halfTileWidth, Y: -halfTileHeight + (halfTileHeight / divideBy)} // south east
	p4 := goutils.Point2DF{X: 0, Y: -halfTileHeight}                                           // south
	p5 := goutils.Point2DF{X: -halfTileWidth, Y: -halfTileHeight + (halfTileHeight / divideBy)}
	p6 := goutils.Point2DF{X: -halfTileWidth, Y: halfTileHeight - (halfTileHeight / divideBy)}

	// count := 0

	// bounds := grid.Bounds()
	// min_x := bounds[0]
	// min_y := bounds[1]
	// max_x := bounds[2]
	// max_y := bounds[3]

	// start_x := 0.0

	// start_x_even := float64(int(min_x))
	// start_x_odd := min_x

	// keys := make([]string, 0)
	// for tile_y := min_y; tile_y <= max_y; tile_y += 1 {
	// 	if int(tile_y)%2 == 0 {
	// 		// is even, x's are all whole numbers
	// 		start_x = start_x_even
	// 	} else {
	// 		// y is odd, x's are all out by 0.5
	// 		start_x = start_x_odd
	// 	}
	// 	for tile_x := start_x; tile_x <= max_x; tile_x += 1.0 {
	// 		key := fmt.Sprintf("%v,%v", tile_x, tile_y)
	// 		keys = append(keys, key)
	// 	}
	// }

	keys := grid.Keys()

	for _, key := range keys {
		x, y := ToXY(key)
		// x_str := fmt.Sprintf("%v", x)
		// y_str := fmt.Sprintf("%v", y)
		// ok := false
		// if strings.Index(y_str, ".5") > -1 {
		// 	// x_str must also be a .5
		// 	if strings.Index(x_str, ".5") > -1 {
		// 		ok = true
		// 	}
		// } else if strings.Index(y_str, ".5") == -1 {
		// 	// x_str must also be a .5
		// 	if strings.Index(x_str, ".5") == -1 {
		// 		ok = true
		// 	}
		// }

		// if !ok {
		// 	continue
		// }

		// if day == 1 {
		// 	if y < -.5 || y > .5 {
		// 		continue
		// 	}
		// }

		hex, exists := grid.Cache[key]
		if !exists {
			// fmt.Printf("(%v) does not exist as a tile, not rendering\n", key)
			continue
			// } else {
			// 	fmt.Printf("(%v) does exists as a tile, rendering\n", key)
		}

		// fmt.Printf("(%v,%v) %v\n", hex.x, hex.y, !hex.White)

		lineWidth := 1.0
		if exists {
			lineWidth = 0.5
		} else {
			lineWidth = 0.1
		}

		// x_pos and y_pos are the centre of this tile on the screen
		x_pos := imageCentreX + float64((x * totalTileWidth))
		floaty := float64(y)
		y_pos := imageCentreY + float64(floaty*totalTileHeight) - (floaty * (totalTileHeight / 5.0))

		tileP1 := p1.Translate(x_pos, y_pos)
		tileP2 := p2.Translate(x_pos, y_pos)
		tileP3 := p3.Translate(x_pos, y_pos)
		tileP4 := p4.Translate(x_pos, y_pos)
		tileP5 := p5.Translate(x_pos, y_pos)
		tileP6 := p6.Translate(x_pos, y_pos)

		if drawHex {
			fill := false
			if exists {
				fill = !hex.White
			}

			if fill {
				dc.ClearPath()
				dc.SetHexColor(COLOR_BLACK)
				dc.MoveTo(tileP1.X, tileP1.Y)
				dc.LineTo(tileP2.X, tileP2.Y)
				dc.LineTo(tileP3.X, tileP3.Y)
				dc.LineTo(tileP4.X, tileP4.Y)
				dc.LineTo(tileP5.X, tileP5.Y)
				dc.LineTo(tileP6.X, tileP6.Y)
				dc.LineTo(tileP1.X, tileP1.Y)
				dc.ClosePath()
				dc.SetHexColor(colorFill)
				dc.SetLineWidth(lineWidth)
				dc.Fill()
				// 				dc.SetHexColor(colorFill)
				if drawText {
					dc.SetHexColor(COLOR_WHITE)
					dc.DrawString(key, x_pos-10, y_pos)
					dc.Fill()
				}
			} else {
				dc.SetHexColor(COLOR_BLACK)
				dc.DrawLine(tileP1.X, tileP1.Y, tileP2.X, tileP2.Y)
				dc.DrawLine(tileP2.X, tileP2.Y, tileP3.X, tileP3.Y)
				dc.DrawLine(tileP3.X, tileP3.Y, tileP4.X, tileP4.Y)
				dc.DrawLine(tileP4.X, tileP4.Y, tileP5.X, tileP5.Y)
				dc.DrawLine(tileP5.X, tileP5.Y, tileP6.X, tileP6.Y)
				dc.DrawLine(tileP6.X, tileP6.Y, tileP1.X, tileP1.Y)
				dc.Stroke()
				if drawText {
					dc.SetHexColor(colorLine)
					dc.SetLineWidth(lineWidth)
					dc.DrawString(key, x_pos-10, y_pos)
					dc.Stroke()
				}
				// dc.DrawString(key, x_pos-10, y_pos)
				if drawPoints {
					dc.DrawString(fmt.Sprintf("P1 (%v,%v)", tileP1.X, tileP1.Y), tileP1.X, tileP1.Y)
					dc.DrawString("P2", tileP2.X, tileP2.Y)
					dc.DrawString(fmt.Sprintf("P3 (%v,%v)", tileP3.X, tileP3.Y), tileP3.X, tileP3.Y)
					dc.DrawString("P4", tileP4.X, tileP4.Y)
					dc.DrawString("P5", tileP5.X, tileP5.Y)
					dc.DrawString("P6", tileP6.X, tileP6.Y)
					dc.Stroke()
				}
			}
		}

		// tile_x1 := x_pos - x_offset
		// tile_y1 := y_pos - y_offset
		// tile_x2 := fullTotalTileWidth  // - (tileMargin * 2.0)
		// tile_y2 := fullTotalTileHeight // - (tileMargin * 2.0)

		// fmt.Printf("Key %v, xpos %v ypos %v, (%v,%v)->(%v,%v)\n", key, x_pos, y_pos, tile_x1, tile_y1, tile_x2, tile_y2)

		// draw each tile
		// fmt.Printf(key)
		// so the box the tile will be in is what we will draw

		// draw the bounding box
		if drawBoxWithBorder {
			box_x := x_pos - (totalTileWidth / 2.0)
			box_y := y_pos - (totalTileHeight / 2.0)
			box_width := totalTileWidth
			box_height := totalTileHeight

			dc.DrawRectangle(box_x, box_y, box_width, box_height)
			dc.SetHexColor(colorLine)
			dc.SetLineWidth(2)
			dc.DrawString(key, x_pos-10, y_pos)
			dc.Stroke()
		}

		if drawBoxWithoutBorder {
			// draw the box we should draw inside
			box_x := x_pos - (tileWidth / 2.0)
			box_y := y_pos - (tileHeight / 2.0)
			box_width := tileWidth
			box_height := tileHeight

			dc.DrawRectangle(box_x, box_y, box_width, box_height)
			dc.SetHexColor(colorLine)
			dc.SetLineWidth(1)
			dc.Stroke()
		}

	}

	if drawInfo {
		lineHeight := 15.0
		infoX := 10.0
		dc.SetHexColor("#000000")
		dc.SetLineWidth(1)
		dc.DrawString(fmt.Sprintf("Day %v", day), infoX, lineHeight)
		dc.DrawString(fmt.Sprintf("Tiles %v", len(grid.Cache)), infoX, lineHeight*2)
		dc.DrawString(fmt.Sprintf("Dimensions %v x %v", numberTilesHorizontal, numberTilesVertical), infoX, lineHeight*3)
		dc.DrawString(fmt.Sprintf("Black %v White %v", numberOfBlackTiles, numberOfWhiteTiles), infoX, lineHeight*4)
		dc.Stroke()
	}

	dc.DrawRectangle(viewportX, viewportY, viewportW, viewportH)
	dc.SetHexColor(colorLine)
	dc.SetLineWidth(1)
	dc.Stroke()

	dc.SavePNG(filename)

}

func (grid *HexGrid) RenderAnimation(data string, filename string, days int, Delay int) {

	grid.PlayPart1(data)

	overallImageWidth := 1000.0
	overallImageHeight := 1000.0

	// frames := make([]image.Image, 0)
	ctx := gg.NewContext(int(overallImageWidth), int(overallImageHeight))
	frame := grid.RenderFrame(ctx, 0)
	bounds := frame.Bounds()
	palettedImage := image.NewPaletted(bounds, nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	quantizer.Quantize(palettedImage, bounds, frame, image.Point{})

	outGif := gif.GIF{}
	outGif.Image = append(outGif.Image, palettedImage)
	outGif.Delay = append(outGif.Delay, Delay)

	// frames = append(frames, frame1)
	for day := 1; day <= days; day++ {
		grid.PlayPart2(day)
		// filename := fmt.Sprintf("real_day_%v.png", day)
		frame := grid.RenderFrame(ctx, day)
		// frames = append(frames, frame)
		// }

		// for _, simage := range frames {
		bounds := frame.Bounds()
		palettedImage := image.NewPaletted(bounds, nil)
		quantizer := gogif.MedianCutQuantizer{NumColor: 64}
		quantizer.Quantize(palettedImage, bounds, frame, image.Point{})

		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, Delay)
	}

	f, _ := os.Create("animation.gif")
	w := bufio.NewWriter(f)
	gif.EncodeAll(w, &outGif)
	w.Flush()
	f.Close()

}

func (grid *HexGrid) RenderFrame(dc *gg.Context, day int) image.Image {

	COLOR_PEACH := "#ffccff"
	COLOR_BLACK := "#000000"
	COLOR_WHITE := "#ffffff"

	colorBackground := COLOR_PEACH // peach
	colorLine := COLOR_BLACK       // black
	colorFill := COLOR_BLACK       // black

	drawHex := true
	drawBoxWithBorder := false
	drawBoxWithoutBorder := false
	drawPoints := false
	drawText := false
	drawInfo := true

	overallImageWidth := 1000.0
	overallImageHeight := 1000.0

	border := 25.0 // 1pct of the overall side as a horizontal margin
	viewportW := overallImageWidth * 0.75
	viewportH := overallImageHeight * 0.75
	viewportX := overallImageWidth - viewportW - border
	viewportY := overallImageHeight/2 - (viewportH / 2) // - border

	// imageOffsetX := 0.0
	// imageOffsetY := 0.0
	imageCentreX := viewportX + viewportW/2.0 // imageOffsetX + (imageWidth / 2.0)
	imageCentreY := viewportY + viewportH/2.0 // imageOffsetY + (imageHeight / 2.0)
	tileMargin := 1.0                         // amound of space 'inside' tile box before tile is rendered

	// fill background entirely
	dc.DrawRectangle(0, 0, float64(overallImageWidth), float64(overallImageHeight))
	dc.SetHexColor(colorBackground)
	dc.Fill()

	// now draw all boxes
	numberOfBlackTiles := grid.BlackCount()
	numberOfWhiteTiles := len(grid.Cache) - grid.BlackCount()
	numberTilesHorizontal, numberTilesVertical := grid.Dimensions()

	totalTileHeight := (viewportH - (border * 2.0)) / float64((numberTilesVertical + 1))
	totalTileWidth := (viewportW - (border * 2.0)) / float64((numberTilesHorizontal + 1))

	totalTileHeight = math.Min(totalTileHeight, totalTileWidth)
	totalTileWidth = math.Min(totalTileHeight, totalTileWidth)

	tileWidth := totalTileWidth - (tileMargin * 2.0)
	tileHeight := totalTileHeight - (tileMargin * 2.0)

	// x_offset := (totalTileWidth - (tileMargin * 2)) / 2.0
	// y_offset := (totalTileHeight - (tileMargin * 2)) / 2.0

	// fmt.Printf("%v tiles in total, width %v, height %v, tile dimensions %v,%v\n", len(grid.Keys()), totalTileWidth, totalTileHeight, numberTilesHorizontal, numberTilesVertical)

	halfTileHeight := tileHeight / 2.0
	halfTileWidth := tileWidth / 2.0
	divideBy := 2.25
	p1 := goutils.Point2DF{X: 0.0, Y: halfTileHeight}                                          // north
	p2 := goutils.Point2DF{X: halfTileWidth, Y: halfTileHeight - (halfTileHeight / divideBy)}  // north east
	p3 := goutils.Point2DF{X: halfTileWidth, Y: -halfTileHeight + (halfTileHeight / divideBy)} // south east
	p4 := goutils.Point2DF{X: 0, Y: -halfTileHeight}                                           // south
	p5 := goutils.Point2DF{X: -halfTileWidth, Y: -halfTileHeight + (halfTileHeight / divideBy)}
	p6 := goutils.Point2DF{X: -halfTileWidth, Y: halfTileHeight - (halfTileHeight / divideBy)}

	// count := 0

	// bounds := grid.Bounds()
	// min_x := bounds[0]
	// min_y := bounds[1]
	// max_x := bounds[2]
	// max_y := bounds[3]

	// start_x := 0.0

	// start_x_even := float64(int(min_x))
	// start_x_odd := min_x

	// keys := make([]string, 0)
	// for tile_y := min_y; tile_y <= max_y; tile_y += 1 {
	// 	if int(tile_y)%2 == 0 {
	// 		// is even, x's are all whole numbers
	// 		start_x = start_x_even
	// 	} else {
	// 		// y is odd, x's are all out by 0.5
	// 		start_x = start_x_odd
	// 	}
	// 	for tile_x := start_x; tile_x <= max_x; tile_x += 1.0 {
	// 		key := fmt.Sprintf("%v,%v", tile_x, tile_y)
	// 		keys = append(keys, key)
	// 	}
	// }

	keys := grid.Keys()

	for _, key := range keys {
		x, y := ToXY(key)
		// x_str := fmt.Sprintf("%v", x)
		// y_str := fmt.Sprintf("%v", y)
		// ok := false
		// if strings.Index(y_str, ".5") > -1 {
		// 	// x_str must also be a .5
		// 	if strings.Index(x_str, ".5") > -1 {
		// 		ok = true
		// 	}
		// } else if strings.Index(y_str, ".5") == -1 {
		// 	// x_str must also be a .5
		// 	if strings.Index(x_str, ".5") == -1 {
		// 		ok = true
		// 	}
		// }

		// if !ok {
		// 	continue
		// }

		// if day == 1 {
		// 	if y < -.5 || y > .5 {
		// 		continue
		// 	}
		// }

		hex, exists := grid.Cache[key]
		if !exists {
			// fmt.Printf("(%v) does not exist as a tile, not rendering\n", key)
			continue
			// } else {
			// 	fmt.Printf("(%v) does exists as a tile, rendering\n", key)
		}

		// fmt.Printf("(%v,%v) %v\n", hex.x, hex.y, !hex.White)

		lineWidth := 1.0
		if exists {
			lineWidth = 0.5
		} else {
			lineWidth = 0.1
		}

		// x_pos and y_pos are the centre of this tile on the screen
		x_pos := imageCentreX + float64((x * totalTileWidth))
		floaty := float64(y)
		y_pos := imageCentreY + float64(floaty*totalTileHeight) - (floaty * (totalTileHeight / 5.0))

		tileP1 := p1.Translate(x_pos, y_pos)
		tileP2 := p2.Translate(x_pos, y_pos)
		tileP3 := p3.Translate(x_pos, y_pos)
		tileP4 := p4.Translate(x_pos, y_pos)
		tileP5 := p5.Translate(x_pos, y_pos)
		tileP6 := p6.Translate(x_pos, y_pos)

		if drawHex {
			fill := false
			if exists {
				fill = !hex.White
			}

			if fill {
				dc.ClearPath()
				dc.SetHexColor(COLOR_BLACK)
				dc.MoveTo(tileP1.X, tileP1.Y)
				dc.LineTo(tileP2.X, tileP2.Y)
				dc.LineTo(tileP3.X, tileP3.Y)
				dc.LineTo(tileP4.X, tileP4.Y)
				dc.LineTo(tileP5.X, tileP5.Y)
				dc.LineTo(tileP6.X, tileP6.Y)
				dc.LineTo(tileP1.X, tileP1.Y)
				dc.ClosePath()
				dc.SetHexColor(colorFill)
				dc.SetLineWidth(lineWidth)
				dc.Fill()
				// 				dc.SetHexColor(colorFill)
				if drawText {
					dc.SetHexColor(COLOR_WHITE)
					dc.DrawString(key, x_pos-10, y_pos)
					dc.Fill()
				}
			} else {
				dc.SetHexColor(COLOR_BLACK)
				dc.DrawLine(tileP1.X, tileP1.Y, tileP2.X, tileP2.Y)
				dc.DrawLine(tileP2.X, tileP2.Y, tileP3.X, tileP3.Y)
				dc.DrawLine(tileP3.X, tileP3.Y, tileP4.X, tileP4.Y)
				dc.DrawLine(tileP4.X, tileP4.Y, tileP5.X, tileP5.Y)
				dc.DrawLine(tileP5.X, tileP5.Y, tileP6.X, tileP6.Y)
				dc.DrawLine(tileP6.X, tileP6.Y, tileP1.X, tileP1.Y)
				dc.Stroke()
				if drawText {
					dc.SetHexColor(colorLine)
					dc.SetLineWidth(lineWidth)
					dc.DrawString(key, x_pos-10, y_pos)
					dc.Stroke()
				}
				// dc.DrawString(key, x_pos-10, y_pos)
				if drawPoints {
					dc.DrawString(fmt.Sprintf("P1 (%v,%v)", tileP1.X, tileP1.Y), tileP1.X, tileP1.Y)
					dc.DrawString("P2", tileP2.X, tileP2.Y)
					dc.DrawString(fmt.Sprintf("P3 (%v,%v)", tileP3.X, tileP3.Y), tileP3.X, tileP3.Y)
					dc.DrawString("P4", tileP4.X, tileP4.Y)
					dc.DrawString("P5", tileP5.X, tileP5.Y)
					dc.DrawString("P6", tileP6.X, tileP6.Y)
					dc.Stroke()
				}
			}
		}

		// tile_x1 := x_pos - x_offset
		// tile_y1 := y_pos - y_offset
		// tile_x2 := fullTotalTileWidth  // - (tileMargin * 2.0)
		// tile_y2 := fullTotalTileHeight // - (tileMargin * 2.0)

		// fmt.Printf("Key %v, xpos %v ypos %v, (%v,%v)->(%v,%v)\n", key, x_pos, y_pos, tile_x1, tile_y1, tile_x2, tile_y2)

		// draw each tile
		// fmt.Printf(key)
		// so the box the tile will be in is what we will draw

		// draw the bounding box
		if drawBoxWithBorder {
			box_x := x_pos - (totalTileWidth / 2.0)
			box_y := y_pos - (totalTileHeight / 2.0)
			box_width := totalTileWidth
			box_height := totalTileHeight

			dc.DrawRectangle(box_x, box_y, box_width, box_height)
			dc.SetHexColor(colorLine)
			dc.SetLineWidth(2)
			dc.DrawString(key, x_pos-10, y_pos)
			dc.Stroke()
		}

		if drawBoxWithoutBorder {
			// draw the box we should draw inside
			box_x := x_pos - (tileWidth / 2.0)
			box_y := y_pos - (tileHeight / 2.0)
			box_width := tileWidth
			box_height := tileHeight

			dc.DrawRectangle(box_x, box_y, box_width, box_height)
			dc.SetHexColor(colorLine)
			dc.SetLineWidth(1)
			dc.Stroke()
		}

	}

	if drawInfo {
		lineHeight := 15.0
		infoX := 10.0
		dc.SetHexColor("#000000")
		dc.SetLineWidth(1)
		dc.DrawString(fmt.Sprintf("Day %v", day), infoX, lineHeight)
		dc.DrawString(fmt.Sprintf("Tiles %v", len(grid.Cache)), infoX, lineHeight*2)
		dc.DrawString(fmt.Sprintf("Dimensions %v x %v", numberTilesHorizontal, numberTilesVertical), infoX, lineHeight*3)
		dc.DrawString(fmt.Sprintf("Black %v White %v", numberOfBlackTiles, numberOfWhiteTiles), infoX, lineHeight*4)
		dc.Stroke()
	}

	dc.DrawRectangle(viewportX, viewportY, viewportW, viewportH)
	dc.SetHexColor(colorLine)
	dc.SetLineWidth(1)
	dc.Stroke()

	return dc.Image()

}

func (grid *HexGrid) Keys() []string {
	arr := make([]string, 0)
	for key := range grid.Cache {
		arr = append(arr, key)
	}
	return arr
}

// TileHeight returns the number of tiles high this grid is
func (grid *HexGrid) Dimensions() (int, int) {
	keys := grid.Keys()
	max_x := 0
	max_y := 0
	for _, key := range keys {
		x, y := ToXY(key)
		max_x = goutils.Max(max_x, int(math.Abs(x)))
		max_y = goutils.Max(max_y, int(math.Abs(float64(y))))
	}
	width := (max_x * 2) + 1  //  * 1.5
	height := (max_y * 2) + 1 //  * 1.5
	return width, height
}

type Hex struct {
	Grid  *HexGrid
	x     float64
	y     int
	White bool
}

func (hex *Hex) Copy() *Hex {
	copy := &Hex{x: hex.x, y: hex.y, White: hex.White}
	return copy
}

func (hex *Hex) NorthEast() string {
	return fmt.Sprintf("%v,%v", hex.x+0.5, hex.y+1)
}

func (hex *Hex) NorthWest() string {
	return fmt.Sprintf("%v,%v", hex.x-0.5, hex.y+1)
}

func (hex *Hex) SouthEast() string {
	return fmt.Sprintf("%v,%v", hex.x+0.5, hex.y-1)
}

func (hex *Hex) SouthWest() string {
	return fmt.Sprintf("%v,%v", hex.x-0.5, hex.y-1)
}

func (hex *Hex) West() string {
	return fmt.Sprintf("%v,%v", hex.x-1, hex.y)
}

func (hex *Hex) East() string {
	return fmt.Sprintf("%v,%v", hex.x+1, hex.y)
}

func (h *Hex) Flip() *Hex {
	// fmt.Printf("Flip(%v,%v) %v\n", h.x, h.x, h.White)
	h.White = !h.White
	return h
}
