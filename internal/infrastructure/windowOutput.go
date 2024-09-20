package output

import (
	"fmt"
	"log"
	"os"
)

func PrintWindow(name_file string) {
	bytes, err := os.ReadFile(fmt.Sprintf("internal/domain/GameWindows/%s", name_file))
	if err != nil {
		log.Fatal(err)
	}
	file_text := string(bytes)
	fmt.Println(file_text)
}
