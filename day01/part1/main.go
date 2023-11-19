package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	largest := 0
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
			if current > largest {
				largest = current
			}
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

	fmt.Printf("Largest value is %v", largest)

}
