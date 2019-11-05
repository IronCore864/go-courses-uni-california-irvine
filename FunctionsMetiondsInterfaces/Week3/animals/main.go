package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	m := map[string]Animal{
		"cow":   Animal{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  Animal{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> Please input the name of the animal and the info you want.")
	fmt.Println("> Format: name info")
	fmt.Println("> Possible names are: cow, bird, and snake.")
	fmt.Println("> Possible infos are: eat, move, and speak.")
	fmt.Println("> Example:")
	fmt.Println("> 	cow speak")
	for {
		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		inputs := strings.Fields(input)
		name := inputs[0]
		info := inputs[1]

		animal := m[name]

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
