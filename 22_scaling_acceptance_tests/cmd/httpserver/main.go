package main

import (
	"net/http"

	"github.com/bernardoer/learn_go_with_tests/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler)
}
