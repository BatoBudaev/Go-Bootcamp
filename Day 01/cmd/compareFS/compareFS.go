package main

import (
	"fmt"
	"os"
	"s21/internal/fileSystem"
)

func main() {
	if len(os.Args) == 5 && os.Args[1] == "--old" && os.Args[3] == "--new" {
		err := fileSystem.CompareFiles(os.Args[2], os.Args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Пример использования: ./compareFS --old <file_path> --new <file_path>")
	}
}
