package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"unicode/utf8"
)

var (
	lineFlag = flag.Bool("l", false, "Count lines")
	charFlag = flag.Bool("m", false, "Count characters")
	wordFlag = flag.Bool("w", false, "Count words")
)

func parseFlags() error {
	flag.Parse()

	if !*lineFlag && !*charFlag && !*wordFlag {
		*wordFlag = true
	}

	if *lineFlag && (*charFlag || *wordFlag) || *charFlag && (*lineFlag || *wordFlag) || *wordFlag && (*lineFlag || *charFlag) {
		return fmt.Errorf("only one of -l, -m, or -w can be specified")
	}

	return nil
}

func main() {
	err := parseFlags()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	var wg sync.WaitGroup
	for _, filename := range flag.Args() {
		wg.Add(1)
		go processFile(filename, &wg)
	}

	wg.Wait()
}

func processFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %s\n", filename, err)
		return
	}

	s := string(data)

	switch {
	case *lineFlag:
		lines := 0
		for _, n := range data {
			if n == byte('\n') {
				lines++
			}
		}
		fmt.Printf("%d\t%s\n", lines, filename)

	case *charFlag:
		chars := utf8.RuneCountInString(s) - 1
		fmt.Printf("%d\t%s\n", chars, filename)

	case *wordFlag:
		words := 0
		inWord := false
		for _, r := range s {
			if r == ' ' || r == '\n' {
				inWord = false
			} else if !inWord {
				inWord = true
				words++
			}
		}

		if inWord {
			words++
		}

		fmt.Printf("%d\t%s\n", words, filename)
	}
}
