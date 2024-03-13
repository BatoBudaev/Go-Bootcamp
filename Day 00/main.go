package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := scan()

	if len(numbers) == 0 {
		fmt.Println("Нет данных.")
		return
	}

	mean := findMean(numbers)
	median := findMedian(numbers)
	mode := findMode(numbers)
	SD := findSD(numbers)

	if len(os.Args) == 1 {
		fmt.Printf("Mean: %.2f\n", mean)
		fmt.Printf("Median: %.2f\n", median)
		fmt.Println("Mode:", mode)
		fmt.Printf("SD: %.2f\n", SD)
	} else {
		args := lowerAndRemoveDuplicates(os.Args[1:])

		for _, arg := range args {
			switch arg {
			case "mean":
				fmt.Printf("Mean: %.2f\n", mean)
			case "median":
				fmt.Printf("Median: %.2f\n", median)
			case "mode":
				fmt.Println("Mode:", mode)
			case "sd":
				fmt.Printf("SD: %.2f\n", SD)
			default:
				fmt.Println("Неверный аргумент. Допустимые аргументы: mean, median, mode, sd.")
				break
			}
		}
	}
}

func scan() []int {
	fmt.Println("Введите числа через Enter. Для завершения ввода введите пустую строку:")

	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		number, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("Ошибка: введите число")
			continue
		}
		if number > 100000 || number < -100000 {
			fmt.Println("Ошибка: число должно быть от -100000 до 100000")
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func findMean(numbers []int) float32 {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return float32(sum) / float32(len(numbers))
}

func findMedian(numbers []int) float32 {
	sort.Ints(numbers)
	mid := len(numbers) / 2

	if len(numbers)%2 == 0 {
		return float32(numbers[mid-1]+numbers[mid]) / 2.0
	}

	return float32(numbers[mid])
}

func findMode(numbers []int) int {
	counts := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		counts[numbers[i]]++
	}

	maxCount := 0
	var mode int

	for num, count := range counts {
		if (count > maxCount) || (count == maxCount && num < mode) {
			mode = num
			maxCount = count
		}
	}

	return mode
}

func findSD(numbers []int) float32 {
	mean := findMean(numbers)
	var sumSquares float32

	for i := 0; i < len(numbers); i++ {
		diff := float32(numbers[i]) - mean
		sumSquares += diff * diff
	}

	variance := sumSquares / float32(len(numbers))
	SD := math.Sqrt(float64(variance))

	return float32(SD)
}

func lowerAndRemoveDuplicates(input []string) []string {
	uniqueStrings := make(map[string]bool)
	var result []string

	for _, s := range input {
		lowerCase := strings.ToLower(s)

		if !uniqueStrings[lowerCase] {
			uniqueStrings[lowerCase] = true
			result = append(result, lowerCase)
		}
	}

	return result
}
