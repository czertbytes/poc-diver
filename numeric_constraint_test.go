package main

import "testing"

func TestNumericConstraint(t *testing.T) {
	type NumericTest[T Numeric] struct {
		v   string
		c   *NumericConstraint[T]
		exp string
		err error
	}

	tests := []NumericTest[int32]{
		{
			v:   "126",
			c:   NewNumericConstraint[int32](0, 256, 0),
			exp: "126",
			err: nil,
		},
		{
			v:   "-5",
			c:   NewNumericConstraint[int32](0, 256, 0),
			exp: "0",
			err: ErrValueOutOfRange,
		},
		{
			v:   "",
			c:   NewNumericConstraint[int32](0, 256, 15),
			exp: "15",
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
