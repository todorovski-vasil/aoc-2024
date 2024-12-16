package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) int {
	var safeConfigs int = 0
	for _, report := range strings.Split(input, "\n") {
		var previousLevel int = 0
		var reportDirection string
		var reportSafe bool = true
		for index, levelString := range strings.Split(report, " ") {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				panic(err)
			}

			switch index {
			case 0:
				previousLevel = level
			case 1:
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
			default:
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

		if reportSafe {
			safeConfigs++
		}
	}
	return safeConfigs
}
