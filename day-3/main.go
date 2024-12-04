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

func scanAndMult(memory string) int {
	result := 0

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	valid_sequences := re.FindAllString(memory, -1)

	re_nums := regexp.MustCompile(`[0-9]{1,3}`)

	for _, sequence := range valid_sequences {
		nums := re_nums.FindAllString(sequence, -1)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		result += num1 * num2
	}

	return result
}

func splitConditionals(memory string) []string {
	re := regexp.MustCompile(`do(n't)?\(\)`)

	conds_indexes := re.FindAllStringIndex(memory, -1)

	var split_memory []string

	i := 0

	for _, index_range := range conds_indexes {

		split_memory = append(split_memory, memory[i:index_range[0]])
		i = index_range[0]
	}

	split_memory = append(split_memory, memory[i:])

	return split_memory
}

func scanAndMultWithConditionals(memory string) int {
	result := 0

	split_conditionals := splitConditionals(memory)

	for _, conditional := range split_conditionals {
		if conditional[:7] != "don't()" {
			result += scanAndMult(conditional)
		}
	}

	return result
}

func main() {
	memory := getMemoryString("input.txt")

	// part 1
	uncorrupts_sum := scanAndMult(memory)
	fmt.Println("Multiplications sum:", uncorrupts_sum)

	//part 2
	uncorrupts_sum_with_conditionals := scanAndMultWithConditionals(memory)
	fmt.Println("Multiplications with condtionals sum:", uncorrupts_sum_with_conditionals)
}
