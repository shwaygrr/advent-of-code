package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertDataToreport(file_name string) [1000][]int {
	var data_report [1000][]int

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

			data_report[i] = append(data_report[i], level_int)
		}
	}

	return data_report
}

func isSafe(report []int) bool {
	is_increasing := report[0]-report[1] < 0 //assumption

	if is_increasing {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]

			if diff > -1 || diff < -3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]

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

func isSafeWithDampener(report []int, assume_increasing bool) bool {
	i := 0
	j := 1
	dampened_index := 80 // guarenteed number thats not an index in the input

	for j < len(report) {
		//assume increasing
		diff := report[i] - report[j]

		var invalid_range bool

		if assume_increasing {
			invalid_range = diff > -1 || diff < -3
		} else {
			invalid_range = diff < 1 || diff > 3
		}

		if invalid_range {
			if dampened_index == 80 { //means first time dampening
				dampened_index = j //track skipped index
				j++                //skip
				// fmt.Println("index skipped for level", report)
			} else {
				return false
			}
		} else {
			if dampened_index == i+1 {
				dampened_index = 81 //prevent triggering another dampening
				i += 2
			} else {
				i++
			}
			j++
		}
	}
	return true
}

func countIsSafeWithDampener(reports [1000][]int) int {
	count := 0

	for _, report := range reports {
		if isSafe(report) || isSafe(report[1:]) || isSafeWithDampener(report, true) || isSafeWithDampener(report, false) {
			count++
		}
	}

	return count
}

func main() {
	reports := convertDataToreport("input.txt")

	// part 1
	safe_report_count := countSafeReports(reports)
	fmt.Println("Safe Report Count:", safe_report_count)

	// part 2
	safe_reports_with_dampener_count := countIsSafeWithDampener(reports)
	fmt.Println("Safe Report Count:", safe_reports_with_dampener_count)
}
