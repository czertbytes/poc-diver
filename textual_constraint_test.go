package main

import "testing"

func TestTextualConstraint(t *testing.T) {
	tests := []struct {
		v   string
		c   *TextualConstraint
		exp string
		err error
	}{
		{
			v:   "Foo",
			c:   NewTextualConstraint(0, 5, "default"),
			exp: "Foo",
			err: nil,
		},
		{
			v:   "LongLongLong",
			c:   NewTextualConstraint(0, 3, "baf"),
			exp: "",
			err: ErrValueOutOfRange,
		},
		{
			v:   "",
			c:   NewTextualConstraint(0, 256, "default"),
			exp: "default",
			err: nil,
		},
	}

	for _, test := range tests {
		res, err := test.c.Run(test.v)
		if err != test.err {
			t.Fatalf("for '%s' expected err %s got %s", test.v, test.err, err)
		}
		if test.err == nil && res != test.exp {
			t.Fatalf("for '%s' expected %s got %s", test.v, test.exp, res)
		}
	}
}
