package domain

import (
	"fmt"
)

type Player struct {
	CountAttempts int
}

func (player *Player) ChooseLevel() (level int, err error) {
	_, err = fmt.Scan(&level)
	return level, err
}

func (player *Player) ChooseCategory() (category int, err error) {
	_, err = fmt.Scan(&category)
	return category, err
}
