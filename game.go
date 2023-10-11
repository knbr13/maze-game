package main

func (b *board) checkWin() bool {
	if b.playerXPos == b.gateXPos && b.playerYPos == b.gateYPos {
		return true
	}
	return false
}
