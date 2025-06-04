package domain

import (
	"fmt"
	"strings"
	"time"

	output "hangman/internal/infrastructure"
)

const strikethroughtText string = "а̶б̶в̶г̶д̶е̶ж̶з̶и̶й̶к̶л̶м̶н̶о̶п̶р̶с̶т̶у̶ф̶х̶ц̶ч̶ш̶щ̶ъ̶ы̶ь̶э̶ю̶я̶"

var Stages = []string{
	`






  ________________`,
	`             ||
             ||
             ||
             ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾‾‾‾‾‾‾‾‾‾||
             ||
             ||
             ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
   |         ||
             ||
             ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
             ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   /         ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   / \       ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   /|\       ||
             ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   /|\       ||
    |        ||
             ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   /|\       ||
    |        ||
   /         ||
             ||
  ___________||___`,
	`   ‾|‾‾‾‾‾‾‾‾||
    |        ||
    0        ||
   /|\       ||
    |        ||
   / \       ||
             ||
  ___________||___`,
}

type Windows struct {
	HangmanStages []string
	Cursor        Cursor
}

func (*Windows) CleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func (window *Windows) Start() {
	window.CleanScreen()
	window.Cursor.HideCursor()
	output.PrintWindow("start.txt")
	time.Sleep(time.Second)
}

func (window *Windows) SelectLevel() {
	window.CleanScreen()
	output.PrintWindow("level.txt")
	window.Cursor.ToChoose()
	window.Cursor.ChangeCursor()
	window.Cursor.ShowCursor()
}

func (window *Windows) SelectCategory() {
	window.CleanScreen()
	output.PrintWindow("category.txt")
	window.Cursor.ToChoose()
}

func (window *Windows) MainWindow() {
	window.CleanScreen()
	output.PrintWindow("game.txt")
}

func (window *Windows) DrawHangman(countAttempts int) {
	window.Cursor.ToHangman()

	currentStage := Stages[len(Stages)-countAttempts-1]
	lines := strings.Split(currentStage, "\n")

	for index, line := range lines {
		fmt.Print(line)
		window.Cursor.ToHangman()
		window.Cursor.Down(rune(index + 1))
	}
}

func (window *Windows) CrossOutLetter(char rune) {
	window.Cursor.ToAlphabet()

	const countRow rune = 7

	letterNumber := (char - 'а')

	if letterNumber >= 0 && letterNumber <= 32 {
		if letterNumber%countRow == 0 {
			window.Cursor.Down((letterNumber/countRow)*2 + 1)
		} else {
			window.Cursor.Down((letterNumber/countRow)*2 + 1)
			window.Cursor.Right((letterNumber % countRow) * 2)
		}

		fmt.Print(string([]rune(strikethroughtText)[letterNumber*2 : letterNumber*2+2]))
	}
}

func (window *Windows) ShowHint(hint string) {
	window.Cursor.ToHint("? - подсказка")
	window.ClearText("? - подсказка")
	window.Cursor.ToHint(hint)
	fmt.Print(hint)
	window.Cursor.ToInput()
	time.Sleep(2 * time.Second)
	window.Cursor.ToHint(hint)
	window.ClearText(hint)
	window.Cursor.ToHint("? - подсказка")
	fmt.Print("? - подсказка")
}

func (*Windows) ClearText(text string) {
	for i := 0; i < len([]rune(text)); i++ {
		fmt.Print(" ")
	}
}

func (window *Windows) Win() {
	window.CleanScreen()
	output.PrintWindow("win.txt")
}

func (window *Windows) Loss() {
	window.CleanScreen()
	output.PrintWindow("loss.txt")
}
