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

func (r SpringRecord) String() string {
    var countList []string
    for _, value := range r.countList {
        converted := strconv.Itoa(value)
        countList = append(countList, converted)
    }

    return r.condition + strings.Join(countList, ",")
}

func ParseSpringRecord(line string, unfold bool) SpringRecord {
    splitRecord := strings.Fields(line)
    condition := splitRecord[0]
    countListStr := strings.Split(splitRecord[1], ",")

    countList := []int {}
    for _, value := range countListStr {
        converted, _ := strconv.Atoi(value)
        countList = append(countList, converted)
    }

    if unfold {
        condition = condition + strings.Repeat("?" + condition, 4)

        countListLen := len(countList)
        for i := 0; i < 4; i++ {
            countList = append(countList, countList[:countListLen]...)
        }
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

func BruteForce(record SpringRecord) int {
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

func DPRecurse(record SpringRecord, memoization map[string]int) int {
    recordString := record.String()
    if val, ok := memoization[recordString]; ok {
        return val
    }

    if len(record.condition) == 0 {
        if len(record.countList) > 0 {
            memoization[recordString] = 0
            return 0
        }

        memoization[recordString] = 1
        return 1
    }

    if len(record.countList) == 0 {
        for _, value := range record.condition {
            if value == '#' {
                memoization[recordString] = 0
                return 0
            }
        }

        memoization[recordString] = 1
        return 1
    }

    current := record.condition[0] 

    if current == '.' {
        subRecord := SpringRecord {
            condition: record.condition[1:],
            countList: record.countList,
        }

        memoization[recordString] = DPRecurse(subRecord, memoization)
        return memoization[recordString]
    }

    arrangements := 0
    if current == '?' {
        subRecord := SpringRecord {
            condition: record.condition[1:],
            countList: record.countList,
        }

        arrangements = DPRecurse(subRecord, memoization)
    }

    index := 1
    for ;
        index < record.countList[0] &&
        index < len(record.condition) &&
        record.condition[index] != '.';
        index++ {
    }

    if index == len(record.condition) {
        // end of a row
        if index == record.countList[0] && len(record.countList) == 1 {
            // feasible state
            memoization[recordString] = arrangements + 1
            return memoization[recordString]
        }
        // row is over but requirements haven't been met
        memoization[recordString] = arrangements
        return memoization[recordString]
    }

    if index < record.countList[0] {
        // couldn't get all the required '#' entries
        memoization[recordString] = arrangements
        return memoization[recordString]
    }

    nextChar := record.condition[index]
    if nextChar == '#' {
        // too many '#' entries
        memoization[recordString] = arrangements
        return memoization[recordString]
    }

    // we got the required number of '#' entries, now make sure the next
    // position becomes a '.', remove one of the damaged counts and recurse
    subRecord := SpringRecord {
        condition: record.condition[index + 1:],
        countList: record.countList[1:],
    }

    memoization[recordString] = arrangements + DPRecurse(subRecord, memoization)
    return memoization[recordString]
}

func DPSolve(record SpringRecord) int {
    memoization := map[string]int {}
    return DPRecurse(record, memoization)
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        spring := ParseSpringRecord(line, true)
        sum += DPSolve(spring)
    }

    fmt.Println(sum)
}

