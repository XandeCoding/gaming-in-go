package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	sprite Sprite
}

func (player *Player) initialize(renderer *sdl.Renderer, x, y float64) {
	sprite := Sprite{}

	sprite.texture = textureFromBMP(renderer, "player.bmp")
	sprite.x = float64(x) / 2.0
	sprite.y = float64(y - SpriteSize)
	player.sprite = sprite
}

func (player Player) draw(renderer *sdl.Renderer) {
	x := int32(player.sprite.x - SpriteSize/2.0)
	y := int32(player.sprite.y - SpriteSize/2.0)

	renderer.Copy(
		player.sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: SpriteSize, H: SpriteSize},
		&sdl.Rect{X: x, Y: y, W: SpriteSize, H: SpriteSize},
	)
}

func (player *Player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		player.sprite.x -= SpriteSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		player.sprite.x += SpriteSpeed
	}
}
