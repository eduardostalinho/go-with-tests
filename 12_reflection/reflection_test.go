package reflection

import (
	"reflect"
	"testing"
)

type profile struct {
	nickname string
}

type nestedStruct struct {
	name    string
	Profile profile
}

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
				firstName string
				lastName  string
			}{"Eduardo", "Carvalho"},
			[]string{"Eduardo", "Carvalho"},
		},
		{
			"Struct with string and int fields",
			struct {
				name string
				age  int
			}{"Eduardo", 22},
			[]string{"Eduardo"},
		},
		{
			"Struct with nested struct",
			nestedStruct{"Eduardo", profile{"stalinho"}},
			[]string{"Eduardo", "stalinho"},
		},
		{
			"Handle pointers",
			&nestedStruct{"Eduardo", profile{"stalinho"}},
			[]string{"Eduardo", "stalinho"},
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
