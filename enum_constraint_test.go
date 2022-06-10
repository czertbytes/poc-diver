package main

import "testing"

func TestEnumStringConstraint(t *testing.T) {
	type EnumTest[T DataTypes] struct {
		v   string
		c   *EnumConstraint[T]
		exp string
		err error
	}

	tests := []EnumTest[string]{
		{
			v:   "Y",
			c:   NewEnumConstraint([]string{"Y", "N"}, "N"),
			exp: "Y",
			err: nil,
		},
		{
			v:   "X",
			c:   NewEnumConstraint([]string{"Y", "N"}, "N"),
			exp: "X",
			err: ErrInvalidEnumValue,
		},
		{
			v:   "",
			c:   NewEnumConstraint([]string{"Y", "N"}, "N"),
			exp: "N",
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

func TestEnumIntConstraint(t *testing.T) {
	type EnumTest[T DataTypes] struct {
		v   string
		c   *EnumConstraint[T]
		exp string
		err error
	}

	tests := []EnumTest[int8]{
		{
			v:   "1",
			c:   NewEnumConstraint([]int8{1, 2, 3}, 2),
			exp: "1",
			err: nil,
		},
		{
			v:   "5",
			c:   NewEnumConstraint([]int8{1, 2, 3}, 2),
			exp: "5",
			err: ErrInvalidEnumValue,
		},
		{
			v:   "",
			c:   NewEnumConstraint([]int8{1, 2, 3}, 2),
			exp: "2",
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
