package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type step struct {
	dir      string
	distance int
}

func getInt(s string) int {
	result, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return result
}

func getInput(path string) []step {
	inputString, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}
	rows := strings.Split(string(inputString), "\n")
	results := make([]step, len(rows))
	for i := 0; i < len(rows); i++ {
		parts := strings.Split(rows[i], " ")
		results[i] = step{dir: parts[0], distance: getInt(parts[1])}
	}

	return results
}

func one(input []step) int {
	xPos := 0
	yPos := 0

	for i := 0; i < len(input); i++ {
		theStep := input[i]
		if theStep.dir == "forward" {
			xPos = xPos + theStep.distance
		} else if theStep.dir == "up" {
			yPos = yPos - theStep.distance
		} else if theStep.dir == "down" {
			yPos = yPos + theStep.distance
		}
	}

	return xPos * yPos
}

func two(input []step) int {
	xPos := 0
	yPos := 0
	aim := 0

	for i := 0; i < len(input); i++ {
		theStep := input[i]
		if theStep.dir == "forward" {
			xPos += theStep.distance
			yPos += aim * theStep.distance
		} else if theStep.dir == "up" {
			aim = aim - theStep.distance
		} else if theStep.dir == "down" {
			aim = aim + theStep.distance
		}
	}

	return xPos * yPos
}

func main() {
	inputRows := getInput("./input.txt")

	fmt.Printf("one: %d \n", one(inputRows))
	fmt.Printf("two: %d \n", two(inputRows))
}
