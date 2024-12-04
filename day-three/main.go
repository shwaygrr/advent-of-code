package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getMemoryString(file_name string) string {
	memory, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	return string(memory)
}
func scanandMult(memory string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	valid_sequences := re.FindAllString(memory, -1)

	fmt.Println(valid_sequences)

	result := 0

	re_nums := regexp.MustCompile(`[0-9]{1,3}`)

	for _, sequence := range valid_sequences {
		nums := re_nums.FindAllString(sequence, -1)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		result += num1 * num2
	}

	return result
}

func main() {
	memory := getMemoryString("input.txt")

	// part 1
	uncorrupts_sum := scanandMult(memory)
	fmt.Println(uncorrupts_sum)

	//poart 2
}
