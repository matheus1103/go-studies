package main

import (
	"errors"
	"log"
)

func main(){

	result, err := divide(5.6, 0)
	if err != nil {
		log.Print(err)
	}
	log.Print(result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}
	result = x/y
	return result, nil
}