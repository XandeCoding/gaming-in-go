package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type BasicEnemy struct {
	texture *sdl.Texture
	x, y    float64
}

const (
	BasicEnemySpeed = 0.35
	BasicEnemySize  = 105
)

func NewBasicEnemy(renderer *sdl.Renderer, x, y float64) (basicEnemy BasicEnemy, err error) {
	// TODO: DEIXAR ISTO GENERICO
	image, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return basicEnemy, fmt.Errorf("loading player sprite: %v", err)
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	basicEnemy.texture = texture

	if err != nil {
		return basicEnemy, fmt.Errorf("creating player Texture: %v", err)
	}

	// TODO: CRIAR FUNCAO QUE ABSTRAI ESSE CALCULO
	basicEnemy.x = x
	basicEnemy.y = y

	return basicEnemy, nil
}

func (basicEnemy *BasicEnemy) Draw(renderer *sdl.Renderer) {
	x := int32(basicEnemy.x - BasicEnemySize/2.0)
	y := int32(basicEnemy.y - BasicEnemySize/2.0)

	renderer.CopyEx(
		basicEnemy.texture,
		&sdl.Rect{X: 0, Y: 0, W: BasicEnemySize, H: BasicEnemySize},
		&sdl.Rect{X: x, Y: y, W: BasicEnemySize, H: BasicEnemySize},
		180, &sdl.Point{X: BasicEnemySize / 2, Y: BasicEnemySize / 2},
		sdl.FLIP_NONE,
	)
}
