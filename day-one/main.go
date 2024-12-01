package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getFileLines(file_name string) []string {
	//get file data as array of strings where each string represents a line in the text file
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%q", data)
	data_arr := strings.Split(string(data), "\r\n")
	return data_arr
}

func sumOfDifferences(arr1 []int, arr2 []int) int {
	total := 0
	for i := range arr1 {
		diff := arr1[i] - arr2[i]

		if diff < 0 {
			diff *= -1
		}

		total += diff
	}
	return total
}

func calcTotalDistance() int {
	var left_arr []int
	var right_arr []int

	data_arr := getFileLines("input.txt")

	// assign values
	for _, pair := range data_arr {
		pair_tuple := strings.Split(pair, "   ")

		left_item, left_err := strconv.Atoi(pair_tuple[0])
		if left_err != nil {
			panic(left_err)
		}

		right_item, right_err := strconv.Atoi(pair_tuple[1])
		if right_err != nil {
			panic(right_err)
		}

		left_arr = append(left_arr, left_item)
		right_arr = append(right_arr, right_item)
	}

	//sort splices
	sort.Ints(left_arr)
	sort.Ints(right_arr)

	// calculate sum of differences
	return sumOfDifferences(left_arr, right_arr)
}

func main() {
	// Puzzle 1
	total_distance := calcTotalDistance()
	fmt.Println("Total:", total_distance)
}
