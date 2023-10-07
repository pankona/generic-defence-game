package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x, y             float64
	targetX, targetY float64
	speed            float64
	attack           int
	arrow            *ebiten.Image

	// TODO: 武器種ごとに設定できるようにする
	framesSinceLastBullet int
	bulletFrameInterval   int
}

func (p *Player) Update() {
	// プレイヤーの動きや攻撃などのロジックをここに書く

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		p.targetX, p.targetY = float64(x), float64(y)
	}

	// Move towards the target position
	dx := p.targetX - p.x
	dy := p.targetY - p.y
	distance := math.Sqrt(dx*dx + dy*dy)

	if distance > p.speed {
		ratio := p.speed / distance
		dx *= ratio
		dy *= ratio
	}

	p.x += dx
	p.y += dy

	p.framesSinceLastBullet++
}

func (p *Player) Draw(screen *ebiten.Image) {
	// プレイヤーの描画ロジックをここに書く

	// 矢印の回転角度を計算
	angle := math.Atan2(p.targetY-p.y, p.targetX-p.x)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-8, -8) // 矢印の中心を原点に移動
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.arrow, op)
}

func (p *Player) RotateTowards(targetX, targetY float64) {
	angle := math.Atan2(targetY-p.y, targetX-p.x)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(angle)
	p.arrow = ebiten.NewImageFromImage(p.arrow)
	p.arrow.DrawImage(p.arrow, op)
}
