package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var name string
	var addr string
	fmt.Println("Please input a name:")
	name, _ = reader.ReadString('\n')
	fmt.Println("Please input an address:")
	addr, _ = reader.ReadString('\n')
	m := map[string]string{
		"name": name,
		"addr": addr,
	}
	jsonObj, _ := json.Marshal(m)
	fmt.Println(jsonObj)
}
