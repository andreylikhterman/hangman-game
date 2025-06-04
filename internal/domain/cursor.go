package domain

import "fmt"

type Cursor struct{}

func (*Cursor) ToChoose() {
	fmt.Print("\033[H")
	fmt.Print("\033[13B\033[29C")
}

func (*Cursor) ToInput() {
	fmt.Print("\033[H")
	fmt.Print("\033[2B\033[18C")
}

func (*Cursor) ToAttempts() {
	fmt.Print("\033[H")
	fmt.Print("\033[15B\033[21C")
}

func (*Cursor) ToHangman() {
	fmt.Print("\033[H")
	fmt.Print("\033[5B\033[20C")
}

func (*Cursor) ToAlphabet() {
	fmt.Print("\033[H")
	fmt.Print("\033[3B\033[43C")
}

func (*Cursor) ToWord() {
	fmt.Print("\033[H")
	fmt.Print("\033[8B\033[3C")
}

func (*Cursor) ToHint(hint string) {
	fmt.Print("\033[H")
	fmt.Printf("\033[15B\033[%dC", 56-len([]rune(hint)))
}

func (*Cursor) HideCursor() {
	fmt.Print("\033[?25l")
}

func (*Cursor) ShowCursor() {
	fmt.Print("\033[?25h")
}

func (*Cursor) ChangeCursor() {
	fmt.Print("\033[6 q")
}

func (*Cursor) Down(count rune) {
	fmt.Printf("\033[%dB", count)
}

func (*Cursor) Right(count rune) {
	fmt.Printf("\033[%dC", count)
}
