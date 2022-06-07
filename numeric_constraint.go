package main

import (
	"fmt"
	"strconv"
)

type NumericConstraint[T Numeric] struct {
	minValue     T
	maxValue     T
	defaultValue T
}

func NewNumericConstraint[T Numeric](minValue, maxValue, defaultValue T) *NumericConstraint[T] {
	return &NumericConstraint[T]{
		minValue:     minValue,
		maxValue:     maxValue,
		defaultValue: defaultValue,
	}
}

func (c *NumericConstraint[T]) Run(v string) (string, error) {
	var numericValue T
	if v == "" {
		switch t := any(c.defaultValue).(type) {
		case int, int8, int16, int32, int64:
			return fmt.Sprintf("%d", t), nil
		case uint, uint8, uint16, uint32, uint64, uintptr:
			return fmt.Sprintf("%d", t), nil
		case float32, float64:
			return fmt.Sprintf("%f", t), nil
		}
	}

	var (
		i64  int64
		ui64 uint64
		f64  float64
		err  error
	)
	switch any(numericValue).(type) {
	case int, int8, int16, int32, int64:
		i64, err = strconv.ParseInt(v, 10, 64)
		numericValue = T(i64)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		ui64, err = strconv.ParseUint(v, 10, 64)
		numericValue = T(ui64)
	case float32, float64:
		f64, err = strconv.ParseFloat(v, 64)
		numericValue = T(f64)
	}

	if err != nil {
		return v, err
	}

	if numericValue < c.minValue || c.maxValue < numericValue {
		return v, ErrValueOutOfRange
	}

	return v, nil
}
