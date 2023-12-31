package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var values []int
	current := 0

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		s := string(line)

		if s == "" {
			fmt.Println("Empty line")
			values = append(values, current)
			current = 0
		} else {
			fmt.Println(string(s))
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Bad number in file")
			}
			current += i
		}
	}

	fmt.Printf("All values are %v\n\n\n", values)

	slices.Sort(values)

	fmt.Printf("Sorted values are %v\n\n\n\n", values)

	slices.Reverse(values)

	fmt.Printf("Reversed values are %v\n\n\n", values)

	fmt.Printf("Highest 3 add up to %d", values[0]+values[1]+values[2])
}
