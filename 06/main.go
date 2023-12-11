package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateDistance(maximumTime int, pressedDown int) int {
    return pressedDown * (maximumTime - pressedDown)
}


type Race struct {
    time int
    distance int
}

func (r *Race) NumberOfWinningScenarios() int {
    count := 0

    for i := 1; i < r.time; i++ {
        distance := CalculateDistance(r.time, i)
        if distance > r.distance {
            count++
        }
    }

    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    timeLine := scanner.Text()
    scanner.Scan()
    distanceLine := scanner.Text()

    // Skip headers
    timeLine = timeLine[5:]
    distanceLine = distanceLine[9:]

    /* First part
    raceTimes := strings.Fields(timeLine)
    raceDistances := strings.Fields(distanceLine)

    var races []Race
    for i, time := range raceTimes {
        timeInt ,_ := strconv.Atoi(time)
        distanceInt, _ := strconv.Atoi(raceDistances[i])

        races = append(races, Race{ timeInt, distanceInt })
    }

    result := 1
    for _, race := range races {
        result *= race.NumberOfWinningScenarios()
    }
    */

    raceTimeStr := strings.ReplaceAll(timeLine, " ", "")
    raceDistanceStr := strings.ReplaceAll(distanceLine, " ", "")

    raceTime, _ := strconv.Atoi(raceTimeStr)
    raceDistance, _ := strconv.Atoi(raceDistanceStr)

    race := Race {
        time: raceTime,
        distance: raceDistance,
    }

    // I expected needing to apply an optimization in the large scenario but
    // apparently not o.o
    fmt.Println(race.NumberOfWinningScenarios())
}
