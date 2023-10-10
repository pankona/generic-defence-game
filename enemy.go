package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x, y    float64
	speed   float64
	HP      int
	active  bool
	reached bool

	slowDuration int     // 鈍足効果の持続時間（フレーム数）
	normalSpeed  float64 // 通常の移動速度を保存する

	collidedWalls []string

	framesSinceLastBullet int
	bulletFrameInterval   int
}

func (e *Enemy) GetX() float64 {
	return e.x
}

func (e *Enemy) GetY() float64 {
	return e.y
}

func NewEnemy(x, y float64) Enemy {
	return Enemy{
		x:                     x,
		y:                     y,
		speed:                 2,
		HP:                    3,
		active:                true,
		reached:               false,
		slowDuration:          0,
		normalSpeed:           0,
		collidedWalls:         []string{},
		framesSinceLastBullet: 0,
		bulletFrameInterval:   30,
	}
}

func (e *Enemy) Update(g *Game) {
	// 壁との当たり判定
	for _, wall := range g.walls {
		if e.isCollidingWithWall(wall) {
			if e.slowDuration == 0 {
				e.normalSpeed = e.speed // 通常のスピードを保存
				e.speed /= 2            // 鈍足効果を適用（スピードを半分に）
			}
			e.slowDuration = 60 // 60フレーム（1秒）の鈍足効果を設定
			break
		}
	}

	// 鈍足効果の持続時間を減少させ、0になったら効果を解除
	if e.slowDuration > 0 {
		e.slowDuration--
		if e.slowDuration == 0 {
			e.speed = e.normalSpeed // 通常のスピードに戻す
		}
	}

	// 右下に到達したかどうかを判定
	if e.x >= 640 && e.y >= 480 && !e.reached {
		e.reached = true // 右下に到達したことをマーク
	}

	e.framesSinceLastBullet++
}

type Point struct {
	x, y float64
}

// 敵が壁と衝突しているかどうかを判定するメソッド
func (e *Enemy) isCollidingWithWall(wall Wall) bool {
	// 壁との当たり判定のロジックを実装
	// すでに衝突している壁に再衝突しているかのチェック
	for _, id := range e.collidedWalls {
		if id == wall.id {
			return false
		}
	}

	// 線分（壁）と矩形（敵）の当たり判定
	// この例では、敵のサイズを16x16と仮定しています。
	rectTopLeft := Point{e.x, e.y}
	rectBottomRight := Point{e.x + 16, e.y + 16}
	if lineIntersectsRect(wall.x1, wall.y1, wall.x2, wall.y2, rectTopLeft, rectBottomRight) {
		// 衝突した壁のIDを保存
		e.collidedWalls = append(e.collidedWalls, wall.id)
		return true
	}
	return false
}

// 線分（x1, y1, x2, y2）と矩形（rectTopLeft, rectBottomRight）の当たり判定
func lineIntersectsRect(x1, y1, x2, y2 float64, rectTopLeft, rectBottomRight Point) bool {
	// 線分と矩形の4つの辺との当たり判定を行う
	// この実装は簡易的なもので、より高精度な判定が必要な場合は別の方法を検討してください。
	return lineIntersectsLine(x1, y1, x2, y2, rectTopLeft.x, rectTopLeft.y, rectBottomRight.x, rectTopLeft.y) || // 上辺
		lineIntersectsLine(x1, y1, x2, y2, rectTopLeft.x, rectBottomRight.y, rectBottomRight.x, rectBottomRight.y) || // 下辺
		lineIntersectsLine(x1, y1, x2, y2, rectTopLeft.x, rectTopLeft.y, rectTopLeft.x, rectBottomRight.y) || // 左辺
		lineIntersectsLine(x1, y1, x2, y2, rectBottomRight.x, rectTopLeft.y, rectBottomRight.x, rectBottomRight.y) // 右辺
}

// 線分（x1, y1, x2, y2）と線分（x3, y3, x4, y4）の当たり判定
func lineIntersectsLine(x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	// 以下の関数は、2つの数の乗算の結果が正か負かを計算します。
	crossProductSign := func(x1, y1, x2, y2, x3, y3 float64) float64 {
		return (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	}

	d1 := crossProductSign(x1, y1, x2, y2, x3, y3)
	d2 := crossProductSign(x1, y1, x2, y2, x4, y4)
	d3 := crossProductSign(x3, y3, x4, y4, x1, y1)
	d4 := crossProductSign(x3, y3, x4, y4, x2, y2)

	// 2つの線分が交差する条件は、各線分の両端が他の線分と異なる側に位置することです。
	// これは、crossProductSignの結果の符号がd1とd2、またはd3とd4で異なる場合に真となります。
	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) && ((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	// 敵の描画ロジック
	enemyColor := color.RGBA{R: 255, G: 0, B: 0, A: 255} // 赤色の敵
	img := ebiten.NewImage(16, 16)
	img.Fill(enemyColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(img, op)
}

// 弾が敵に当たったかどうかを判定するメソッド
func (e *Enemy) IsHit(bulletX, bulletY float64) bool {
	const enemyRadius, bulletRadius = 8, 2 // 敵と弾の半径。適切なサイズに調整してください

	// 敵と弾の中心間の距離を計算
	dx := e.x + enemyRadius - bulletX
	dy := e.y + enemyRadius - bulletY
	distance := math.Sqrt(dx*dx + dy*dy)

	// 2つの円の半径の合計よりも距離が小さい場合、当たりと判定
	return distance < (enemyRadius + bulletRadius)
}

func (e *Enemy) GetRadius() float64 {
	return 8
}
