package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getListFromFile(file_name string) ([]int, []int) {
	var left_arr []int
	var right_arr []int

	//get file data as array of strings where each string represents a line in the text file
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%q", data)
	data_arr := strings.Split(string(data), "\r\n")

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

	return left_arr, right_arr
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

	left_arr, right_arr := getListFromFile("input.txt")

	//sort splices
	sort.Ints(left_arr)
	sort.Ints(right_arr)

	// calculate sum of differences
	return sumOfDifferences(left_arr, right_arr)
}

func main() {
	// puzzle 1
	total_distance := calcTotalDistance()
	fmt.Println("Total:", total_distance)

	// puzzle 2

}
