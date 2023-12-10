package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Solver struct {
    input []string
    currentRow int
    currentCol int
    gears map[string][]int
}

func IsSymbol(c byte) bool {
    return 0x2E != c &&
	((0x20 <= c && c <= 0x40 && !(0x30 <= c && c <= 0x39)) ||
	 (0x5B <= c && c <= 0x60) ||
	 (0x7B <= c))
}

func IsNumber(c byte) bool {
    return 0x30 <= c && c <= 0x39
}

func (s *Solver) ParseNumber() (bool, int) {
    isValid := IsSymbol(s.input[s.currentRow - 1][s.currentCol - 1]) ||
               IsSymbol(s.input[s.currentRow][s.currentCol - 1]) ||
               IsSymbol(s.input[s.currentRow + 1][s.currentCol - 1])

    surroundingGears := [][]int {}
    if s.input[s.currentRow - 1][s.currentCol - 1] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow - 1, s.currentCol - 1})
    }
    if s.input[s.currentRow][s.currentCol - 1] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow, s.currentCol - 1})
    }
    if s.input[s.currentRow + 1][s.currentCol - 1] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow + 1, s.currentCol - 1})
    }

    currentByte := s.input[s.currentRow][s.currentCol]
    value := 0
    for IsNumber(currentByte) {
	value = value * 10 + int(currentByte - 0x30)

	if IsSymbol(s.input[s.currentRow - 1][s.currentCol]) ||
	   IsSymbol(s.input[s.currentRow + 1][s.currentCol]) {
	    isValid = true
	}
	if s.input[s.currentRow - 1][s.currentCol] == 0x2A {
	    surroundingGears = append(surroundingGears,
		[]int{s.currentRow - 1, s.currentCol})
	}
	if s.input[s.currentRow + 1][s.currentCol] == 0x2A {
	    surroundingGears = append(surroundingGears,
		[]int{s.currentRow + 1, s.currentCol})
	}

	s.currentCol++
	currentByte = s.input[s.currentRow][s.currentCol]
    }

    if IsSymbol(s.input[s.currentRow - 1][s.currentCol]) ||
       IsSymbol(s.input[s.currentRow][s.currentCol]) ||
       IsSymbol(s.input[s.currentRow + 1][s.currentCol]) {
	isValid = true
    }

    if s.input[s.currentRow - 1][s.currentCol] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow - 1, s.currentCol})
    }
    if s.input[s.currentRow][s.currentCol] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow, s.currentCol})
    }
    if s.input[s.currentRow + 1][s.currentCol] == 0x2A {
	surroundingGears = append(surroundingGears,
	    []int{s.currentRow + 1, s.currentCol})
    }

    for _, gear := range surroundingGears {
	index := fmt.Sprintf("%d, %d", gear[0], gear[1])
	s.gears[index] = append(s.gears[index], value)
    }

    return isValid, value
}

func (s *Solver) NextNumber() int {
    for s.currentRow < len(s.input) - 1 {
	for s.currentCol < len(s.input[0]) - 1 {
	    if IsNumber(s.input[s.currentRow][s.currentCol]) {
		isValid, value := s.ParseNumber()
		if isValid {
		    return value
		}
	    }

	    s.currentCol++
	}

	s.currentRow++
	s.currentCol = 1
    }

    return -1
}

func (s *Solver) Solve() int {
    s.currentRow = 1
    s.currentCol = 1
    s.gears = make(map[string][]int)

    for next := 0; next != -1; next = s.NextNumber() {}

    sum := 0
    for _, gear := range s.gears {
	if len(gear) == 2 {
	    sum += gear[0] * gear[1]
	}
    }

    return sum
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    lines := []string {}

    for scanner.Scan() {
        line := scanner.Text()
        line = "." + line + "."

	lines = append(lines, line)
    }

    emptyLine := strings.Repeat(".", len(lines[0]))
    lines = append([]string{emptyLine}, lines...)
    lines = append(lines, emptyLine)

    solver := Solver {
	input: lines,
	currentRow: 0,
	currentCol: 0,
    }
    sum := solver.Solve()
    fmt.Println(sum)
}
