package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
    label string
    focalLen int
}

type Operation struct {
    label string
    operation rune
    focalLen int
}

func CalculateHash(entry string) byte {
    result := byte(0)

    for _, char := range entry {
        result += byte(char)
        result = result * 17
    }

    return result
}

func ParseOperation(entry string) Operation {
    if strings.Contains(entry, "=") {
	split := strings.Split(entry, "=")
	focalLen, _ := strconv.Atoi(split[1])

	return Operation{
	    split[0],
	    '=',
	    focalLen,
	}
    }

    return Operation{
	entry[:len(entry) - 1],
	'-',
	0,
    }
}

func FilterOperations(ops []Operation) []Operation {
    result := []Operation{}
    removed := make([]bool, len(ops))

    for i := len(ops) - 1; i >= 0; i-- {
	if !removed[i] {
	    if ops[i].operation == '=' {
		result = append(result, ops[i])
	    } else {
		for j := i-1; j >= 0; j-- {
		    if ops[j].label == ops[i].label &&
		       ops[j].operation == '=' && !removed[j] {
			removed[j] = true
			break
		    }
		}
	    }
	}
    }

    slices.Reverse(result)
    return result
}

func InsertOperation(boxes [256][]Lens, op Operation) [256][]Lens {
    boxIndex := CalculateHash(op.label)

    lensIndex := -1
    for i, lens := range boxes[boxIndex] {
	if lens.label == op.label {
	    lensIndex = i
	    break
	}
    }

    if lensIndex != -1 {
	boxes[boxIndex][lensIndex].focalLen = op.focalLen
    } else {
	boxes[boxIndex] = append(boxes[boxIndex], Lens{
	    op.label,
	    op.focalLen,
	})
    }

    return boxes
}

func RemoveOperation(boxes [256][]Lens, op Operation) [256][]Lens {
    boxIndex := CalculateHash(op.label)

    for i, lens := range boxes[boxIndex] {
	if lens.label == op.label {
	    boxes[boxIndex] = append(
		boxes[boxIndex][:i],
		boxes[boxIndex][i+1:]...
	    )

	    return boxes
	}
    }

    return boxes
}

func ApplyOperation(boxes [256][]Lens, op Operation) [256][]Lens {
    if op.operation == '=' {
	return InsertOperation(boxes, op)
    }

    return RemoveOperation(boxes, op)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    line := scanner.Text()

    fields := strings.Split(line, ",")

    var operations []Operation
    for _, field := range fields {
	operations = append(operations, ParseOperation(field))
    }

    boxes := [256][]Lens{}
    for i := range boxes {
	boxes[i] = make([]Lens, 0)
    }

    for _, op := range operations {
	boxes = ApplyOperation(boxes, op)
    }

    power := 0
    for i, box := range boxes {
	for j, lens := range box {
	    power += (i + 1) * (j + 1) * lens.focalLen
	}
    }

    fmt.Println(power)
}

