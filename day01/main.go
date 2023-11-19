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

	fmt.Printf("All values %v", values)

}
