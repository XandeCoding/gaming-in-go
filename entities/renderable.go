package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Renderable interface {
	initialize(renderer *sdl.Renderer, x, y float64)
	draw(renderer *sdl.Renderer)
	update()
}

const (
	PlayerInstance     = "PLAYER"
	BasicEnemyInstance = "BASIC_ENEMY"
)

func Create(renderableName string, renderer *sdl.Renderer, x, y float64) Renderable {
	if renderableName == PlayerInstance {
		player := Player{}
		player.initialize(renderer, x, y)
		return &player
	}

	if renderableName == BasicEnemyInstance {
		basicEnemy := BasicEnemy{}
		basicEnemy.initialize(renderer, x, y)
		// fmt.Println("basicEnemy", basicEnemy)
		return &basicEnemy
	}
	return nil
}

func Draw(renderable Renderable, renderer *sdl.Renderer) {
	renderable.draw(renderer)
}

func Update(renderable Renderable) {
	renderable.update()
}
