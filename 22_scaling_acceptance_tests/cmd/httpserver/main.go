package main

import (
	"net/http"

	go_specs_greet "github.com/bernardoer/learn_go_with_tests/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	http.ListenAndServe(":8080", handler)
}
