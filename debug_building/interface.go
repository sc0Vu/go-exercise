package main

type ShoppingCart interface {
	AddItem(key string, num int) error
	RemoveItem(key string) error
}