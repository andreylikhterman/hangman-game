package hangman

import (
	"fmt"
)

type Player struct {
	CountAttempts int
}

func (player *Player) ChooseLevel() string {
	var level string
	fmt.Scan(&level)
	return level

}

func (player *Player) ChooseCategory() string {
	var category string
	fmt.Scan(&category)
	return category
}
