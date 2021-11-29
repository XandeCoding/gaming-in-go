package entities

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	sprite   Sprite
	lastShot time.Time
}

const (
	playerSpeed        = 0.35
	playerSize         = 105
	playerShotCooldown = time.Millisecond * 250
)

func (player *Player) initialize(renderer *sdl.Renderer, x, y float64) {
	sprite := Sprite{}
	x = float64(x) / 2.0
	y = float64(y - playerSize)

	sprite.initialize(renderer, "player.bmp", x, y)

	player.sprite = sprite
}

func (player Player) draw(renderer *sdl.Renderer) {
	x, y := int32(player.sprite.x-playerSize/2.0), int32(player.sprite.y-playerSize/2.0)

	renderer.Copy(
		player.sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: x, Y: y, W: playerSize, H: playerSize},
	)
}

func (player *Player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && player.sprite.x-(playerSize/2.0) > 0 {
		player.sprite.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && player.sprite.x+(playerSize/2.0) < 800 {
		player.sprite.x += playerSpeed
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(player.lastShot) >= playerShotCooldown {
			player.shoot(player.sprite.x+25, player.sprite.y-20)
			player.shoot(player.sprite.x-25, player.sprite.y-20)

			player.lastShot = time.Now()
		}
	}
}

func (player *Player) shoot(x, y float64) {
	if bullet, ok := getFromPool(); ok {
		bullet.active = true
		bullet.sprite.x = x
		bullet.sprite.y = y
		bullet.angle = 270 * (math.Pi / 180)
	}
}
