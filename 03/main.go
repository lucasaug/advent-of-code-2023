package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Solver struct {
    input []string
    currentRow uint64
    currentCol uint64
}

func (s Solver) NextNumber() int {
    for i, _ := range s.input[1:len(s.input)-1] {
	fmt.Println(i)
    }

    return -1
}

func (s Solver) Solve() int {
    s.currentRow = 0
    s.currentCol= 0

    sum := 0

    for next := 0; next != -1; next = s.NextNumber() {
	sum += next
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

