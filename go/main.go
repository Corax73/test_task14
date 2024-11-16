package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Solutions struct{}

func main() {
	solution := Solutions{}
	solution.GetFiborow()
}

func (solution *Solutions) GetBinarySearch() {
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

func (solution *Solutions) GetFindWeekend() {
	fmt.Println(solution.findWeekend("05.05.2024", "25.05.2024"))
}

// findWeekend if the arguments are correct, returns the number of days off in the passed time period, otherwise returns -1.
func (solutions *Solutions) findWeekend(begin, end string) int {
	resp := -1
	_, err := regexp.Compile(`06.06.2020`)
	if err != nil {
		fmt.Println(err)
	}
	matchedBegin, _ := regexp.MatchString(`\d+.\d+.\d+`, begin)
	matchedEnd, _ := regexp.MatchString(`\d+.\d+.\d+`, end)
	if matchedBegin && matchedEnd {
		begin = solutions.formatDateString(begin)
		timeBegin, err := time.Parse(time.DateOnly, begin)
		if err != nil {
			fmt.Println(err)
		} else {
			end = solutions.formatDateString(end)
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

// formatDateString formats the passed time string to the time.DateOnly format.
func (solutions *Solutions) formatDateString(str string) string {
	strArr := strings.Split(str, ".")
	var strBuilder strings.Builder
	for i := 2; i >= 0; i-- {
		strBuilder.WriteString(strArr[i])
		strBuilder.WriteString("-")
	}
	resp := strings.Trim(strBuilder.String(), "-")
	strBuilder.Reset()
	return resp
}

func (solution *Solutions) GetRgbToInt() {
	fmt.Println(solution.rgbToInt(255, 0, 255))
}

// rgbToInt converts a set of rgb color values ​​to an integer.
func (solutions *Solutions) rgbToInt(red, green, blue int64) int64 {
	var resp int64
	resp = -1
	if red >= 0 && red <= 255 && green >= 0 && green <= 255 && blue >= 0 && blue <= 255 {
		arr := []int64{red, green, blue}
		var strBuilder strings.Builder
		for _, color := range arr {
			temp := strconv.FormatInt(color, 16)
			if len(temp) < 2 {
				strBuilder.WriteString("0")
			}
			strBuilder.WriteString(temp)
		}
		str := strBuilder.String()
		strBuilder.Reset()
		temp, err := strconv.ParseInt(str, 16, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			resp = temp
		}
	}
	return resp
}

func (solution *Solutions) GetFiborow() {
	fmt.Println(solution.fiborow(math.MaxInt64))
}

// fiborow if the argument is valid, returns the Fibonacci sequence string up to the value of the argument, otherwise -1.
func (solutions *Solutions) fiborow(limit int64) string {
	resp := "-1"
	if limit > 0 && limit <= math.MaxInt64 {
		var fibArr []int64
		fibArr = append(fibArr, int64(0))
		fibArr = append(fibArr, int64(1))
		for i := 2; fibArr[i - 1] + fibArr[i - 2] < limit; i++ {
			newEl := fibArr[i - 1] + fibArr[i - 2]
			fibArr = append(fibArr, newEl)
			if math.MaxInt64 - fibArr[i - 1] <= newEl {
				break
			}
		}
		var strBuilder strings.Builder
		for _, val := range fibArr {
			strBuilder.WriteString(strconv.FormatInt(val, 10))
			strBuilder.WriteString(" ")
		}
		resp = strings.Trim(strBuilder.String(), " ")
		strBuilder.Reset()
	}
	return resp
}
