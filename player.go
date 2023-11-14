package main

import (
	"image/color"
	"math"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	id               string
	x, y             float64
	targetX, targetY float64
	speed            float64
	attack           int
	arrow            *ebiten.Image

	// TODO: 武器種ごとに設定できるようにする
	framesSinceLastBullet int
	bulletFrameInterval   int
}

func NewPlayer() Player {
	arrow := ebiten.NewImage(16, 16)
	arrow.Fill(color.White)
	return Player{
		id:                  uuid.New().String(),
		x:                   screenWidth / 2,
		y:                   (screenHeight - infoAreaHeight) / 2, // 情報表示領域を除いた領域の中央に配置
		targetX:             screenWidth / 2,
		targetY:             (screenHeight - infoAreaHeight) / 2, // 情報表示領域を除いた領域の中央に配置
		speed:               4,
		attack:              1,
		arrow:               arrow,
		bulletFrameInterval: 30,
	}
}

func (p *Player) Update(g *Game) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if selectedPlayer, ok := g.unitInfo.(*Player); ok {
			if selectedPlayer.id == p.id {
				x, y := ebiten.CursorPosition()
				// マウスの位置をターゲット位置に設定
				// ターゲット位置がプレイヤーの中央と重なるように移動するために、ターゲット位置をプレイヤーの半径分ずらす
				p.targetX, p.targetY = float64(x)-p.GetRadius(), float64(y)-p.GetRadius()
			}
		}

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
	op := &ebiten.DrawImageOptions{}
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

func (p *Player) GetPosition() (x, y int) {
	return int(p.x), int(p.y)
}

func (p *Player) GetRadius() float64 {
	return 8
}

func (p *Player) GetSize() (width, height int) {
	radius := p.GetRadius()
	return int(radius * 2), int(radius * 2)
}
