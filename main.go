package main

import (
	"sdl"
	"fmt"
	"blokus"
)

var running = true
var screen *sdl.Surface

var cellSize int
var offsetX int
var offsetY int

const margin = 4
const borderWidth = 1
const boardStart = margin + borderWidth
const gridSize = 2*margin + 2*borderWidth + blokus.BoardSize

const emptyCellColor = 0x0

func main() {
	fmt.Println("loaded", len(blokus.Tiles), "shapes")
	blokus.Board[2][4] = blokus.Red
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}
	resize(640, 480)
	sdl.WM_SetCaption("Blokus", "")

	for running {
		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
			case sdl.KEYDOWN:
				keyDown(e.Keyboard())
			case sdl.KEYUP:
				keyUp(e.Keyboard())
			case sdl.MOUSEBUTTONDOWN:
				mouseDown(e.MouseButton())
			case sdl.VIDEORESIZE:
				r := e.Resize()
				resize(int(r.W), int(r.H))
			default:
			}
			if !running {
				break
			}
		}

		draw()

		sdl.Delay(25)
	}

	sdl.Quit()
}
func resize(width int, height int) {
	screen = sdl.SetVideoMode(width, height, 32, sdl.RESIZABLE)
	if screen == nil {
		panic(sdl.GetError())
	}

	gridSizePixels := min(width, height)
	cellSize = gridSizePixels / gridSize
	offsetX = (width - gridSizePixels) / 2
	offsetY = (height - gridSizePixels) / 2
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func stop() {
	running = false
}

func draw() {
	screen.FillRect(nil, 0xa0a0a0)
	drawBorder()
	screen.Flip()
}

func drawBorder() {
	for y := 0; y < blokus.BoardSize; y++ {
		for x := 0; x < blokus.BoardSize; x++ {
			var color uint32
			switch blokus.Board[x][y] {
			case blokus.Empty:
				color = 0x0
			case blokus.Blue:
				color = 0x0000ff
			case blokus.Yellow:
				color = 0xffff00
			case blokus.Red:
				color = 0xff0000
			case blokus.Green:
				color = 0x00ff00
			}
			screen.FillRect(rect(pixelX(x)+1, pixelY(y)+1, cellSize-2, cellSize-2), color)
		}
	}
}
func pixelX(cellX int) int {
	return offsetX + cellSize*(boardStart+cellX)
}
func pixelY(cellY int) int {
	return offsetY + cellSize*(boardStart+cellY)
}

func rect(x int, y int, w int, h int) *sdl.Rect {
	return &sdl.Rect{int16(x), int16(y), uint16(w), uint16(h)}
}

func keyUp(k *sdl.KeyboardEvent) {
	//debugKeyboardEvent(k)
}
func keyDown(k *sdl.KeyboardEvent) {
	//debugKeyboardEvent(k)
	if k.Keysym.Sym == 27 {
		stop()
	}
}
func debugKeyboardEvent(k *sdl.KeyboardEvent) {
	println("")
	println(k.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(k.Keysym.Sym)))

	fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", k.Type, k.Which, k.State, k.Pad0[0])
	fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", k.Keysym.Scancode, k.Keysym.Sym, k.Keysym.Mod, k.Keysym.Unicode)
}

func mouseDown(mouseButton *sdl.MouseButtonEvent) {
	println("Click:", mouseButton.Button, mouseButton.X, mouseButton.Y)
}
