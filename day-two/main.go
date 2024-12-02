package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertDataToArr(file_name string) [1000][]int {
	var data_arr [1000][]int

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")

	for i, line := range lines {
		levels := strings.Split(line, " ")

		for _, level := range levels {
			level_int, err := strconv.Atoi(level)

			if err != nil {
				panic(err)
			}

			data_arr[i] = append(data_arr[i], level_int)
		}
	}

	return data_arr
}

func isSafe(levels []int) bool {

	is_increasing := levels[0]-levels[1] < 0 //assumption

	if is_increasing {
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i] - levels[i+1]

			if diff > -1 || diff < -3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i] - levels[i+1]

			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	return true
}

func countSafeReports(reports [1000][]int) int {
	count := 0

	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	return count
}

func main() {
	reports := convertDataToArr("input.txt")

	// part 1
	safe_report_count := countSafeReports(reports)
	fmt.Println("Safe Report Count:", safe_report_count)

	// part 2
}
