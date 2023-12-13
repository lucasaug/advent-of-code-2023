package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func AllEqual(values []int) bool {
    for _, v := range values {
        if v != values[0] {
            return false
        }
    }

    return true
}

func PredictNext(values []int) int {
    finalValues := []int { values[len(values)-1] }

    for !AllEqual(values) {
        newValues := []int {}
        for i := range values[:len(values)-1] {
            newValues = append(newValues, values[i+1] - values[i])
        }
        values = newValues

        finalValues = append(finalValues, values[len(values)-1])
    }

    result := 0
    for _, value := range finalValues {
        result += value
    }

    return result
}

func PredictPrevious(values []int) int {
    slices.Reverse(values)
    return PredictNext(values)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        values := strings.Fields(line)

        intValues := []int {}
        for _, v := range values {
            value, _ := strconv.Atoi(v)
            intValues = append(intValues, value)
        }

        sum += PredictPrevious(intValues)
    }

    fmt.Println(sum)
}
