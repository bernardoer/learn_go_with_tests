package generics_test

import "testing"

func TestAssertFunctinos(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

// stacks LIFO (Last In First Out)

type StackOfInts = Stack
type StackOfStrings = Stack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

// type StackOfInts struct {
// 	values []int
// }

// func (s *StackOfInts) Push(value int) {
// 	s.values = append(s.values, value)
// }

// func (s *StackOfInts) IsEmpty() bool {
// 	return len(s.values) == 0
// }

// func (s *StackOfInts) Pop() (int, bool) {
// 	if s.IsEmpty() {
// 		return 0, false
// 	}

// 	index := len(s.values) - 1
// 	el := s.values[index]
// 	s.values = s.values[:index]
// 	return el, true
// }

// type StackOfStrings struct {
// 	values []string
// }

// func (s *StackOfStrings) Push(value string) {
// 	s.values = append(s.values, value)
// }

// func (s *StackOfStrings) IsEmpty() bool {
// 	return len(s.values) == 0
// }

// func (s *StackOfStrings) Pop() (string, bool) {
// 	if s.IsEmpty() {
// 		return "", false
// 	}
// 	index := len(s.values) - 1
// 	el := s.values[index]
// 	s.values = s.values[:index]
// 	return el, true
// }

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)
		// check if stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add something, then check if it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		//add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again

		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
