package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Wall struct {
	id             string
	x1, y1, x2, y2 float64
}

func (w *Wall) Draw(screen *ebiten.Image) {
	wallColor := color.RGBA{R: 150, G: 150, B: 150, A: 255} // 灰色の壁
	vector.StrokeLine(screen, float32(w.x1), float32(w.y1), float32(w.x2), float32(w.y2), 1, wallColor, false)
}
