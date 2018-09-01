package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)

	now := time.Now().Unix()
	mins := now % 2

	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
		fallthrough
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}