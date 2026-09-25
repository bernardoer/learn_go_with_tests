package go_specs_greet

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/bernardoer/learn_go_with_tests/go-specs-greet"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_specs_greet.Greet(name))
}
