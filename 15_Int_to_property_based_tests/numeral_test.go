package inttopropertybasedtests

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
	// t.Run("1 gets converted to I", func(t *testing.T) {
	// 	got := ConvertToRoman(1)
	// 	want := "I"

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })
	// t.Run("2 gets converted to II", func(t *testing.T) {
	// 	got := ConvertToRoman(2)
	// 	want := "II"

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })
}
