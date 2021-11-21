package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture *sdl.Texture
	x, y    float64
}

type Renderable interface {
	initialize(renderer *sdl.Renderer, x, y float64)
	draw(renderer *sdl.Renderer)
	update()
}

const (
	SpriteSpeed        = 0.35
	SpriteSize         = 105
	PlayerInstance     = "PLAYER"
	BasicEnemyInstance = "BASIC_ENEMY"
)

// TODO: MAYBE PUT THIS IN A UTILS FILE
func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	image, err := sdl.LoadBMP("sprites/" + filename)
	if err != nil {
		panic(fmt.Errorf("loading player sprite: %v", err))
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(fmt.Errorf("loading player sprite: %v", err))
	}

	return texture
}

func Create(renderableName string, renderer *sdl.Renderer, x, y float64) Renderable {
	if renderableName == PlayerInstance {
		player := Player{}
		player.initialize(renderer, x, y)
		return &player
	}

	return nil
}

func Draw(renderable Renderable, renderer *sdl.Renderer) {
	renderable.draw(renderer)
}

func Update(renderable Renderable) {
	renderable.update()
}
