package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture *sdl.Texture
	x, y    float64
}

const (
	SpriteSpeed        = 0.35
	SpriteSize         = 105
)

// TODO: MAYBE PUT THIS IN A UTILS FILE
func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	image, err := sdl.LoadBMP("sprites/" + filename)
	if err != nil {
		panic(fmt.Errorf("loading sprite sprite: %v", err))
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(fmt.Errorf("loading sprite sprite: %v", err))
	}

	return texture
}

func (sprite *Sprite) initialize(renderer *sdl.Renderer, filename string, x, y float64) {
	sprite.texture = textureFromBMP(renderer, filename)
	sprite.x = x
	sprite.y = y
}

// TODO: MAYBE CREATE ANOTHER INTERFACE TO ACCEPT X AND Y ARGUMENTS AND REMOVE CAST
func (sprite Sprite) draw(renderer *sdl.Renderer) {
	x := int32(sprite.x)
	y := int32(sprite.y)

	renderer.Copy(
		sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: SpriteSize, H: SpriteSize},
		&sdl.Rect{X: x, Y: y, W: SpriteSize, H: SpriteSize},
	)
}
