package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int)
)

func coinsForUser(name string) int {
	coins := 0
	for _, char := range name {
		switch char {
		case 'a', 'A':
			coins++
		case 'e', 'E':
			coins++
		case 'i', 'I':
			coins += 2
		case 'o', 'O':
			coins += 3
		case 'u', 'U':
			coins += 4
		}
	}

	return coins
}

func main() {
	for _, name := range users {
		v := coinsForUser(name)

		if v > 10 {
			v = 10
		}

		distribution[name] = v
		coins = coins - v
	}

	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}