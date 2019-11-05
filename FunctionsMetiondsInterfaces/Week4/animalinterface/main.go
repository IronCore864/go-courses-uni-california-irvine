package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow *Cow) Eat() string   { return "grass" }
func (cow *Cow) Move() string  { return "walk" }
func (cow *Cow) Speak() string { return "moo" }

func (bird *Bird) Eat() string   { return "worms" }
func (bird *Bird) Move() string  { return "fly" }
func (bird *Bird) Speak() string { return "peep" }

func (snake *Snake) Eat() string   { return "mice" }
func (snake *Snake) Move() string  { return "slither" }
func (snake *Snake) Speak() string { return "hsss" }

var m = make(map[string]Animal)

func addAnimal(name string, animalType string) {
	animals := map[string]Animal{
		"cow":   &Cow{},
		"bird":  &Bird{},
		"snake": &Snake{},
	}
	if _, ok := animals[animalType]; ok {
		m[name] = animals[animalType]
		fmt.Println("Created it!")
	} else {
		fmt.Println("Wrong type!")
	}
}

func queryByName(name string, action string) {
	funcs := map[string]func(Animal) string{
		"eat":   (Animal).Eat,
		"move":  (Animal).Move,
		"speak": (Animal).Speak,
	}
	if _, ok := funcs[action]; ok {
		fmt.Println(funcs[action](m[name]))
	} else {
		fmt.Println("Wrong action!")
	}
}

func main() {
	funcs := map[string]func(string, string){
		"newanimal": addAnimal,
		"query":     queryByName,
	}

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
		if len(inputs) < 3 {
			fmt.Println("Wrong command!")
			continue
		}
		cmd := inputs[0]
		name := inputs[1]
		action := inputs[2]

		if _, ok := funcs[cmd]; ok {
			funcs[cmd](name, action)
		} else {
			fmt.Println("Wrong command!")
		}
	}
}
