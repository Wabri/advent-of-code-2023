package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	totalCalibration := 0
	firstDigit := -1
	lastDigit := -1

	for {
		char, _, err := scanner.ReadRune()
		if err != nil {
			break
		}

		if unicode.IsDigit(char) {
			if firstDigit == -1 {
				firstDigit = int(char) - 48
			} else {
				lastDigit = int(char) - 48
			}
		} else if char == '\n' {
			if lastDigit == -1 {
				lastDigit = firstDigit
			}
			totalCalibration = totalCalibration + firstDigit*10 + lastDigit
			firstDigit = -1
			lastDigit = -1
		}
	}

	fmt.Println(totalCalibration)

}
