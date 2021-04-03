package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNum = ":8080"

// HomeEx - / main route request handler
func HomeEx(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Here we are on the Home Page")
}

// AboutEx - /about route request handler
func AboutEx(rw http.ResponseWriter, r *http.Request) {
	value := addValues(2, 2)
	_, _ = fmt.Fprintf(rw, fmt.Sprintf("here is the about page, and 2 + 2 is %d", value))
}

func addValues(x, y int) int {
	return x + y
}

// DivideEx - /didivide request handler
func DivideEx(rw http.ResponseWriter, r *http.Request) {
	var x, y float32
	x, y = 10.5, -0.008
	value, err := divide(x, y)
	if err != nil {
		fmt.Fprintf(rw, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(rw, fmt.Sprintf("%f divided by %f is %f", x, y, value))
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("devided by 0")
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/ex", HomeEx)
	http.HandleFunc("/aboutex", AboutEx)
	http.HandleFunc("/divideex", DivideEx)

	fmt.Println(fmt.Sprintf("Server listening on port %s....", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
