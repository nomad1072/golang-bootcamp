package main

import (
	"fmt"
	"time"
)

const (
	Pi = 3.14
	Truth = false
	Big = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "Hello, World!"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
	fmt.Println(Small)

	fmt.Println(add(42, 13))
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s\n", region, continent)

	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s has released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

	// Type conversions
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Printf("Type of u is %T\n", u)

	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)
}

func add(x, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

type Artist struct {
	Name, Genre string
	Songs int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func timeMap(y interface{}) {
	// Type assertion
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}