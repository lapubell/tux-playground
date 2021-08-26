package game

func (g *Game) updateSprites() {
	frames := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1}
	if tux.facingDirection == "D" {
		if tux.isWalking {
			tux.frameNumber = frames[g.count%len(frames)]
		} else {
			tux.frameNumber = frames[0]
		}
		return
	}

	if tux.facingDirection == "U" {
		frames = []int{3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
		if tux.isWalking {
			tux.frameNumber = frames[g.count%len(frames)]
		} else {
			tux.frameNumber = frames[0]
		}
		return
	}

	if tux.facingDirection == "R" {
		frames = []int{7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9}
		if tux.isWalking {
			tux.frameNumber = frames[g.count%len(frames)]
		} else {
			tux.frameNumber = frames[0]
		}
		return
	}

	if tux.facingDirection == "L" {
		frames = []int{10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 10, 10, 10, 10, 10, 12, 12, 12, 12, 12}
		if tux.isWalking {
			tux.frameNumber = frames[g.count%len(frames)]
		} else {
			tux.frameNumber = frames[0]
		}
		return
	}
}
