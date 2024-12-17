package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {

	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)

	correctExpressions := r.FindAllString(input, -1)
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