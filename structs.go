package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date time.Time
}

type Point struct {
	x, y int
}

type User struct {
	Id int
	Name string
	Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	User
	GameId int
}

func main() {
	fmt.Println(Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})

	p := Point{1, 2}
	q := &Point{1, 2}
	r := Point{x: 1}
	s := Point{}

	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of q: %T\n", q)
	fmt.Printf("Type of r: %T\n", r)
	fmt.Printf("Type of s: %T\n", s)

	x := new(Bootcamp)
	y := &Bootcamp{}
	y.Lat = 10
	fmt.Println(x.Lat)

	player := &Player{}
	player.Id = 42
	player.Name = "Matt"
	player.Location = "LA"
	fmt.Println(player.Greetings())
}