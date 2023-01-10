package main

import (
	"fmt"
	"net/http"

	router "example.com/practice/router"
)

func main() {
	printBanner()
	router.StartRouter()
	// http.HandleFunc("/", index)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
}

func printBanner() {
	fmt.Println("=======================================================")
	fmt.Println("=================== START SERVER ======================")
	fmt.Println("=======================================================")
}
