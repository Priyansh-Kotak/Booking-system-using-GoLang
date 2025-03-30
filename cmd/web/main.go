package main

import (
	"fmt"
	"net/http"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/handlers"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting port at ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
