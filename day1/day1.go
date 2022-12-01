package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var elves []int
	sum := 0

	for scanner.Scan() {

		cals := scanner.Text()

		if len(cals) > 0 {
			converted, err := strconv.Atoi(cals)
			if err != nil {
				log.Fatal(err)
			} else {
				sum += converted
			}
		} else {
			elves = append(elves, sum)
			sum = 0
		}

	}
	sort.Ints(elves)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		sum = 0
		for _, value := range elves[len(elves) - 3:] {
		    sum += value
		}
		fmt.Println("Elves :", sum)
	}
}
