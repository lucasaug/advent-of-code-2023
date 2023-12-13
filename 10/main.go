package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
    NORTH = iota
    EAST
    SOUTH
    WEST
    NONEDIR
)

func ParseMap(scanner *bufio.Scanner) ([][]string, int, int) {
    // the whole map is padded with ground tiles to avoid annoying corner cases
    result := [][]string {}
    startingRow := 0
    startingCol := 0

    row := 0
    for scanner.Scan() {
        line := scanner.Text()
        line = "." + line + "."

        result = append(result, strings.Split(line, ""))

        for col, char := range line {
            if char == 'S' {
                // account for padding
                startingRow = row + 1
                startingCol = col
            }
        }

        row++
    }

    emptyLine := strings.Repeat(".", len(result[0]))
    emptyRow := strings.Split(emptyLine,  "")

    result = append([][]string { emptyRow }, result...)
    result = append(result, [][]string { emptyRow }...)

    return result, startingRow, startingCol
}

func OppositeDirection(direction int) int {
    if direction == NORTH {
        return SOUTH
    }
    if direction == EAST {
        return WEST
    }
    if direction == SOUTH {
        return NORTH
    }
    if direction == WEST {
        return EAST
    }

    return NONEDIR
}

func DirectionToOffset(direction int) (int, int) {
    if direction == NORTH {
        return -1, 0
    }
    if direction == EAST {
        return 0, 1
    }
    if direction == SOUTH {
        return 1, 0
    }
    if direction == WEST {
        return 0, -1
    }

    return 0, 0
}

func GetAvailableDirections(puzzleMap [][]string, row, col int) [2]int {
    char := puzzleMap[row][col]

    if char == "|" {
        return [2]int { NORTH, SOUTH }
    }
    if char == "-" {
        return [2]int { EAST, WEST }
    }
    if char == "L" {
        return [2]int { NORTH, EAST }
    }
    if char == "J" {
        return [2]int { NORTH, WEST }
    }
    if char == "7" {
        return [2]int { SOUTH, WEST }
    }
    if char == "F" {
        return [2]int { SOUTH, EAST }
    }
    if char == "S" {
        result := [2]int{}

        resultCount := 0
        for i := 0; i < 4; i++ {
            offsetRow, offsetCol := DirectionToOffset(i)

            offsetDirections := GetAvailableDirections(puzzleMap,
                row + offsetRow, col + offsetCol)

            for _, dir := range offsetDirections {
                if dir == OppositeDirection(i) {
                    // this is fine because the problem input assures us that
                    // S will have at most 2 connecting pipes
                    result[resultCount] = i
                    resultCount++
                }
            }
        }

        return result
    }

    return [2]int { NONEDIR, NONEDIR }
}

func CalculateFarthest(
    puzzleMap [][]string,
    startingRow, startingCol int) int {

    directions := GetAvailableDirections(puzzleMap, startingRow, startingCol)

    currentlyVisiting := [2][2]int {
        { startingRow, startingCol },
        { startingRow, startingCol },
    }
    distance := 1
    lastVisited := [2]int { startingRow, startingCol }

    lastDirection := directions
    
    for i, dir := range directions {
        offsetRow, offsetCol := DirectionToOffset(dir)
        currentlyVisiting[i][0] += offsetRow
        currentlyVisiting[i][1] += offsetCol
    }

    for lastVisited != currentlyVisiting[1] &&
        currentlyVisiting[0] != currentlyVisiting[1] {

        lastVisited = currentlyVisiting[0]
        for i, cursor := range currentlyVisiting {
            row := cursor[0]
            col := cursor[1]
            directions := GetAvailableDirections(puzzleMap, row, col)

            selected := directions[0]
            if directions[0] == OppositeDirection(lastDirection[i]) {
                selected = directions[1]
            }

            offsetRow, offsetCol := DirectionToOffset(selected)
            currentlyVisiting[i][0] += offsetRow
            currentlyVisiting[i][1] += offsetCol

            lastDirection[i] = selected
        }

        distance++
    }

    if currentlyVisiting[0] == currentlyVisiting[1] {
        return distance
    }

    return distance - 1
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    puzzleMap, startingRow, startingCol := ParseMap(scanner)
    fmt.Println(CalculateFarthest(puzzleMap, startingRow, startingCol))
}
