package main

import (
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

var ErrValueOutOfRange error = errors.New("constraint: value out of allowed range")
var ErrInvalidEnumValue error = errors.New("constraint: invalid enum value")
var ErrValueNotMatchingRegExp error = errors.New("constraint: value is not matching regexp")

type DataTypes interface {
	Numeric | Textual
}

type Numeric interface {
	constraints.Signed | constraints.Unsigned | constraints.Float
}

type Textual interface {
	string
}

type DataConstraintRunner interface {
	Run(string) (string, error)
}
