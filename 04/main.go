package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func GameScore(s string) int {
    count := 0

    splitGame := strings.Split(s, ": ")
    splitGame = strings.Split(splitGame[1], " | ")

    winningNumbers := strings.Fields(splitGame[0])
    scratchNumbers := strings.Fields(splitGame[1])

    for _, num := range scratchNumbers {
        if slices.Contains(winningNumbers, num) {
            count++
        }
    }

    // return int(math.Pow(2, float64(count - 1)))
    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    extraCopies := []int {}
    for scanner.Scan() {
        line := scanner.Text()
        currentCopies := 1
        if len(extraCopies) > 0 {
            currentCopies += extraCopies[0]
            extraCopies = extraCopies[1:]
        }
        sum += currentCopies

        score := GameScore(line)
        if len(extraCopies) < score {
            newEntries := make([]int, score - len(extraCopies))
            extraCopies = append(extraCopies, newEntries...)
        }

        for i := range extraCopies[:score] {
            extraCopies[i] += currentCopies
        }

    }

    fmt.Println(sum)
}
