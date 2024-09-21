package hangman

import (
	"fmt"
)

type Player struct {
	CountAttempts int
	Window        *Windows
}

func (player *Player) ChooseLevel() (level int) {
	if _, err := fmt.Scan(&level); err != nil {
		player.Window.SelectLevel()
		return player.ChooseLevel()
	}

	return level
}

func (player *Player) ChooseCategory() (category int) {
	if _, err := fmt.Scan(&category); err != nil {
		player.Window.SelectCategory()
		return player.ChooseCategory()
	}

	return category
}
