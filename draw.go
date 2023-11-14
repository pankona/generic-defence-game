package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

const marginBottom = 10
const sideMargin = 10
const infoAreaY = screenHeight - infoAreaHeight - marginBottom // 情報表示領域のY座標
const infoAreaX = sideMargin

// クリックされたユニットの情報を表示する関数
func (g *Game) drawUnitInfo(screen *ebiten.Image, unit Clickable) {
	// 情報表示領域のX座標
	switch u := unit.(type) {
	case *Player:
		ebitenutil.DebugPrintAt(screen, "Player", infoAreaX+sideMargin, infoAreaY+marginBottom)
	case *Enemy:
		ebitenutil.DebugPrintAt(screen, "Enemy", infoAreaX+sideMargin, infoAreaY+marginBottom)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HP: %d", u.HP), infoAreaX+sideMargin, infoAreaY+marginBottom+20) // EnemyのHPを表示
	case *Base:
		g.drawBaseInfo(screen)
	}

	for _, button := range g.unitInfoPanel.buttons {
		x, y := int(button.x), int(button.y)
		drawRectBorder(screen, x, y, 100, infoAreaHeight-10, color.White)
		for _, text := range button.text {
			ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
			y += 20
		}
	}
}

func (g *Game) drawBaseInfo(screen *ebiten.Image) { // 情報表示領域のX座標
	ebitenutil.DebugPrintAt(screen, "Base", infoAreaX+sideMargin, infoAreaY+marginBottom)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HP: %d", g.base.HP), infoAreaX+sideMargin, infoAreaY+marginBottom+20)

	/*
		recoverButton := &RecoverButton{
			x:      infoAreaX + sideMargin + 100,
			y:      infoAreaY + 5,
			width:  100,
			height: infoAreaHeight - 10,
		}
		drawRectBorder(screen, recoverButton.x, recoverButton.y, recoverButton.width, recoverButton.height, color.White)
		ebitenutil.DebugPrintAt(screen, "Recover HP", recoverButton.x+10, recoverButton.y+10)
		ebitenutil.DebugPrintAt(screen, "+10HP / $10", recoverButton.x+10, recoverButton.y+30)
	*/
}

func drawRectBorder(screen *ebiten.Image, x, y, width, height int, clr color.Color) {
	const thickness = 1.0
	vector.StrokeLine(screen, float32(x), float32(y), float32(x+width), float32(y), thickness, clr, false)
	vector.StrokeLine(screen, float32(x+width), float32(y), float32(x+width), float32(y+height), thickness, clr, false)
	vector.StrokeLine(screen, float32(x+width), float32(y+height), float32(x), float32(y+height), thickness, clr, false)
	vector.StrokeLine(screen, float32(x), float32(y+height), float32(x), float32(y), thickness, clr, false)
}

func (g *Game) drawGame(screen *ebiten.Image) {
	drawMoney(screen, g.money)

	for _, player := range g.players {
		player.Draw(screen)
	}
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
