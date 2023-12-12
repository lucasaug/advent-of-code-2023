package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
    FIVE_OF_A_KIND = 7
    FOUR_OF_A_KIND = 6
    FULL_HOUSE = 5
    THREE_OF_A_KIND = 4
    TWO_PAIR = 3
    ONE_PAIR = 2
    HIGH_CARD = 1
)

type Hand struct {
    hand string
    bid int
}

func ParseHand(s string) Hand {
    fields := strings.Fields(s)

    hand := fields[0]
    bid, _ := strconv.Atoi(fields[1])

    return Hand {
        hand: hand,
        bid: bid,
    }
}

func ClassifyHand(c Hand) int {
    letterCount := map[rune]int {}
    jCount := 0
    mostCommon := '0'

    for _, letter := range c.hand {
        if letter != 'J' {
            _, ok := letterCount[letter]
            if !ok {
                letterCount[letter] = 0
            }

            letterCount[letter]++

            _, mostCommonIsSet := letterCount[mostCommon]
            if !mostCommonIsSet || letterCount[mostCommon] < letterCount[letter] {
                mostCommon = letter
            }
        } else {
            jCount++
        }
    }

    letterCount[mostCommon] += jCount

    if len(letterCount) == 1 {
        return FIVE_OF_A_KIND
    } else if len(letterCount) == 2 {
        for _, value := range letterCount {
            if value == 1 || value == 4 {
                return FOUR_OF_A_KIND
            }
            return FULL_HOUSE
        }
    } else if len(letterCount) == 3 {
        for _, value := range letterCount {
            if value == 3 {
                return THREE_OF_A_KIND
            }
        }
        return TWO_PAIR
    } else if len(letterCount) == 4 {
        return ONE_PAIR
    }

    return HIGH_CARD
}

func BreakTie(a, b Hand) int {
    cards := []rune {
        'A',
        'K',
        'Q',
        'J',
        'T',
        '9',
        '8',
        '7',
        '6',
        '5',
        '4',
        '3',
        '2',
        'J',
    }

    for i, charA := range a.hand {
        // The input is ASCII encoded
        charB, _ := utf8.DecodeRuneInString(b.hand[i:])

        if charA != charB {
            firstIndex := -1
            secondIndex := -1

            for i, char := range cards {
                if char == charA {
                    firstIndex = i
                }
                if char == charB {
                    secondIndex = i
                }
            }

            return secondIndex - firstIndex
        }
    }

    return 0
}

func CompareHands(a, b Hand) int {
    rankA := ClassifyHand(a)
    rankB := ClassifyHand(b)

    if rankA != rankB {
        return rankA - rankB
    }

    return BreakTie(a, b)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var hands []Hand
    for scanner.Scan() {
        line := scanner.Text()
        hand := ParseHand(line)

        hands = append(hands, hand)
    }

    slices.SortFunc(hands, CompareHands)

    result := 0
    for i, hand := range hands {
        fmt.Println(hand.hand, ":", hand.bid, "*", (i + 1))
        result += hand.bid * (i + 1)
    }

    fmt.Println(result)
}
