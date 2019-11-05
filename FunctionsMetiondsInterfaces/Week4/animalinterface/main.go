package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow *Cow) Eat()   { fmt.Println("grass") }
func (cow *Cow) Move()  { fmt.Println("walk") }
func (cow *Cow) Speak() { fmt.Println("moo") }

func (bird *Bird) Eat()   { fmt.Println("worms") }
func (bird *Bird) Move()  { fmt.Println("fly") }
func (bird *Bird) Speak() { fmt.Println("peep") }

func (snake *Snake) Eat()   { fmt.Println("mice") }
func (snake *Snake) Move()  { fmt.Println("slither") }
func (snake *Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> Please input command.")
	fmt.Println("> Format: newanimal name type")
	fmt.Println("> Possible types are: cow, bird, and snake.")
	fmt.Println("> Format: query name info.")
	fmt.Println("> Possible infos are: eat, move, and speak.")
	fmt.Println("> Example:")
	fmt.Println("> 	newanimal Bobby cow")
	fmt.Println("> 	query Bobby speak")

	for {
		fmt.Printf("> ")
		input, _ := reader.ReadString('\n')
		inputs := strings.Fields(input)

		cmd := inputs[0]
		name := inputs[1]
		action := inputs[2]

		switch cmd {
		case "newanimal":
			switch action {
			case "cow":
				animals[name] = &Cow{}
			case "bird":
				animals[name] = &Bird{}
			case "snake":
				animals[name] = &Snake{}
			}
			fmt.Println("Created it!")
		case "query":
			switch action {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			}
		}
	}
}
