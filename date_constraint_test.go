package main

import (
	"testing"
	"time"
)

func TestDateConstraint(t *testing.T) {
	from := time.Now().Add(-30 * 24 * time.Hour)       // -30 days
	to := time.Now().Add(5 * 12 * 30 * 24 * time.Hour) // +5 years
	tests := []struct {
		v   string
		c   *DateConstraint
		exp string
		err error
	}{
		{
			v:   "2022-11-22",
			c:   NewDateConstraint(from, to, "1111-11-11"),
			exp: "2022-11-22",
			err: nil,
		},
		{
			v:   "foobarbaz",
			c:   NewDateConstraint(from, to, "1111-11-11"),
			exp: "",
			err: ErrInvalidDateValue,
		},
		{
			v:   "",
			c:   NewDateConstraint(from, to, "1111-11-11"),
			exp: "1111-11-11",
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
