package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	firstname string
	lastname  string
}

func main() {
	var filename string
	fmt.Println("Please input filename:")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)

	people := make([]person, 0, 3)

	for scanner.Scan() {
		text := scanner.Text()
		fullName := strings.Fields(text)
		if len(fullName) >= 2 {
			p := person{firstname: fullName[0], lastname: fullName[1]}
			people = append(people, p)
		}
	}

	for _, p := range people {
		fmt.Println(p.firstname, p.lastname)
	}
}
