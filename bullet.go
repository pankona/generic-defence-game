package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	x, y   float64
	speed  float64
	image  *ebiten.Image
	active bool
	target targetable
}

type targetable interface {
	GetX() float64
	GetY() float64
}

func NewBullet(x, y float64, target targetable) Bullet {
	b := Bullet{
		x:      x,
		y:      y,
		speed:  8.0,
		active: true,
		target: target,
	}
	b.image = ebiten.NewImage(4, 4)
	b.image.Fill(color.White)
	return b
}

func (b *Bullet) Update() {
	// 弾の動きのロジック
	dx := b.target.GetX() - b.x
	dy := b.target.GetY() - b.y
	dist := math.Sqrt(dx*dx + dy*dy)

	if dist != 0 {
		dx /= dist
		dy /= dist
	}

	b.x += dx * b.speed
	b.y += dy * b.speed

	// 画面外に出たら弾を消す
	if b.x < 0 || b.x > 640 || b.y < 0 || b.y > 480 {
		b.active = false
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	// 弾の描画ロジック
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, op)
}
