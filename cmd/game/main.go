package main

import (
	hangman "github.com/backend-academy-2024-go-template/internal/application"
)

func main() {
	game := hangman.NewHangmanGame(6)
	game.Play()
}
