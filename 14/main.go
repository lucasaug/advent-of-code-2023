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

func Cycle(puzzle [][]rune) [][]rune {
    result := [][]rune{}
    for _, row := range puzzle {
        rowCopy := append([]rune{}, row...)
        result = append(result, rowCopy)
    }

    for k := 0; k <= 4; k++ {
        scanLength := len(puzzle[0])
        orthogonalLength := len(puzzle)
        if k % 2 == 1 {
            scanLength, orthogonalLength = orthogonalLength, scanLength
        }

        lastAvailablePosition := make([]int, scanLength)
        if k == 2 || k == 3 {
            for i := range lastAvailablePosition {
                lastAvailablePosition[i] = orthogonalLength - 1
            }
        }

        for i := 0; i < orthogonalLength; i++ {
            for j := 0; j < scanLength; j++ {
                x, y := i, j
                maxX, maxY := len(result), len(result[i])

                if k % 2 == 1 {
                    x, y = y, x
                    maxX, maxY = maxY, maxX
                }

                if k == 2 || k == 3 {
                    x = maxX - 1 - x
                    y = maxY - 1 - y
                }

                value := result[x][y]
                if value == 'O' {
                    result[x][y] = '.'
                    result[lastAvailablePosition[y]][y] = 'O'

                    if k == 2 || k == 3 {
                        lastAvailablePosition[y]--
                    } else {
                        lastAvailablePosition[y]++
                    }

                }

                if value == '#' {
                    if (k == 2 || k == 3) && x > 0 {
                        lastAvailablePosition[y] = x - 1
                    } else if x < maxX - 1 {
                        lastAvailablePosition[y] = x + 1
                    }
                }
            }
        }
    }

    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    puzzle := [][]rune{}
    for scanner.Scan() {
        line := scanner.Text()
        puzzle = append(puzzle, []rune(line))
    }

    // fmt.Println(CalculateLoad(puzzle))
    result := Cycle(puzzle)
    for _, row := range result {
        fmt.Println(string(row))
    }

}

