package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture *sdl.Texture
	x, y    float64
}

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
