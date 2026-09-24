package main_test

import (
	"testing"

	"github.com/bernardoer/learn_go_with_tests/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_spects_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
