package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input_2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalCalibration := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstDigit := extractFrom(line)
		line = reverseButKeepOrder(line)
		lastDigit := extractFrom(line)
		// lastDigit := 0
		// if len(line) > 1 {
		// 	lastDigit = extractFrom(line)
		// } else {
		// 	lastDigit = firstDigit
		// 	firstDigit = 0
		// }
		totalCalibration += firstDigit*10 + lastDigit
	}

	fmt.Println(totalCalibration)
}

func reverseButKeepOrder(word string) string {
	var alphaPart string
	result := ""
	for _, char := range word {
		if unicode.IsLetter(char) {
			alphaPart += string(char)
			digit := extractDigit(alphaPart)
			if digit != -1 {
				result = strconv.Itoa(digit) + result
				alphaPart = alphaPart[1:]
			}
		} else if unicode.IsDigit(char) {
			digit := extractDigit(alphaPart)
			if digit != -1 {
				result = strconv.Itoa(digit) + result
			}
			result = string(char) + result
			alphaPart = ""
		}
	}

	return result
}

func extractFrom(line string) int {
	characters := ""
	for _, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0')
		} else {
			characters += string(char)
		}
		if len(characters) >= 3 {
			digit := 0
			digit = extractDigit(characters)
			if digit != -1 {
				return digit
			}
		}
	}
	return -1
}

func extractDigit(word string) int {
	for i := 3; i <= len(word) && i <= 5; i++ {
		isDigit, digit := isADigit(word[0:i])
		if isDigit {
			return digit
		}
	}
	if len(word) < 3 {
		return -1
	}
	return extractDigit(word[1:])
}

func isADigit(word string) (bool, int) {
	switch word {
	case "one":
		return true, 1
	case "two":
		return true, 2
	case "three":
		return true, 3
	case "four":
		return true, 4
	case "five":
		return true, 5
	case "six":
		return true, 6
	case "seven":
		return true, 7
	case "eight":
		return true, 8
	case "nine":
		return true, 9
	}
	return false, -1
}
