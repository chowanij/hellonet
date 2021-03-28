package main

import (
	"fmt"
	"net/http"
)

const portNum = ":8080"

// Home - / main route request handler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Here we are on the Home Page")
}

// About - /about route request handler
func About(rw http.ResponseWriter, r *http.Request) {
	value := addValues(2, 2)
	_, _ = fmt.Fprintf(rw, fmt.Sprintf("here is the about page, and 2 + 2 is %d", value))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server listening on port %s....", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
