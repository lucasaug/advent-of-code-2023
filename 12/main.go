package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type SpringRecord struct {
    condition string
    countList []int
}

func ParseSpringRecord(line string) SpringRecord {
    splitRecord := strings.Fields(line)
    condition := splitRecord[0]
    countListStr := strings.Split(splitRecord[1], ",")

    countList := []int {}
    for _, value := range countListStr {
        converted, _ := strconv.Atoi(value)
        countList = append(countList, converted)
    }

    return SpringRecord{
        condition,
        countList,
    }
}

func ArrangementFromInt(condition string, arrangementIndex int) string {
    currentBit := 0
    result := ""

    for i := 0; i < len(condition); i++ {
        character := condition[i]

        if character == '?' {
            character = '.'
            if (arrangementIndex >> currentBit) & 1 == 1 {
                character = '#'
            }
            currentBit++
        }

        result += string(character)
    }

    return result
}

func CountListFromArrangement(arrangement string) []int {
    currentDamagedCount := 0

    result := []int{}
    for _, value := range arrangement {
        if value == '.' {
            if currentDamagedCount > 0 {
                result = append(result, currentDamagedCount)
            }
            currentDamagedCount = 0
        } else {
            currentDamagedCount++
        }
    }

    if currentDamagedCount > 0 {
        result = append(result, currentDamagedCount)
    }

    return result
}

func CalculateArrangements(record SpringRecord) int {
    unknowns := strings.Count(record.condition, "?")
    numArrangements := int(math.Pow(2.0, float64(unknowns)))

    count := 0
    for i := 0; i < numArrangements; i++ {
        arrangement := ArrangementFromInt(record.condition, i)
        countList := CountListFromArrangement(arrangement)
        if reflect.DeepEqual(countList, record.countList) {
            count++
        }
    }

    return count
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        spring := ParseSpringRecord(line)
        sum += CalculateArrangements(spring)
    }

    fmt.Println(sum)
}

