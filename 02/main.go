package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct  {
    red uint64
    green uint64
    blue uint64
}

type Game struct {
    id uint64
    CubeSet
}

var cubeAmounts = CubeSet {
    red: 12,
    green: 13,
    blue: 14,
}

func ParseGameId(s string) uint64 {
    var gameId uint64
    fmt.Sscanf(s, "Game %d", &gameId)

    return gameId
}

func ReadCubeSet(s string) CubeSet {
    var set CubeSet

    for _, cubeAmount := range strings.Split(s, ",") {
	splitAmount := strings.Split(strings.TrimSpace(cubeAmount), " ")

	switch splitAmount[1] {
	case "red":
	    set.red, _ = strconv.ParseUint(splitAmount[0], 10, 64)
	case "green":
	    set.green, _ = strconv.ParseUint(splitAmount[0], 10, 64)
	case "blue":
	    set.blue, _ = strconv.ParseUint(splitAmount[0], 10, 64)
	}
    }

    return set
}

func ReadGame(s string) Game {
    splitLine := strings.Split(s, ":")
    gameId := ParseGameId(splitLine[0])

    resultCubeSet := CubeSet { 0, 0, 0 }
    for _, cubeSetString := range strings.Split(splitLine[1], ";") {
	cubeSet := ReadCubeSet(cubeSetString)

	resultCubeSet.red = max(resultCubeSet.red, cubeSet.red)
	resultCubeSet.green = max(resultCubeSet.green, cubeSet.green)
	resultCubeSet.blue = max(resultCubeSet.blue, cubeSet.blue)
    }

    return Game { gameId, resultCubeSet }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var sum uint64 = 0
    for scanner.Scan() {
	line := scanner.Text()
	game := ReadGame(line)

	power := game.red * game.green * game.blue
	sum += power

	//if game.CubeSet.red <= cubeAmounts.red &&
	//   game.CubeSet.green <= cubeAmounts.green &&
        //   game.CubeSet.blue <= cubeAmounts.blue {
	//    sum += game.id
	// }
    }

    fmt.Println(sum)

}
