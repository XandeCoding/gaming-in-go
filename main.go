package main

import (
	"fmt"
	"gaming-in-go/entities"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL: ", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenHeight, screenHeight,
		sdl.WINDOW_OPENGL,
	)

	if err != nil {
		fmt.Println("Initializing Window: ", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Println("Initializing Renderer", err)
		return
	}

	defer renderer.Destroy()

	player, err := entities.NewPlayer(renderer)

	if err != nil {
		fmt.Println("Creating player", err)
		return
	}

	for {
		// TODO: simplify this
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.Draw(renderer)

		renderer.Present()
	}
}
