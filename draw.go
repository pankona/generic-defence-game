package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func drawBaseHP(screen *ebiten.Image, hp int) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Base HP: %d", hp), 0, 10)
}

func drawMoney(screen *ebiten.Image, money int) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Money: %d", money), screenWidth-100, 10)
}

func drawGameOver(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Game Over", screenWidth/2, screenHeight/2)
}

func drawWaiting(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Click to Start", screenWidth/2, screenHeight/2)
}

// クリックされたユニットの情報を表示する関数
func (g *Game) drawUnitInfo(screen *ebiten.Image, unit Clickable) {
	switch unit.(type) {
	case *Player:
		ebitenutil.DebugPrintAt(screen, "Player", 0, 30)
	case *Enemy:
		ebitenutil.DebugPrintAt(screen, "Enemy", 0, 30)
	case *Base:
		ebitenutil.DebugPrintAt(screen, "Base", 0, 30)
	}
}

func drawGameClear(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Congratulations! Game Clear!", screenWidth/2, screenHeight/2)
}

func (g *Game) drawGame(screen *ebiten.Image) {
	drawBaseHP(screen, g.base.HP)
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
}
