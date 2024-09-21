package output

import (
	"fmt"
	"log"
	"os"
)

func PrintWindow(nameFile string) {
	bytes, err := os.ReadFile(fmt.Sprintf("internal/domain/GameWindows/%s", nameFile))
	if err != nil {
		log.Fatal(err)
	}

	fileText := string(bytes)
	fmt.Println(fileText)
}
