package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	maxLen := 0
	for _, name := range names {
		if l := len(name); l >= maxLen {
			maxLen = l
		}
	}

	res := make([][]string, maxLen)

	for _, name := range names {
		res[len(name) - 1] = append(res[len(name) - 1], name)
	}

	fmt.Println(res)
}