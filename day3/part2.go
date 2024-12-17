package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part2(input string) int {
	cleanInput := ""
	uncleanInput := input

	for {
		split1 := strings.SplitN(uncleanInput, "don't()", 2)
		cleanInput += split1[0]

		if(len(split1) == 1) {
			break
		}

		split2 := strings.SplitN(split1[1], "do()", 2)
		if(len(split2) == 1) {
			break
		}

		uncleanInput = split2[1]
	}	

	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	correctExpressions := r.FindAllString(cleanInput, -1)

	var sum int = 0

	for _, expression := range correctExpressions {
		operands := strings.Split(expression[4:len(expression)-1], ",")
		op1, err := strconv.Atoi(operands[0])
		if err != nil {
			panic(err)
		}
		op2, err := strconv.Atoi(operands[1])
		if err != nil {
			panic(err)
		}

		sum += op1 * op2
	}

	return sum
}