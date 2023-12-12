package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseNode(s string) (string, []string) {
    splitLine := strings.Split(s, " = ")

    nodeName := splitLine[0]
    directionsStr := splitLine[1]
    directionsStr = directionsStr[1:len(directionsStr)-1]

    directions := strings.Split(directionsStr, ", ")

    return nodeName, directions
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    instructions := scanner.Text()

    // ignore empty line
    scanner.Scan()
    scanner.Text()

    nodes := map[string][]string {}
    initialNodes := []string {}
    for scanner.Scan() {
        nodeName, nodeDirections := ParseNode(scanner.Text())
        nodes[nodeName] = nodeDirections

        if nodeName[len(nodeName) - 1] == byte('A') {
            initialNodes = append(initialNodes, nodeName)
        }
    }

    minStepCount := 1
    for _, node := range initialNodes {
        stepCount := 0
        currentNode := node

        for currentNode[len(currentNode)-1] != 'Z' {
            currentInstruction := instructions[stepCount % len(instructions)]

            index := 0
            if currentInstruction == byte('R') {
                index = 1
            }

            currentNode = nodes[currentNode][index]

            stepCount++
        }

        // for some reason using the LCM is the correct solution (even though
        // it's incorrect for the general case and the problem statement
        // doesn't tell you to assume this would work -.-)
        minStepCount = LCM(minStepCount, stepCount)
    }

    fmt.Println(minStepCount)
}
