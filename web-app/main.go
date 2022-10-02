package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2+2 is: %d", sum))
}

// addValues adds two integer and return the sum
func addValues(x, y int) int {
	return x + y
}

// Divide is a page
func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32

	x = 100.0
	y = 0.0
	f, err := division(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is equal to %f", x, y, f))
}

// division divides two float32 and return de result and the error if exists
func division(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
