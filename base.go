package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Base (本拠地)を表す構造体
type Base struct {
	x, y float64
	HP   int
	img  *ebiten.Image
}

// Baseの初期化
func NewBase() *Base {
	img := ebiten.NewImage(32, 32) // 本拠地の画像サイズ
	img.Fill(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	return &Base{
		x:   600, // 位置の調整
		y:   440, // 位置の調整
		HP:  20,  // 本拠地のヒットポイント
		img: img,
	}
}

// Baseの描画
func (b *Base) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.img, op)
}

func (b *Base) IsHit(bulletX, bulletY float64) bool {
	const enemyRadius, bulletRadius = 8, 2 // 敵と弾の半径。適切なサイズに調整してください

	// 敵と弾の中心間の距離を計算
	dx := b.x + enemyRadius - bulletX
	dy := b.y + enemyRadius - bulletY
	distance := math.Sqrt(dx*dx + dy*dy)

	// 2つの円の半径の合計よりも距離が小さい場合、当たりと判定
	return distance < (enemyRadius + bulletRadius)
}

func (b *Base) GetX() float64 {
	return b.x
}

func (b *Base) GetY() float64 {
	return b.y
}

func (b *Base) GetRadius() float64 {
	return 16
}
