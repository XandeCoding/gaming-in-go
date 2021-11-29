package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type BasicEnemy struct {
	sprite Sprite
}

const (
	basicEnemySpeed = 0.35
	basicEnemySize  = 105
)

func (basicEnemy *BasicEnemy) initialize(renderer *sdl.Renderer, x, y float64) {
	basicEnemy.sprite.initialize(renderer, "basic_enemy.bmp", x, y)
	basicEnemy.sprite.x = x
	basicEnemy.sprite.y = y
}

func (basicEnemy BasicEnemy) draw(renderer *sdl.Renderer) {
	x := int32(basicEnemy.sprite.x - basicEnemySize/2.0)
	y := int32(basicEnemy.sprite.y - basicEnemySize/2.0)

	renderer.CopyEx(
		basicEnemy.sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: x, Y: y, W: basicEnemySize, H: basicEnemySize},
		180, &sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE,
	)
}

func (basicEnemy *BasicEnemy) update() {
	// TODO: IMPLEMENT ANIMATION
}
