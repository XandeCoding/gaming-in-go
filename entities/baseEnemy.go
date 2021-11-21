// TODO: create generic interface to all sprites that has animation
package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO: CRIAR INTERFACES GENERICAS
type BaseEnemy struct {
	texture *sdl.Texture
	x, y    float64
}

const (
	BaseEnemySpeed = 0.35
	BaseEnemySize  = 105
)

func NewBaseEnemy(renderer *sdl.Renderer, x, y float64) (baseEnemy BaseEnemy, err error) {
	// TODO: DEIXAR ISTO GENERICO
	image, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return baseEnemy, fmt.Errorf("loading player sprite: %v", err)
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	baseEnemy.texture = texture

	if err != nil {
		return baseEnemy, fmt.Errorf("creating player Texture: %v", err)
	}

	// TODO: CRIAR FUNCAO QUE ABSTRAI ESSE CALCULO 
	baseEnemy.x = x
	baseEnemy.y = y

	return baseEnemy, nil
}

func (baseEnemy *BaseEnemy) Draw(renderer *sdl.Renderer) {
	x := int32(baseEnemy.x - BaseEnemySize/2.0)
	y := int32(baseEnemy.y - BaseEnemySize/2.0)

	renderer.CopyEx(
		baseEnemy.texture,
		&sdl.Rect{X: 0, Y: 0, W: BaseEnemySize, H: BaseEnemySize},
		&sdl.Rect{X: x, Y: y, W: BaseEnemySize, H: BaseEnemySize},
		180, &sdl.Point{X: BaseEnemySize / 2, Y: BaseEnemySize / 2},
		sdl.FLIP_NONE,
	)
}
