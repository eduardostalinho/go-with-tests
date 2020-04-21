package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ name string }{"Eduardo"},
			[]string{"Eduardo"},
		},

		{
			"Struct with two string fields",
			struct {
				fisrtName string
				lastName  string
			}{"Eduardo", "Carvalho"},
			[]string{"Eduardo", "Carvalho"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})

	}

}
