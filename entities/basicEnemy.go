package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type BasicEnemy struct {
	sprite Sprite
}

const (
	BasicEnemySpeed = 0.35
	BasicEnemySize  = 105
)

func (basicEnemy *BasicEnemy) initialize(renderer *sdl.Renderer, x, y float64) {
	basicEnemy.sprite.initialize(renderer, "basic_enemy.bmp", x, y)
	basicEnemy.sprite.x = x
	basicEnemy.sprite.y = y
}

func (basicEnemy BasicEnemy) draw(renderer *sdl.Renderer) {
	x := int32(basicEnemy.sprite.x - BasicEnemySize/2.0)
	y := int32(basicEnemy.sprite.y - BasicEnemySize/2.0)

	renderer.CopyEx(
		basicEnemy.sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: BasicEnemySize, H: BasicEnemySize},
		&sdl.Rect{X: x, Y: y, W: BasicEnemySize, H: BasicEnemySize},
		180, &sdl.Point{X: BasicEnemySize / 2, Y: BasicEnemySize / 2},
		sdl.FLIP_NONE,
	)
}

func (basicEnemy *BasicEnemy) update() {
	// TODO: IMPLEMENT ANIMATION
}
