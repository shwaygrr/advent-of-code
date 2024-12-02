package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getListsFromFile(file_name string) ([]int, []int) {
	var left_arr []int
	var right_arr []int

	//get file data as array of strings where each string represents a line in the text file
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%q", data)
	lines := strings.Split(string(data), "\r\n")

	// assign values
	for _, line := range lines {
		pair_tuple := strings.Split(line, "   ")

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

func calcTotalDistance(arr1, arr2 []int) int {
	//sort splices
	sort.Ints(arr1)
	sort.Ints(arr2)

	// calculate sum of differences
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

func calcSimilarityScore(left_arr, right_arr []int) int {
	right_freq_map := make(map[int]int)

	for _, num := range right_arr {
		right_freq_map[num]++
	}

	result := 0

	for _, id := range left_arr {
		result += id * right_freq_map[id]
	}

	return result
}

func main() {
	left_arr, right_arr := getListsFromFile("input.txt")

	// part 1
	total_distance := calcTotalDistance(left_arr, right_arr)
	fmt.Println("Total:", total_distance)

	// part 2
	similarity_score := calcSimilarityScore(left_arr, right_arr)
	fmt.Println("Similarity Score:", similarity_score)
}
