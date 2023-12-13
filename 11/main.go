package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const EXPANSION_LENGTH = 1000000

type Galaxy struct {
    row, col int
}

func ReadUniverse(scanner *bufio.Scanner) []Galaxy {
    var result []Galaxy
    var emptyCols []bool

    rowCount := 0
    for scanner.Scan() {
        line := scanner.Text()

        if len(emptyCols) == 0 {
            lineLength := len(line)
            for i := 0; i < lineLength; i++ {
                emptyCols = append(emptyCols, true)
            }
        }

        isEmpty := true
        for col, char := range line {
            if char == '#' {
                isEmpty = false
                result = append(result, Galaxy { rowCount, col })
                emptyCols[col] = false
            }
        }

        if isEmpty {
            rowCount += EXPANSION_LENGTH
        } else {
            rowCount++
        }
    }

    var colOffsets []int
    currOffset := 0
    for _, value := range emptyCols {
        colOffsets = append(colOffsets, currOffset)

        if value {
            currOffset += EXPANSION_LENGTH - 1
        }
    }

    for i := range result {
        result[i].col += colOffsets[result[i].col]
    }

    return result
}

func Distance(a, b Galaxy) int {
    return int(math.Abs(float64(a.row - b.row))) +
        int(math.Abs(float64(a.col - b.col)))
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    universe := ReadUniverse(scanner)

    sum := 0
    for i, galaxy := range universe[:len(universe)-1] {
        for _, pair := range universe[i+1:] {
            sum += Distance(galaxy, pair)
        }
    }

    fmt.Println(sum)
}
