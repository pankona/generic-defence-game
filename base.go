package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Base (本拠地)を表す構造体
type Base struct {
	x, y float64
	HP   int
	img  *ebiten.Image
}

// Baseの初期化
func NewBase() Base {
	img := ebiten.NewImage(32, 32) // 本拠地の画像サイズ
	img.Fill(color.White)          // 本拠地の色を白に設定（カスタマイズ可能）
	return Base{
		x:   600, // 位置の調整
		y:   440, // 位置の調整
		HP:  10,  // 本拠地のヒットポイント
		img: img,
	}
}

// Baseの描画
func (b *Base) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.img, op)
}
