package main

import (
	"sdl"
	"fmt"
	// TODO: why is the explicit name required?
	blokus "blokus"
)

var running = true

func main() {
	fmt.Println("loaded", len(blokus.Tiles), "shapes")
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}
	var screen = sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)
	if screen == nil {
		panic(sdl.GetError())
	}
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
				screen = sdl.SetVideoMode(int(e.Resize().W), int(e.Resize().H), 32, sdl.RESIZABLE)
				if screen == nil {
					panic(sdl.GetError())
				}
			default:
			}
			if !running {
				break
			}
		}

		draw(screen)

		sdl.Delay(25)
	}

	sdl.Quit()
}

func stop() {
	running = false
}

func draw(screen *sdl.Surface) {
	screen.FillRect(nil, 0xa0a0a0)
	screen.Flip()
}

func keyUp(k *sdl.KeyboardEvent) {
	debugKeyboardEvent(k)
}
func keyDown(k *sdl.KeyboardEvent) {
	debugKeyboardEvent(k)
}
func debugKeyboardEvent(k *sdl.KeyboardEvent) {
	println("")
	println(k.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(k.Keysym.Sym)))

	fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", k.Type, k.Which, k.State, k.Pad0[0])
	fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", k.Keysym.Scancode, k.Keysym.Sym, k.Keysym.Mod, k.Keysym.Unicode)

	if k.Keysym.Sym == 27 {
		stop()
	}
}

func mouseDown(mouseButton *sdl.MouseButtonEvent) {
	println("Click:", mouseButton.Button, mouseButton.X, mouseButton.Y)
}
