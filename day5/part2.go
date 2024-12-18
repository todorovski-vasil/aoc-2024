package main

func swap[T any](items []T, j int, i int) []T {
	var newItems []T
	newItems = append(newItems, items[0:i]...)
	newItems = append(newItems, items[j])
	newItems = append(newItems, items[i+1:j]...)
	newItems = append(newItems, items[i])
	newItems = append(newItems, items[j+1:]...)
	return newItems
}

func isUpdateValidWithIndexes(update []int, rules []rule) (bool, int, int) {
	for _, rule := range rules {
		var before int = -1
		var after int = -1
		for index, page := range update {
			if page == rule.before {
				before = index
			} else if page == rule.after {
				after = index
			}

			if before != -1 && after != -1 {
				if before > after {
					return false, before, after
				}
			}
		}
	}

	return true, -1, -1
}

func fixUpdate(update []int, rules []rule) (bool, []int) {
	needsFixing := false
	var currentUpdate = make([]int, len(update))
	copy(currentUpdate, update)
	var valid, before, after = isUpdateValidWithIndexes(currentUpdate, rules)
	for {
		if valid {
			break
		}
		needsFixing = true

		currentUpdate = swap(currentUpdate, before, after)
		valid, before, after = isUpdateValidWithIndexes(currentUpdate, rules)
	}

	return needsFixing, currentUpdate
}

func part2(input string) int {
	var sum int = 0
	var rules, updates = inputToRulesAndUpdates(input)

	for _, update := range updates {
		var neededFixing, fixedUpdate = fixUpdate(update, rules)
		if neededFixing {
			sum += fixedUpdate[(len(fixedUpdate)-1)/2]
		}
	}

	return sum
}