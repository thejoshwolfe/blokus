package main

import (
	"sdl"
	"fmt"
)

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}
	var screen = sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)
	if screen == nil {
		panic(sdl.GetError())
	}
	sdl.WM_SetCaption("Blokus", "")

	running := true
	for running {
		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
				break
			case sdl.KEYDOWN, sdl.KEYUP:
				println("")
				println(e.Keyboard().Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keyboard().Keysym.Sym)))

				if e.Keyboard().Keysym.Sym == 27 {
					running = false
				}

				fmt.Printf("%04x ", e.Type)

				for i := 0; i < len(e.Pad0); i++ {
					fmt.Printf("%02x ", e.Pad0[i])
				}
				println()

				k := e.Keyboard()

				fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", k.Type, k.Which, k.State, k.Pad0[0])
				fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", k.Keysym.Scancode, k.Keysym.Sym, k.Keysym.Mod, k.Keysym.Unicode)
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().Button, e.MouseButton().X, e.MouseButton().Y)
			case sdl.VIDEORESIZE:
				println("resize screen ", e.Resize().W, e.Resize().H)

				screen = sdl.SetVideoMode(int(e.Resize().W), int(e.Resize().H), 32, sdl.RESIZABLE)

				if screen == nil {
					panic(sdl.GetError())
				}
			default:
			}
		}

		screen.FillRect(nil, 0xa0a0a0)

		screen.Flip()
		sdl.Delay(25)
	}

	sdl.Quit()
}
