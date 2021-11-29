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
	BulletInstance     = "BULLET"
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
		return &basicEnemy
	}

	if renderableName == BulletInstance {
		bullet := Bullet{}
		bullet.initialize(renderer, 0, 0)
		return &bullet
	}

	return nil
}

func Draw(renderable Renderable, renderer *sdl.Renderer) {
	renderable.draw(renderer)
}

func Update(renderable Renderable) {
	renderable.update()
}
