package hangman

import (
	"fmt"
	"strings"
	"time"

	output "github.com/backend-academy-2024-go-template/internal/infrastructure"
)

const strikethrought_text string = "а̶б̶в̶г̶д̶е̶ж̶з̶и̶й̶к̶л̶м̶н̶о̶п̶р̶с̶т̶у̶ф̶х̶ц̶ч̶ш̶щ̶ъ̶ы̶ь̶э̶ю̶я̶"

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

type cursor struct{}

func (*cursor) ToChoose() {
	fmt.Print("\033[H")
	fmt.Print("\033[13B\033[29C")
}

func (*cursor) ToInput() {
	fmt.Print("\033[H")
	fmt.Print("\033[2B\033[18C")
}

func (*cursor) ToAttempts() {
	fmt.Print("\033[H")
	fmt.Print("\033[15B\033[21C")
}

func (*cursor) ToHangman() {
	fmt.Print("\033[H")
	fmt.Print("\033[5B\033[20C")
}

func (*cursor) ToAlphabet() {
	fmt.Print("\033[H")
	fmt.Print("\033[3B\033[43C")
}

func (*cursor) ToWord() {
	fmt.Print("\033[H")
	fmt.Print("\033[8B\033[3C")
}

func (*cursor) ToHint(hint string) {
	fmt.Print("\033[H")
	fmt.Printf("\033[15B\033[%dC", 56-len([]rune(hint)))
}

func (*cursor) HideCursor() {
	fmt.Print("\033[?25l")
}

func (*cursor) ShowCursor() {
	fmt.Print("\033[?25h")
}

func (*cursor) ChangeCursor() {
	fmt.Print("\033[6 q")
}

func (*cursor) Down(count rune) {
	fmt.Printf("\033[%dB", count)
}

func (*cursor) Right(count rune) {
	fmt.Printf("\033[%dC", count)
}

type HangmanWindows struct {
	HangmanStages []string
	Cursor        cursor
}

func (*HangmanWindows) CleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func (hangman_windows *HangmanWindows) Start() {
	hangman_windows.CleanScreen()
	hangman_windows.Cursor.HideCursor()
	output.PrintWindow("start.txt")
	time.Sleep(time.Second)
}

func (hangman_windows *HangmanWindows) SelectLevel() {
	hangman_windows.CleanScreen()
	output.PrintWindow("level.txt")
	hangman_windows.Cursor.ToChoose()
	hangman_windows.Cursor.ChangeCursor()
	hangman_windows.Cursor.ShowCursor()
}

func (hangman_windows *HangmanWindows) SelectCategory() {
	hangman_windows.CleanScreen()
	output.PrintWindow("category.txt")
	hangman_windows.Cursor.ToChoose()
}

func (hangman_windows *HangmanWindows) MainWindow() {
	hangman_windows.CleanScreen()
	output.PrintWindow("game.txt")
}

func (hangman_windows *HangmanWindows) DrawHangman(count_attemts int) {
	hangman_windows.Cursor.ToHangman()
	current_stage := Stages[len(Stages)-count_attemts-1]
	lines := strings.Split(current_stage, "\n")
	for index, line := range lines {
		fmt.Print(line)
		hangman_windows.Cursor.ToHangman()
		hangman_windows.Cursor.Down(rune(index + 1))
	}
}

func (hangman_windows *HangmanWindows) CrossOutLetter(char rune) {
	hangman_windows.Cursor.ToAlphabet()
	const count_row rune = 7
	letter_number := (char - 'а')
	if letter_number >= 0 && letter_number <= 32 {
		if letter_number%count_row == 0 {
			hangman_windows.Cursor.Down((letter_number/count_row)*2 + 1)
		} else {
			hangman_windows.Cursor.Down((letter_number/count_row)*2 + 1)
			hangman_windows.Cursor.Right((letter_number % count_row) * 2)
		}
		fmt.Print(string([]rune(strikethrought_text)[letter_number*2 : letter_number*2+2]))
	}
}

func (window *HangmanWindows) ShowHint(hint string) {
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

func (*HangmanWindows) ClearText(text string) {
	for i := 0; i < len([]rune(text)); i++ {
		fmt.Print(" ")
	}
}

func (hangman_windows *HangmanWindows) Win() {
	hangman_windows.CleanScreen()
	output.PrintWindow("win.txt")
}

func (hangman_windows *HangmanWindows) Loss() {
	hangman_windows.CleanScreen()
	output.PrintWindow("loss.txt")
}
