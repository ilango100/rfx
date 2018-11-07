package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting rfx on http://localhost:%d...\n", set.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", set.Port), http.HandlerFunc(backHandler))
}
