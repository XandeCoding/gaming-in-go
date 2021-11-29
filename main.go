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

	player := entities.Create(entities.PlayerInstance, renderer, screenWidth, screenHeight)

	var enemies []entities.Renderable

	for column := 0; column < 5; column++ {
		for line := 0; line < 3; line++ {
			// TODO: Create struct to create this enemies
			x := (float64(column)/5)*screenWidth + (105 / 2.0)
			y := float64(line)*105 + (105 / 2.0)

			enemy := entities.Create(entities.BasicEnemyInstance, renderer, x, y)
			if enemy != nil {
				enemies = append(enemies, enemy)
			}
		}
	}

	entities.InitializePool(renderer)

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

		entities.Draw(player, renderer)
		entities.Update(player)

		for _, enemy := range enemies {
			entities.Draw(enemy, renderer)
		}

		for _, bullet := range entities.BulletPool {
			entities.Draw(bullet, renderer)
			entities.Update(bullet)
		}

		renderer.Present()
	}
}
