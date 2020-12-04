// +build !debug

package main

import (
	"sync"
)

type Cart struct {
	items map[string]Item
	mx sync.Mutex
}

func NewCart() Cart {
	c := Cart{
		items: make(map[string]Item, 3),
		mx: sync.Mutex{},
	}
	return c
}

func (c *Cart) AddItem(key string, num int) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	if num <= 0 {
		return ErrNumBeZero
	}
	if i, ok := items[key]; !ok {
		return ErrItemNotExist
	} else {
		if i.Num <= 0 || i.Num < num {
			return ErrItemNotAvailable
		}
		if ii, ok := c.items[key]; !ok {
			c.items[key] = Item{
				Key: i.Key,
				Num: num,
			}
		} else {
			ii.Num += num
			c.items[key] = ii
		}
		i.Num -= num
		items[key] = i
	}
	return nil
}

func (c *Cart) RemoveItem(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	if i, ok := items[key]; !ok {
		return ErrItemNotExist
	} else {
		if ii, ok := c.items[key]; !ok {
			return nil
		} else {
			i.Num += ii.Num
			items[key] = i
			delete(c.items, key)
		}
	}
	return nil
}