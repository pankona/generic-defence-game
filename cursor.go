package main

import "github.com/hajimehoshi/ebiten/v2"

func getPointerPosition() (x, y float64, eventOccurred bool) {
	// マウスクリックの処理
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		return float64(mx), float64(my), true
	}

	// タッチイベントの処理
	touchIDs := ebiten.AppendTouchIDs(nil)
	if len(touchIDs) > 0 {
		tx, ty := ebiten.TouchPosition(touchIDs[0])
		return float64(tx), float64(ty), true
	}

	// イベントが発生しなかった場合
	return 0, 0, false
}
