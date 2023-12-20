package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    NORTH = iota
    WEST
    SOUTH
    EAST
)

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


type EnergizedTile struct {
    x, y int
    direction int
}

func (tile *EnergizedTile) String() string {
    return fmt.Sprintf("%d,%d,%d", tile.x, tile.y, tile.direction)
}

func (tile *EnergizedTile) PosString() string {
    return fmt.Sprintf("%d,%d", tile.x, tile.y)
}


func Interact(value rune, direction int) []int {
    switch value {
    case '/':
        if direction == EAST {
            return []int{NORTH}
        } else if direction == WEST {
            return []int{SOUTH}
        } else if direction == SOUTH {
            return []int{WEST}
        } else if direction == NORTH {
            return []int{EAST}
        }
    case '\\':
        if direction == EAST {
            return []int{SOUTH}
        } else if direction == WEST {
            return []int{NORTH}
        } else if direction == SOUTH{
            return []int{EAST}
        } else if direction == NORTH{
            return []int{WEST}
        }
    case '-':
        if direction == NORTH || direction == SOUTH {
            return []int{ WEST, EAST }
        }
    case '|':
        if direction == WEST || direction == EAST {
            return []int{ NORTH, SOUTH }
        }
    }

    // '.' rune or pointy end of a splitter
    return []int{ direction }
}

func Iterate(layout []string, current EnergizedTile) []EnergizedTile {
    currTile := layout[current.x][current.y]
    nextDirections := Interact(rune(currTile), current.direction)

    nextTiles := []EnergizedTile{}

    for _, direction := range nextDirections {
        xOffset, yOffset := DirectionToOffset(direction)
        x, y := current.x + xOffset, current.y + yOffset

        if x >= 0 && y >= 0 && x < len(layout[0]) && y < len(layout) {
            nextTiles = append(nextTiles, EnergizedTile{
                x, y, direction,
            })
        }

    }

    return nextTiles
}

func CountEnergized(layout []string) int {
    isEnergized := make([][]bool, len(layout))
    for i := range isEnergized {
        isEnergized[i] = make([]bool, len(layout[0]))
    }

    toEvaluate := []EnergizedTile{{ 0, 0, EAST }}
    evaluated := make(map[string]bool)

    for len(toEvaluate) > 0 {
        current := toEvaluate[0]
        toEvaluate = toEvaluate[1:]

        if !evaluated[current.String()] {
            evaluated[current.String()] = true
            isEnergized[current.x][current.y] = true

            var next []EnergizedTile
            next = Iterate(layout, current)
            toEvaluate = append(toEvaluate, next...)
        }
    }

    count := 0
    for _, row := range isEnergized {
        for _, value := range row {
            if value {
                count++
            }
        }
    }

    return count
}

func DPCountEnergized(
    layout []string,
    current EnergizedTile,
    evaluated *map[string]bool,
    memoization *map[string]int,
) int {
    if val, ok := (*memoization)[current.String()]; ok {
        return val
    }

    (*evaluated)[current.String()] = true

    currentChar := layout[current.x][current.y] 
    if currentChar == '-' &&
        (current.direction == NORTH || current.direction == SOUTH) {
        oppositeDir := current
        oppositeDir.direction = (oppositeDir.direction + 2) % 4
        (*evaluated)[oppositeDir.String()] = true
    }
    if currentChar == '|' &&
        (current.direction == WEST || current.direction == EAST) {
        oppositeDir := current
        oppositeDir.direction = (oppositeDir.direction + 2) % 4
        (*evaluated)[oppositeDir.String()] = true
    }

    var next []EnergizedTile
    next = Iterate(layout, current)

    maxEnergized := 0
    for _, tile := range next {
        if !(*evaluated)[tile.String()] {
            if _, ok := (*memoization)[tile.String()]; !ok {
                DPCountEnergized(layout, tile, evaluated, memoization)
            }

            maxEnergized += (*memoization)[tile.String()]
        }
    }

    // This next part is an ugly hack so that a tile is not counted as
    // energized twice if two light beams in different directions pass through
    // it
    if !(*evaluated)[current.PosString()] {
        maxEnergized++
    }
    (*evaluated)[current.PosString()] = true

    (*memoization)[current.String()] = maxEnergized
    return maxEnergized
}

func CalculateAllPossiblePaths(layout []string) int {
    initialTiles := []EnergizedTile{}
    for i := 0; i < len(layout); i++ {
        initialTiles = append(initialTiles, EnergizedTile{ i, 0, EAST })
        initialTiles = append(initialTiles, EnergizedTile{
            i, len(layout[0]) - 1, WEST,
        })
    }
    for i := 0; i < len(layout[0]); i++ {
        initialTiles = append(initialTiles, EnergizedTile{ 0, i, SOUTH })
        initialTiles = append(initialTiles, EnergizedTile{
            len(layout) - 1, i, NORTH,
        })
    }

    result := 0
    for _, start := range initialTiles {
        evaluated := make(map[string]bool)
        memoization := make(map[string]int)
        result = max(DPCountEnergized(
            layout, start, &evaluated, &memoization,
        ), result)
    }

    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var layout []string
    for scanner.Scan() {
        line := scanner.Text()
        layout = append(layout, line)
    }

    fmt.Println(CalculateAllPossiblePaths(layout))
}
