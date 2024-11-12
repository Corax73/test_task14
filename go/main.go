package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"time"
)

type Solutions struct{}

func main() {
	solution := Solutions{}
	fmt.Println(solution.findWeekend( /*"05.05.2024", "25.05.2024"*/ "2024-05-05", "2024-05-25"))
}

func (solution *Solutions) getBinarySearch() {
	slice := []int{4, 8, 15, 15, 16, 23, 42}
	searchedInt := 15
	fmt.Println(solution.binarySearch(slice, searchedInt))
}

// binarySearch searches the passed array for the key of the passed value. If input checks fail or there is no number in the array, returns -1.
func (solutions *Solutions) binarySearch(array []int, searchedInt int) int {
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

func (solutions *Solutions) findWeekend(begin, end string) int {
	resp := -1
	_, err := regexp.Compile(`06.06.2020`)
	if err != nil {
		fmt.Println(err)
	}
	matchedBegin, _ := regexp.MatchString(`\d+.\d+.\d+`, begin)
	matchedEnd, _ := regexp.MatchString(`\d+.\d+.\d+`, end)
	if matchedBegin && matchedEnd {
		timeBegin, err := time.Parse(time.DateOnly, begin)
		if err != nil {
			fmt.Println(err)
		} else {
			timeEnd, err := time.Parse(time.DateOnly, end)
			if err != nil {
				fmt.Println(err)
			} else {
				if timeBegin.Unix() <= math.MaxInt && timeEnd.Unix() <= math.MaxInt {
					resp = 0
					weekendsNumbers := []int{5, 6}
					for timeBegin.Unix() <= timeEnd.Unix() {
						if slices.Contains(weekendsNumbers, int(timeBegin.Weekday())) {
							resp++
							timeBegin = timeBegin.Add(24 * time.Hour)
						} else {
							hours := 24 * (5 - int(timeBegin.Weekday()))
							timeBegin = timeBegin.Add(time.Duration(hours) * time.Hour)
						}
					}
				}
			}
		}
	}
	return resp
}
