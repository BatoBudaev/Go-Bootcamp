package main

import (
	"fmt"
	"os"
	"s21/internal/recipes"
)

func main() {
	if len(os.Args) == 3 && os.Args[1] == "-f" {
		reader, data := recipes.ReadFile(os.Args[2])
		reader.PrintData(data)
	} else {
		fmt.Println("Пример использования: ./read -f <file_path>")
	}
}
