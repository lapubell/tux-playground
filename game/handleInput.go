package game

func (g *Game) handleInput() {
	if len(g.keys) == 0 {
		tux.isWalking = false
		return
	}

	tux.isWalking = false

	for _, key := range g.keys {
		for _, walkingKey := range walkingKeys {
			if key.String() == walkingKey {
				tux.isWalking = true
			}
		}

		if key.String() == "ArrowUp" && tux.positionY > 0 {
			tux.facingDirection = "U"
			tux.positionY -= tuxSpeed
		}
		if key.String() == "ArrowDown" && tux.positionY < screenHeight-frameHeight {
			tux.facingDirection = "D"
			tux.positionY += tuxSpeed
		}
		if key.String() == "ArrowRight" && tux.positionX < screenWidth-frameWidth {
			tux.facingDirection = "R"
			tux.positionX += tuxSpeed
		}
		if key.String() == "ArrowLeft" && tux.positionX > 0 {
			tux.facingDirection = "L"
			tux.positionX -= tuxSpeed
		}
	}
}
