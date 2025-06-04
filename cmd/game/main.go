package main

import (
	hangman "hangman/internal/application"
)

func main() {
	attempts := 6

	game := hangman.NewGame(attempts)
	game.Play()
}
