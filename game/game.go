package main

import (
	"fmt"
	"log"
)

func main() {
	var i Item
	fmt.Println("i:", i)
	fmt.Printf("v: %v\n", i)
	fmt.Printf("+v: %+v\n", i)
	fmt.Printf("#v: %#v\n", i)

	i = Item{1, 2} // in this style must provide all fields in order
	fmt.Printf("#v: %#v\n", i)

	i = Item{
		Y: 10,
		X: 20,
	}
	fmt.Printf("#v: %#v\n", i)

	i = Item{
		Y: 10, // X isn't given so defaults to the default for the type (0 for int)
	}
	fmt.Printf("#v: %#v\n", i)

	i.X = 9
	fmt.Printf("#v: %#v\n", i)

	fmt.Println(NewItem(100, 200))
	// throws error
	// fmt.Println(NewItem(100, 2000))

	i2, err := NewItem(50, 60)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(i2.X)

	i.Move(15, 17)
	fmt.Printf("after move #v: %#v\n", i)

	p1 := Player{
		Name: "Parzival",
		Item: Item{
			X: 500,
			Y: 700,
		},
	}
	fmt.Printf("p1: %#v\n", p1)
	// fmt.Printf("p1.X: %#v\n", p1.X) // only works if it isn't ambiguous based on embedded types
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	fmt.Printf("p1.First.X: %#v\n", p1.First.X)
	fmt.Printf("p1.First.Other.X: %#v\n", p1.Other.Foo.X)

	ms := []Mover{
		&i, // now isn't a receiver;  when using interfaces like this have to be explicit with passing pointer per Move def
		i2,
		&p1,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}
}

type Player struct {
	Name string
	Item // Player embeds Item
	Other
	First  Item
	Second Item
}

/*
- Add a Keys []string field to Player
- Add a Found(key string) error method to Player
	- key should be one of "copper", "jade", "crystal"
	- a key should be added only once
		- p1.Found("copper")
		- p1.Found("copper")
		- printing p1.Keys ----> [copper]
*/

/*
Go
type Reader interface {
	Read([]byte) (int, error)
}
^^ for the purpose of performance
   re-use same byte slice again and again

Python
type Reader interface {
	Read(int) ([]byte, error)
}

Go
wanted sortable but at the time there were no generics, just interfaces
func Sort(s Sortable) {
	...
}

// See sort package in Go, Sortable.Interface
type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

*/

// Mover
// interfaces specify what we need, not what we provide
// this is inverted from many other languages;  "I need something that can Move"
// so interfaces in Go are small
// rule of thumb: accept interfaces, return types (not interfaces)
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Move i is called "the receiver"
// with func (i Item) Move(x, y int) {
// everything is being passed by value including receivers...so what the Move method receives
// is a copy
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds for %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	// the Go compiler does escape analysis and will allocate i on the heap (instead of on the stack)
	// this way the allocated object can be garbage collected
	// try running "go build -gcflags=-m"
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

type Item struct {
	X int
	Y int
}

type Other struct {
	X   float32
	Foo Item
}
