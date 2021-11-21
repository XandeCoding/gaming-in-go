package entities

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	sprite Sprite
}

func (player *Player) initialize(renderer *sdl.Renderer, x, y float64) {
	sprite := Sprite{}
	x = float64(x) / 2.0
	y = float64(y - SpriteSize)
	
	sprite.initialize(renderer, "player.bmp", x, y)
	
	player.sprite = sprite
}

func (player Player) draw(renderer *sdl.Renderer) {
	player.sprite.x = player.sprite.x - SpriteSize/2.0
	player.sprite.y = player.sprite.y - SpriteSize/2.0

	player.sprite.draw(renderer)
}

func (player *Player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		player.sprite.x -= SpriteSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		player.sprite.x += SpriteSpeed
	}
}
