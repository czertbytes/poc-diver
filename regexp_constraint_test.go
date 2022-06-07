package main

import "testing"

func TestRegexpConstraint(t *testing.T) {
	tests := []struct {
		v   string
		c   *RegexpConstraint
		exp string
		err error
	}{
		{
			v:   "Foo",
			c:   NewRegexpConstraint(`Fo.*`, "default"),
			exp: "Foo",
			err: nil,
		},
		{
			v:   "LongLongLong",
			c:   NewRegexpConstraint(`foo.*`, "baf"),
			exp: "",
			err: ErrValueNotMatchingRegExp,
		},
		{
			v:   "",
			c:   NewRegexpConstraint(`foo.*`, "default"),
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
