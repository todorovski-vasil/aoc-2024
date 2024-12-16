package main

import (
	"math"
	"strconv"
	"strings"
)

func checkReport(levels []string) (int, bool) {
	var previousLevel int = 0
	var reportDirection string
	var reportSafe bool = true
	var index = 0

	for ; index < len(levels); index++ {
		level, err := strconv.Atoi(levels[index])
		if err != nil {
			panic(err)
		}

		if index == 1 {
			var diff = int(math.Abs(float64(previousLevel - level)))
			if diff < 1 || diff > 3 {
				reportSafe = false
				break
			}
			if level > previousLevel {
				reportDirection = "up"
			} else {
				reportDirection = "down"
			}
		} else if index > 1 {
			var diff = int(math.Abs(float64(previousLevel - level)))
			if diff < 1 || diff > 3 {
				reportSafe = false
				break
			}
			if level > previousLevel {
				if reportDirection == "down" {
					reportSafe = false
					break
				}
			} else {
				if reportDirection == "up" {
					reportSafe = false
					break
				}
			}
		}
		previousLevel = level
	}

	return index, reportSafe
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func part2(input string) int {
	var safeConfigs int = 0
	for _, report := range strings.Split(input, "\n") {
		var levels = strings.Split(report, " ")
		var index, reportSafe = checkReport(levels)

		if reportSafe {
			safeConfigs++
		} else {
			levels = remove(strings.Split(report, " "), index)
			var _, reportSafe = checkReport(levels)
			if reportSafe {
				safeConfigs++
				continue
			} else if index > 0 {
				levels = remove(strings.Split(report, " "), index-1)
				var _, reportSafe = checkReport(levels)
				if reportSafe {
					safeConfigs++
					continue
				} else if index > 1 {
					levels = remove(strings.Split(report, " "), index-2)
					var _, reportSafe = checkReport(levels)
					if reportSafe {
						// fmt.Printf("%s, %d\n", report, index)
						// fmt.Printf("%v\n", remove(strings.Split(report, " "), index))
						// fmt.Printf("%v\n", remove(strings.Split(report, " "), index-1))
						// fmt.Printf("%v\n", remove(strings.Split(report, " "), index-2))

						safeConfigs++
						continue
					}
				}
			}
		}
	}
	return safeConfigs
}
