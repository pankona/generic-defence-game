package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const infoAreaHeight = 120

func drawMoney(screen *ebiten.Image, money int) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Money: %d", money), screenWidth-100, 10)
}

func drawGameOver(screen *ebiten.Image) {

	const message = "Game Over"
	messageWidth := len(message) * 6 // 6 is the width of a character
	messageX := (screenWidth - messageWidth) / 2
	messageY := (screenHeight - infoAreaHeight) / 2
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
}

func drawWaiting(screen *ebiten.Image) {
	const message = "Click to Start"
	messageWidth := len(message) * 6 // 6 is the width of a character
	messageX := (screenWidth - messageWidth) / 2
	messageY := (screenHeight - infoAreaHeight) / 2
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
}

func drawGameClear(screen *ebiten.Image) {
	const message = "Congratulations! Game Clear!"
	messageWidth := len(message) * 6 // 6 is the width of a character
	messageX := (screenWidth - messageWidth) / 2
	messageY := (screenHeight - infoAreaHeight) / 2
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
}

// クリックされたユニットの情報を表示する関数
func (g *Game) drawUnitInfo(screen *ebiten.Image, unit Clickable) {
	const marginBottom = 10
	const sideMargin = 10
	screenHeight := screen.Bounds().Dy()
	infoAreaY := screenHeight - infoAreaHeight - marginBottom // 情報表示領域のY座標
	infoAreaX := sideMargin                                   // 情報表示領域のX座標

	switch u := unit.(type) {
	case *Player:
		ebitenutil.DebugPrintAt(screen, "Player", infoAreaX+sideMargin, infoAreaY+marginBottom)
	case *Enemy:
		ebitenutil.DebugPrintAt(screen, "Enemy", infoAreaX+sideMargin, infoAreaY+marginBottom)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HP: %d", u.HP), infoAreaX+sideMargin, infoAreaY+marginBottom+20) // EnemyのHPを表示
	case *Base:
		ebitenutil.DebugPrintAt(screen, "Base", infoAreaX+sideMargin, infoAreaY+marginBottom)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HP: %d", u.HP), infoAreaX+sideMargin, infoAreaY+marginBottom+20)
	}
}

func (g *Game) drawGame(screen *ebiten.Image) {
	drawMoney(screen, g.money)

	g.player.Draw(screen)
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}
	for _, bullet := range g.playerBullets {
		bullet.Draw(screen)
	}
	for _, bullet := range g.enemyBullets {
		bullet.Draw(screen)
	}
	for _, wall := range g.walls {
		wall.Draw(screen)
	}
	if g.unitInfo != nil {
		g.drawUnitInfo(screen, g.unitInfo)
	}

	g.base.Draw(screen)
}

func drawInfoArea(screen *ebiten.Image) {
	const borderThickness = 2
	const marginBottom = 10
	const sideMargin = 10
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	rect := image.Rect(sideMargin, screenHeight-infoAreaHeight-marginBottom, screenWidth-sideMargin, screenHeight-marginBottom)
	borderColor := color.RGBA{R: 255, G: 255, B: 255, A: 255} // white

	// 上辺
	upperBorder := ebiten.NewImage(rect.Dx(), borderThickness)
	upperBorder.Fill(borderColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
	screen.DrawImage(upperBorder, op)

	// 下辺
	lowerBorder := ebiten.NewImage(rect.Dx(), borderThickness)
	lowerBorder.Fill(borderColor)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.Min.X), float64(rect.Max.Y-borderThickness))
	screen.DrawImage(lowerBorder, op)

	// 左辺
	leftBorder := ebiten.NewImage(borderThickness, rect.Dy())
	leftBorder.Fill(borderColor)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
	screen.DrawImage(leftBorder, op)

	// 右辺
	rightBorder := ebiten.NewImage(borderThickness, rect.Dy())
	rightBorder.Fill(borderColor)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(rect.Max.X-borderThickness), float64(rect.Min.Y))
	screen.DrawImage(rightBorder, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	switch g.gameState {
	case Waiting:
		drawWaiting(screen)
		return
	case GameOver:
		drawGameOver(screen)
	case GameClear:
		drawGameClear(screen)
	}
	g.drawGame(screen)
	drawInfoArea(screen)
}
