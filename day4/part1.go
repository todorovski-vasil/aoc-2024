package main

import (
	"fmt"
	"strings"
)

func textToArray(text string) [][]byte {
	lines := strings.Split(text, "\n")

	// var h = len(lines)

	var array [][]byte
	// array := [][]byte{}
	for i, line := range lines {
		linija := []byte{}
		array = append(array, linija)
		for j := 0; j < len(line); j++ {
			// array[i][j] = line[j]
			array[i] = append(array[i], line[j])
		}
	}

	return array;
}

func part1(input string) int {

	array := textToArray(input)
	var xmas = 0

	if(len(array) <= 0) {
		return xmas
	}

	var h = len(array)
	var w = len(array[0])

	for i:=0; i< h; i++ {
		for j:=0; j<w; j++ {
			if array[i][j] == 'X' {
				if(j>2) {
					word := fmt.Sprintf("%c%c%c", array[i][j-1], array[i][j-2], array[i][j-3]) 
					if word == "MAS" {
						xmas += 1
					}
				}

				if(j<w-3) {
					word := fmt.Sprintf("%c%c%c", array[i][j+1], array[i][j+2], array[i][j+3]) 
					if word == "MAS" {
						xmas += 1
					}
				}

				if(i>2) {
					word := fmt.Sprintf("%c%c%c", array[i-1][j], array[i-2][j], array[i-3][j]) 
					if word == "MAS" {
						xmas += 1
					}

					if(j>2) {
						word := fmt.Sprintf("%c%c%c", array[i-1][j-1], array[i-2][j-2], array[i-3][j-3]) 
						if word == "MAS" {
							xmas += 1
						}
					}

					if(j<w-3) {
						word := fmt.Sprintf("%c%c%c", array[i-1][j+1], array[i-2][j+2], array[i-3][j+3]) 
						if word == "MAS" {
							xmas += 1
						}
					}
				}

				if(i<h-3) {
					word := fmt.Sprintf("%c%c%c", array[i+1][j], array[i+2][j], array[i+3][j]) 
					if word == "MAS" {
						xmas += 1
					}

					if(j>2) {
						word := fmt.Sprintf("%c%c%c", array[i+1][j-1], array[i+2][j-2], array[i+3][j-3]) 
						if word == "MAS" {
							xmas += 1
						}
					}

					if(j<w-3) {
						word := fmt.Sprintf("%c%c%c", array[i+1][j+1], array[i+2][j+2], array[i+3][j+3]) 
						if word == "MAS" {
							xmas += 1
						}
					}
				}
			}
		}
	}
	return xmas
}