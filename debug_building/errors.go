package main

import (
	"fmt"
)

var (
	ErrItemNotExist = fmt.Errorf("item was not exist")
	ErrItemNotAvailable = fmt.Errorf("item was not available")
	ErrNumBeZero = fmt.Errorf("num should be bigger than 0")
)