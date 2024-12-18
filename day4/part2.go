package main

func part2(input string) int {

	array := textToArray(input)
	var xmas = 0

	if(len(array) <= 0) {
		return xmas
	}

	var h = len(array)
	var w = len(array[0])

	for i:=0; i< h; i++ {
		for j:=0; j<w; j++ {
			if array[i][j] == 'A' {
				if(i>0 && j>0 && i<h-1 && j<w-1) {
					if array[i+1][j+1] == 'M' && array[i-1][j-1] == 'S' || array[i-1][j-1] == 'M' && array[i+1][j+1] == 'S' {
						if array[i-1][j+1] == 'M' && array[i+1][j-1] == 'S' || array[i+1][j-1] == 'M' && array[i-1][j+1] == 'S' {
							xmas += 1
						}	
					}
				}
			}
		}
	}
	return xmas
}