package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Please specify a command")
		return
	}

	command := flag.Args()[0]
	args := flag.Args()[1:]

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			args = append(args, word)
		}
	}

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Print(string(output))
}
