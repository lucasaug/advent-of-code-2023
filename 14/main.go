package main

import (
	"bufio"
	"fmt"
	"os"
)

func CalculateLoad(puzzle [][]rune) int {
    lastAvailablePosition := make([]int, len(puzzle[0]))

    sum := 0
    for i := range puzzle {
        for j, value := range puzzle[i] {
            if value == 'O' {
                sum += len(puzzle) - lastAvailablePosition[j] 
                lastAvailablePosition[j]++
            }
            if value == '#' && i < len(puzzle) - 1 {
                lastAvailablePosition[j] = i + 1
            }
        }
    }

    return sum
}

func RotateMatrix(slice [][]rune) [][]rune {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]rune, xl)

    for i := range result {
        result[i] = make([]rune, yl)
    }

    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][yl - j - 1] = slice[j][i]
        }
    }

    return result
}

func Cycle(puzzle [][]rune) [][]rune {
    result := [][]rune{}
    for _, row := range puzzle {
        rowCopy := append([]rune{}, row...)
        result = append(result, rowCopy)
    }

    for k := 0; k < 4; k++ {
        lastAvailablePosition := make([]int, len(result[0]))

        for i := range result {
            for j, value := range result[i] {
                if value == 'O' {
                    result[i][j] = '.'
                    result[lastAvailablePosition[j]][j] = 'O'
                    lastAvailablePosition[j]++
                }
                if value == '#' && i < len(result) - 1 {
                    lastAvailablePosition[j] = i + 1
                }
            }
        }
        result = RotateMatrix(result)
    }

    return result
}

func MatrixEqual(a [][]rune, b [][]rune) bool {
    for i, row := range a {
        for j, value := range row {
            if value != b[i][j] {
                return false
            }
        }
    }

    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    puzzle := [][]rune{}
    for scanner.Scan() {
        line := scanner.Text()
        puzzle = append(puzzle, []rune(line))
    }

    // fmt.Println(CalculateLoad(puzzle))
    previous := puzzle
    result := Cycle(puzzle)

    for !MatrixEqual(previous, result) {
        previous = result
        result = Cycle(result)
    }

    fmt.Println(CalculateLoad(result))
}

