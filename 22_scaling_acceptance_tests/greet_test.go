package go_specs_greet_test

import (
	"testing"

	gospecsgreet "github.com/bernardoer/learn_go_with_tests/go-specs-greet"
	"github.com/bernardoer/learn_go_with_tests/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(gospecsgreet.Greet),
	)
}
