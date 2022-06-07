package main

import (
	"fmt"
)

type EnumConstraint[T DataTypes] struct {
	Values  []T
	Default T
}

func NewEnumConstraint[T DataTypes](values []T, defaultValue T) *EnumConstraint[T] {
	return &EnumConstraint[T]{
		Values:  values,
		Default: defaultValue,
	}
}

func (c *EnumConstraint[T]) Run(v string) (string, error) {
	if v == "" {
		switch t := any(c.Default).(type) {
		case int, int8, int16, int32, int64:
			return fmt.Sprintf("%d", t), nil
		case uint, uint8, uint16, uint32, uint64, uintptr:
			return fmt.Sprintf("%d", t), nil
		case float32, float64:
			return fmt.Sprintf("%f", t), nil
		case string:
			return t, nil
		}
	}

	var s string
	for _, e := range c.Values {
		switch t := any(e).(type) {
		case int, int8, int16, int32, int64:
			s = fmt.Sprintf("%d", t)
		case uint, uint8, uint16, uint32, uint64, uintptr:
			s = fmt.Sprintf("%d", t)
		case float32, float64:
			s = fmt.Sprintf("%f", t)
		case string:
			s = t
		}

		if s == v {
			return v, nil
		}
	}

	return v, ErrInvalidEnumValue
}
