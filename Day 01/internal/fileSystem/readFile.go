package fileSystem

import (
	"bufio"
	"os"
)

func ReadFile(filePath string) (map[string]struct{}, error) {
	fileMap := make(map[string]struct{})

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileMap[scanner.Text()] = struct{}{}
	}

	return fileMap, nil
}
