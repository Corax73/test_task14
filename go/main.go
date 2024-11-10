package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{4, 8, 15, 15, 16, 23, 42}
	searchedInt := 15
	fmt.Println(binarySearch(slice, searchedInt))
}

// binarySearch searches the passed array for the key of the passed value. If input checks fail or there is no number in the array, returns -1.
func binarySearch(array []int, searchedInt int) int {
	resp := -1
	end := float64(len(array) - 1)
	if end+1 <= math.MaxInt {
		var start float64
		for start <= end {
			middle := math.Round((start + end) / 2)
			checkedValue := array[int(middle)]
			if searchedInt == checkedValue {
				resp = int(middle)
				if array[int(middle)-1] == searchedInt {
					resp = int(middle) - 1
				}
			} else if checkedValue > searchedInt {
				end = middle - 1
			} else {
				start = middle + 1
			}
			if resp > -1 {
				break
			}
		}
	}
	return resp
}
