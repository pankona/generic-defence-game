package main

import (
	"fmt"
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player         Player
	enemies        []Enemy
	bullets        []Bullet
	gameState      string
	spawnInterval  int
	spawnedEnemies int

	maxEnemies int

	isDragging     bool
	startX, startY float64
	walls          []Wall

	reachedEnemies int
}

const (
	Waiting   = "waiting"
	Playing   = "playing"
	GameOver  = "gameover"
	GameClear = "gameclear"
)

func (g *Game) Update() error {
	// ゲーム開始待機状態で左クリックが押された場合、ゲームを開始
	if g.gameState == Waiting && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.gameState = Playing
		return nil
	}

	// ゲームオーバーの状態で左クリックが押された場合、ゲームをリセット
	if g.gameState == GameOver && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g = NewGame()
		return nil
	}

	if g.gameState == Playing {
		g.UpdateGame()
	}

	// 壁の生成に関するロジック
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.isDragging {
			g.isDragging = true
			x, y := ebiten.CursorPosition()
			g.startX, g.startY = float64(x), float64(y)
		}
	} else if g.isDragging {
		x, y := ebiten.CursorPosition()
		endX, endY := float64(x), float64(y)
		g.walls = append(g.walls, Wall{id: uuid.New().String(), x1: g.startX, y1: g.startY, x2: endX, y2: endY})
		g.isDragging = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	if g.gameState == Waiting {
		DrawWaiting(screen)
	}

	if g.gameState == Playing {
		g.DrawGame(screen)
	}

	if g.gameState == GameOver {
		DrawGameOver(screen)
	}

	if g.gameState == GameClear {
		DrawGameClear(screen)
	}

	for _, wall := range g.walls {
		wall.Draw(screen)
	}
}

func DrawGameClear(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Congratulations! Game Clear!")
}

func (g *Game) UpdateGame() {
	// 敵の生成
	g.spawnInterval++
	if g.spawnInterval > 100 && g.spawnedEnemies < g.maxEnemies {
		g.enemies = append(g.enemies, Enemy{x: 0, y: 0, speed: 2, HP: 3, active: true})
		g.spawnInterval = 0
		g.spawnedEnemies++
	}

	// ゲームクリアの判定
	if len(g.enemies) == 0 && g.spawnedEnemies == g.maxEnemies {
		if g.reachedEnemies < 3 {
			g.gameState = GameClear
		}
	}

	// プレイヤーが敵に近づいたら攻撃するロジック
	for i := range g.enemies {
		enemy := &g.enemies[i]

		if enemy.reached {
			g.reachedEnemies++
			enemy.reached = false
			enemy.active = false
		}

		distX := g.player.x - enemy.x
		distY := g.player.y - enemy.y
		distance := distX*distX + distY*distY

		attackRange := 100.0 * 100.0
		if distance < attackRange && g.player.framesSinceLastBullet >= g.player.bulletFrameInterval {
			// 弾を発射する
			bullet := NewBullet(g.player.x, g.player.y, enemy)
			g.bullets = append(g.bullets, bullet)

			g.player.framesSinceLastBullet = 0
		}
	}

	g.player.Update()
	for i := range g.enemies {
		g.enemies[i].Update(g.walls)
	}

	// 弾の更新と敵との当たり判定
	for i := range g.bullets {
		bullet := &g.bullets[i]
		bullet.Update()
		for i := range g.enemies {
			enemy := &g.enemies[i]
			if bullet.active && enemy.IsHit(bullet.x, bullet.y) {
				bullet.active = false
				enemy.HP -= g.player.attack
				if enemy.HP <= 0 {
					enemy.active = false
				}
			}
		}
	}

	// 無効になった敵を削除
	activeEnemies := g.enemies[:0]
	for _, enemy := range g.enemies {
		if enemy.active {
			activeEnemies = append(activeEnemies, enemy)
		}
	}
	g.enemies = activeEnemies

	// 無効になった弾を削除
	activeBullets := g.bullets[:0]
	for _, bullet := range g.bullets {
		if bullet.active {
			activeBullets = append(activeBullets, bullet)
		}
	}
	g.bullets = activeBullets
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	// reachedEnemies の値を表示
	debugText := fmt.Sprintf("Enemies Reached: %d", g.reachedEnemies)
	ebitenutil.DebugPrint(screen, debugText)

	g.player.Draw(screen)
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func DrawGameOver(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Game Over")
}

func DrawWaiting(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Click to Start")
}

func NewGame() *Game {
	arrow := ebiten.NewImage(16, 16)
	arrow.Fill(color.White)
	return &Game{
		player: Player{
			x:                   320,
			y:                   240,
			targetX:             320,
			targetY:             240,
			speed:               4,
			attack:              1,
			arrow:               arrow,
			bulletFrameInterval: 30,
		},
		maxEnemies: 10,
		gameState:  Waiting,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Generic Shooting Game")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
