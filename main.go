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

	if err != nil {
		fmt.Println("Creating player", err)
		return
	}

	var enemies []entities.BasicEnemy = make([]entities.BasicEnemy, 20)

	// TODO: CREATE A MONSTERS FACTORY
	for column := 0; column < 5; column++ {
		for line := 0; line < 3; line++ {
			x := (float64(column)/5)*screenWidth + (entities.BasicEnemySize / 2.0)
			y := float64(line)*entities.BasicEnemySize + (entities.BasicEnemySize / 2.0)

			enemy, err := entities.NewBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("Creating enemy", err)
				return
			}

			enemies = append(enemies, enemy)
		}
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

		entities.Draw(player, renderer)
		entities.Update(player)

		for _, enemy := range enemies {
			enemy.Draw(renderer)
		}

		renderer.Present()
	}
}
