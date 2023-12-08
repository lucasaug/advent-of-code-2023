package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ExtractNumber(s string) int {
	numbers := [9]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	firstDigit := 0
	lastDigit := 0

	for strIndex, c := range s {
		if unicode.IsNumber(c) {
			if firstDigit == 0 {
				firstDigit, _ = strconv.Atoi(string(c))
			}
			lastDigit, _ = strconv.Atoi(string(c))
		} else {
			for i, num := range numbers {
				if strings.HasPrefix(s[strIndex:], num) {
					if firstDigit == 0 {
						firstDigit = i + 1
					}
					lastDigit = i + 1
				}
			}
		}
	}

	return firstDigit*10 + lastDigit
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		num := ExtractNumber(text)
		sum = sum + num
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
