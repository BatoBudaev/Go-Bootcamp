package fileSystem

import (
	"bufio"
	"fmt"
	"os"
)

func CompareFiles(oldFilePath, newFilePath string) error {
	oldFileMap, err := ReadFile(oldFilePath)
	if err != nil {
		return err
	}

	newFile, err := os.Open(newFilePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)

	for scanner.Scan() {
		line := scanner.Text()
		_, isExists := oldFileMap[line]
		if isExists {
			delete(oldFileMap, line)
		} else {
			fmt.Println("ADDED", line)
		}
	}

	for line := range oldFileMap {
		fmt.Println("REMOVED", line)
	}

	return nil
}
