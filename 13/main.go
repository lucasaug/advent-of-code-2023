package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPuzzle(scanner *bufio.Scanner) []string {
    puzzle := []string{}

    for scanner.Scan() {
        line := scanner.Text()

        if len(line) == 0 {
            return puzzle
        }

        puzzle = append(puzzle, line)
    }

    return puzzle
}

func HorizontalReflection(puzzle []string, hasSmudge bool) int {
    for i := 1; i < len(puzzle); i++ {
        j := 0
        smudgeUsed := !hasSmudge

        for ; i + j < len(puzzle) && i - j - 1 >= 0; j++ {
            numDiffs := 0

            for k := 0; k < len(puzzle[0]); k++ {
                if puzzle[i + j][k] != puzzle[i - j - 1][k] {
                    numDiffs++
                }
            }

            if numDiffs > 1 || (numDiffs == 1 && smudgeUsed) {
                break
            }

            if numDiffs == 1 {
                smudgeUsed = true
            }
        }

        if (i + j == len(puzzle) || i - j == 0) && smudgeUsed {
            return i
        }
    }

    return -1
}

func VerticalReflection(puzzle []string, hasSmudge bool) int {
    for i := 1; i < len(puzzle[0]); i++ {
        j := 0
        smudgeUsed := !hasSmudge

        for ; i + j < len(puzzle[0]) && i - j - 1 >= 0; j++ {
            numDiffs := 0

            for k := 0; k < len(puzzle); k++ {
                if puzzle[k][i + j] != puzzle[k][i - j - 1] {
                    numDiffs++
                }
            }

            if numDiffs > 1 || (numDiffs == 1 && smudgeUsed) {
                break
            }

            if numDiffs == 1 {
                smudgeUsed = true
            }
        }

        if (i + j == len(puzzle[0]) || i - j == 0) && smudgeUsed {
            return i
        }
    }

    return -1
}

func SolvePuzzle(puzzle []string, hasSmudge bool) int {
    vertical := VerticalReflection(puzzle, hasSmudge)
    horizontal := HorizontalReflection(puzzle, hasSmudge)

    if vertical != -1 {
        return vertical
    }

    return 100 * horizontal
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    puzzle := ReadPuzzle(scanner)

    for len(puzzle) > 0 {
        sum += SolvePuzzle(puzzle, true)
        puzzle = ReadPuzzle(scanner)
    }

    fmt.Println(sum)
}

