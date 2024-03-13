package main

import (
	"fmt"
	"os"
	"s21/internal/recipes"
)

func main() {
	if len(os.Args) == 5 && os.Args[1] == "--old" && os.Args[3] == "--new" {
		_, data1 := recipes.ReadFile(os.Args[2])
		_, data2 := recipes.ReadFile(os.Args[4])
		recipes.CompareDB(*data1, *data2)
	} else {
		fmt.Println("Пример использования: ./compareDB --old <file_path> --new <file_path>")
	}
}
