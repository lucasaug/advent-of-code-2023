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
	firstDigit := '\x00'
	lastDigit := '\x00'

	for _, c := range s{
		if unicode.IsNumber(c) {
			if firstDigit == '\x00' {
				firstDigit = c
			}
			lastDigit = c
		}
	}

	fullNumber := string(firstDigit) + string(lastDigit)

	i, err := strconv.Atoi(fullNumber)

	if err != nil {
		panic(err)
	}

	return i
}

func ConvertNumberNamesToValue(s string) string {
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	return replacer.Replace(s)
}

func Rev(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		converted := ConvertNumberNamesToValue(text)
		fmt.Println(converted)
		num := ExtractNumber(converted)
		fmt.Println(num)
		sum = sum + num
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
