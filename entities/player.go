package entities

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	texture *sdl.Texture
	x, y    float64
}

const (
	playerSpeed = 0.4
	playerSize  = 105
)

func NewPlayer(renderer *sdl.Renderer, screenWidth, screenHeight int32) (player Player, err error) {
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

	player.x = float64(screenWidth) / 2.0
	player.y = float64(screenHeight - playerSize)

	return player, nil
}

func (player *Player) Draw(renderer *sdl.Renderer) {
	x := int32(player.x - playerSize/2.0)
	y := int32(player.y - playerSize/2.0)

	renderer.Copy(
		player.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: x, Y: y, W: playerSize, H: playerSize},
	)
}

func (player *Player) Update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		player.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		player.x += playerSpeed
	}
}
