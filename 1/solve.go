package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInt(s string) int {
	result, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return result
}

func sum(items []string) int {
	total := 0
	for i := 0; i < len(items); i++ {
		total += getInt(items[i])
	}

	return total
}

func one(items []string) int {
	total := 0
	for i := 1; i < len(items); i++ {
		intOne := getInt(items[i-1])
		intTwo := getInt(items[i])

		if intOne < intTwo {
			total += 1
		}
	}
	return total
}

func two(items []string) int {
	total := 0
	for i := 4; i < len(items); i++ {
		windowOne := sum(items[i-3 : i])
		windowTwo := sum(items[i-2 : i+1])
		fmt.Printf("%d:%d\n", windowOne, windowTwo)
		if windowOne < windowTwo {
			total += 1
		}
	}
	return total
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	items := strings.Fields(string(dat))

	answer := one(items)

	answerTwo := two(items)

	fmt.Printf("one %d \n", answer)
	fmt.Printf("two %d \n", answerTwo)
}
