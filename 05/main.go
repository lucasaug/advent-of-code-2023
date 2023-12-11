package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type GardeningFunctionSection struct {
    destinationStart int
    sourceStart int
    length int
}

func (s GardeningFunctionSection) IsInputInRange(input int) bool {
    return s.sourceStart <= input && input < s.sourceStart + s.length
}

func (s GardeningFunctionSection) Apply(input int) int {
    return (input - s.sourceStart) + s.destinationStart
}

type GardeningFunction struct {
    functionSections []GardeningFunctionSection
    nextFunction *GardeningFunction
}

func (f GardeningFunction) Apply(input int) int {
    for _, section := range f.functionSections {
        if section.IsInputInRange(input) {
            return section.Apply(input)
        }
    }

    return input
}

func (f GardeningFunction) ApplyAll(input int) int {
    currentReturn := f.Apply(input)

    if f.nextFunction == nil {
        return currentReturn
    }

    return f.nextFunction.ApplyAll(currentReturn)
}

func ParseGardeningFunction(scanner *bufio.Scanner) GardeningFunction {
    result := GardeningFunction {}

    // Skip header line
    scanner.Text()

    for scanner.Scan() {
        line := scanner.Text()

        if len(line) == 0 {
            return result
        }

        functionParametersStr := strings.Fields(line)
        var functionParameters []int
        for _, param := range functionParametersStr {
            integerParam, _ := strconv.Atoi(param)
            functionParameters = append(functionParameters, integerParam)
        }

        functionSection := GardeningFunctionSection {
            destinationStart: functionParameters[0],
            sourceStart: functionParameters[1],
            length: functionParameters[2],
        }
        result.functionSections = append(
            result.functionSections,
            functionSection,
        )
    }

    return result
}

type SeedSpec struct {
    start int
    length int
}

func ParseSeeds(scanner *bufio.Scanner) []SeedSpec {
    scanner.Scan()
    line := scanner.Text()

    // Skip "seeds: "
    seedStr := line[7:]
    seedsParams := strings.Fields(seedStr)

    var seeds []SeedSpec
    i := 0
    for i < len(seedsParams) {
        start, _ := strconv.Atoi(seedsParams[i])
        length, _ := strconv.Atoi(seedsParams[i + 1])

        seeds = append(seeds, SeedSpec{
            start: start,
            length: length,
        })

        i += 2
    }

    return seeds
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    seeds := ParseSeeds(scanner)

    startFunction := GardeningFunction {}
    currentFunction := &startFunction

    // Skip empty line
    scanner.Scan()
    scanner.Text()

    for scanner.Scan() {
        function := ParseGardeningFunction(scanner)
        currentFunction.nextFunction = &function
        currentFunction = &function
    }

    var minimumResult = math.MaxInt64
    // This is painfully slow but I won't bother with something more
    // sophisticated
    for _, seedSpec := range seeds {
        for i := 0; i < seedSpec.length; i++ {
            minimumResult = min(minimumResult, startFunction.ApplyAll(
                seedSpec.start + i,
            ))
        }
    }

    fmt.Println(minimumResult)
}
