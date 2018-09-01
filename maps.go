package main

import "fmt"

func main() {
	celebs := map[string]int{
		"Nicolas Cage": 50,
		"Selena Gomez": 21,
		"Jude Law": 41,
		"Scarlett Johansson": 29,
	}

	fmt.Printf("%#v\n", celebs)

	emptyMap := map[string]int{}
	emptyMap["Hello"] = 1
	fmt.Println(emptyMap)

	elem, ok := emptyMap["Hello"]
	fmt.Println(elem, ok)

	delete(emptyMap, "Hello")
	fmt.Println(emptyMap)

	elem, ok = emptyMap["Hello"]
	fmt.Println(elem, ok)
}