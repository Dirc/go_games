package main

func setState(score1 int, score2 int) gameState {
	if score1 == 3 || score2 == 3 {
		return gameover
	} else {
		return start
	}
}
