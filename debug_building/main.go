package main

import (
	"os"
)

type Item struct {
	Key string
	Num int
	Price int
}

var items = map[string]Item{
	"apple": {
		Key: "apple",
		Num: 10,
		Price: 20,
	},
	"banana": {
		Key: "banana",
		Num: 2,
		Price: 50,
	},
	"guava": {
		Key: "guava",
		Num: 7,
		Price: 15,
	},
}

func main() {
	var err error
	cart := NewCart()
	err = cart.AddItem("apple", 1)
	err = cart.AddItem("banana", 3)
	err = cart.AddItem("guava", -1)
	err = cart.AddItem("banana", 2)
	err = cart.AddItem("strawberry", 1)
	err = cart.RemoveItem("banana")
	err = cart.AddItem("banana", 2)
	if err != nil {
		os.Exit(2)
	}
	os.Exit(0)
}