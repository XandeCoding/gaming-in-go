package entities

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Bullet struct {
	sprite Sprite
	angle  float64
	active bool
}

const (
	bulletSize  = 32
	bulletSpeed = 0.15
)

func (bullet *Bullet) initialize(renderer *sdl.Renderer, x, y float64) {
	sprite := Sprite{}
	sprite.initialize(renderer, "player_bullet.bmp", x, y)

	bullet.sprite = sprite
}

func (bullet Bullet) draw(renderer *sdl.Renderer) {
	if !bullet.active {
		return
	}

	x := int32(bullet.sprite.x - bulletSize/2.0)
	y := int32(bullet.sprite.y - bulletSize/2.0)

	renderer.Copy(
		bullet.sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: x, Y: y, W: bulletSize, H: bulletSize},
	)
}

func (bullet *Bullet) update() {
	bullet.sprite.x += bulletSpeed * math.Cos(bullet.angle)
	bullet.sprite.y += bulletSpeed * math.Sin(bullet.angle)

	// TODO: Move consts screenWidth and screenHeight to a public place
	if bullet.sprite.x > 600 || bullet.sprite.x < 0 || bullet.sprite.y > 800 || bullet.sprite.y < 0 {
		bullet.active = false
	}
}

// TODO: Create a BulletPool struct, and improve names if necessarily
var BulletPool []*Bullet

// TODO: return pool
func InitializePool(renderer *sdl.Renderer) {
	for index := 0; index < 30; index++ {
		bullet := Bullet{}
		bullet.initialize(renderer, 0, 0)
		BulletPool = append(BulletPool, &bullet)
	}
}

func getFromPool() (*Bullet, bool) {
	for _, bullet := range BulletPool {
		if !bullet.active {
			return bullet, true
		}
	}

	return nil, false
}
