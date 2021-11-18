package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	texture *sdl.Texture
}

func NewPlayer(renderer *sdl.Renderer) (player Player, err error) {
	image, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player, fmt.Errorf("loading player sprite: %v", err)
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	player.texture = texture

	if err != nil {
		return player, fmt.Errorf("creating player Texture: %v", err)
	}

	return player, nil
}

func (player Player) Draw(renderer *sdl.Renderer) {
	renderer.Copy(
		player.texture,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
	)
}
