package main

import (
	"strconv"
	"strings"
)

type rule struct {
	before int
	after int
}

func inputToRulesAndUpdates (input string) ([]rule, [][]int) {
	var parts = strings.SplitN(input, "\n\n", 2)

	var rulesStr = strings.Split(parts[0], "\n")
	var rules []rule
	for _, ruleStr := range rulesStr {
		// var rule rule
		ruleParts := strings.Split(ruleStr, "|")
		before, err := strconv.Atoi(ruleParts[0])
		if err != nil {
			panic(err)
		}
		after, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			panic(err)
		}

		rule := rule{before, after}

		rules = append(rules, rule)
	}

	var updatesStr = strings.Split(parts[1], "\n")
	var updates [][]int
	for _, updateStr := range updatesStr {
		var update []int
		updateParts := strings.Split(updateStr, ",")
		for _, updatePart := range updateParts {
			number, err := strconv.Atoi(updatePart)
			if err != nil {
				panic(err)
			}
			update = append(update, number)
		}

		updates = append(updates, update)
	}

	return rules, updates
}

func isUpdateValid(update []int, rules []rule) (bool, int) {
	for _, rule := range rules {
		var before int = -1
		var after int = -1
		for index, page := range update {
			if(page == rule.before) {
				before = index
			} else if page == rule.after {
				after = index
			}
			
			if(before != -1 && after != -1) {
				if(before > after) {
					return false, 0
				}
			}
		}		
	}

	return true, update[(len(update)-1)/2]
}

func part1(input string) int {
	var sum int = 0

	var rules, updates = inputToRulesAndUpdates(input)
	
	// fmt.Printf("Rules: %v\n", rules)
	// fmt.Printf("Updates: %v\n", updates)

	for _, update:= range updates {
		var valid, middle = isUpdateValid(update, rules)
		if valid {
			// fmt.Printf("Middle: %d\n", middle)
			sum += middle
		}
	}

	return sum
}