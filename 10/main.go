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

func TraverseAndReplace(
    puzzleMap [][]string,
    startingRow, startingCol int) [][]string {
    directions := GetAvailableDirections(puzzleMap, startingRow, startingCol)
    puzzleMap[startingRow][startingCol] = "X"

    offsetRow, offsetCol := DirectionToOffset(directions[0])
    currentlyVisiting := [2]int {
        startingRow + offsetRow,
        startingCol + offsetCol,
    }

    lastDirection := directions[0]

    for currentlyVisiting != [2]int {startingRow, startingCol} {
        row := currentlyVisiting[0]
        col := currentlyVisiting[1]
        directions := GetAvailableDirections(puzzleMap, row, col)

        selected := directions[0]
        if directions[0] == OppositeDirection(lastDirection) {
            selected = directions[1]
        }

        puzzleMap[currentlyVisiting[0]][currentlyVisiting[1]] = "X"
        offsetRow, offsetCol := DirectionToOffset(selected)
        currentlyVisiting[0] += offsetRow
        currentlyVisiting[1] += offsetCol

        lastDirection = selected
    }

    return puzzleMap
}

func CountEnclosedTiles(
    puzzleMap [][]string,
    startingRow, startingCol int) int {
    var markedPuzzleMap [][]string
    for _, row := range puzzleMap {
        newRow := append([]string{}, row...)
        markedPuzzleMap = append(markedPuzzleMap, newRow)
    }

    markedPuzzleMap = TraverseAndReplace(markedPuzzleMap, startingRow, startingCol)

    for i, row := range markedPuzzleMap {
        inside := false
        inwall := false
        wallIncomingDir := NONEDIR

        for j, value := range row {
            if value == "X" {
                if !inwall {
                    dirs := GetAvailableDirections(puzzleMap, i, j)
                    enteringWall := false
                    orthogonalDir := NONEDIR

                    for _, dir := range dirs {
                        if dir == EAST {
                            enteringWall = true
                        } else {
                            orthogonalDir = dir
                        }
                    }

                    if enteringWall {
                        wallIncomingDir = orthogonalDir
                        inwall = true
                    } else {
                        inside = !inside
                    }
                } else {
                    dirs := GetAvailableDirections(puzzleMap, i, j)
                    orthogonalDir := NONEDIR

                    for _, dir := range dirs {
                        if dir == NORTH || dir == SOUTH {
                            orthogonalDir = dir
                        }
                    }

                    if orthogonalDir != NONEDIR {
                        // leaving wall
                        inwall = false
                        if orthogonalDir != wallIncomingDir {
                            inside = !inside
                        }
                    }
                }
            } else if inside {
                markedPuzzleMap[i][j] = "O"
            }
        }
    }

    for i := 0; i < len(markedPuzzleMap[0]); i++ {
        inside := false
        inwall := false
        wallIncomingDir := NONEDIR

        for j := range markedPuzzleMap {
            if markedPuzzleMap[j][i] == "X" {
                if !inwall {
                    dirs := GetAvailableDirections(puzzleMap, j, i)
                    enteringWall := false
                    orthogonalDir := NONEDIR

                    for _, dir := range dirs {
                        if dir == SOUTH {
                            enteringWall = true
                        } else {
                            orthogonalDir = dir
                        }
                    }

                    if enteringWall {
                        wallIncomingDir = orthogonalDir
                        inwall = true
                    } else {
                        inside = !inside
                    }
                } else {
                    dirs := GetAvailableDirections(puzzleMap, j, i)
                    orthogonalDir := NONEDIR

                    for _, dir := range dirs {
                        if dir == WEST || dir == EAST {
                            orthogonalDir = dir
                        }
                    }

                    if orthogonalDir != NONEDIR {
                        // leaving wall
                        if orthogonalDir != wallIncomingDir {
                            inside = !inside
                        }
                        inwall = false
                    }
                }
            } else if inside {
                markedPuzzleMap[j][i] = "O"
            }
        }
    }

    count := 0
    for _, row := range markedPuzzleMap {
        for _, value := range row {
            if value == "O" {
                count++
            }
        }
    }

    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    puzzleMap, startingRow, startingCol := ParseMap(scanner)
    // fmt.Println(CalculateFarthest(puzzleMap, startingRow, startingCol))
    fmt.Println(CountEnclosedTiles(puzzleMap, startingRow, startingCol))
}
