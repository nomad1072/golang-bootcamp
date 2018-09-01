package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [2]string{"hello", "world"}
	fmt.Println(b)

	c := [...]string{"hello", "world"}
	fmt.Printf("%q\n", c)

	var ndim [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			ndim[i][j] = fmt.Sprintf("row %d - column %d", i+i, j+1)
		}
	}
	fmt.Printf("%q\n", ndim)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)

	fmt.Println(mySlice[1:4])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[4:])

	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"

	fmt.Printf("%q\n", cities)

	otherCities := []string{}
	otherCities = append(otherCities, "San Diego", "Mountain View")
	fmt.Println(otherCities)

	// Appending a slice to another slice using ellipsis:
	moreCities := []string{"Santa Monica", "Mountain View"}
	otherCities = append(otherCities, moreCities...)
	fmt.Println(otherCities)

	fmt.Println(len(otherCities))

	countries := make([]string, 42)
	fmt.Println(len(countries))
	fmt.Println(cap(countries))

	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("z == nil is true")
	}
}