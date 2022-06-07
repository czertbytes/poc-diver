package main

import (
	"time"
)

type DateConstraint struct {
	from         time.Time
	to           time.Time
	defaultValue string
}

func NewDateConstraint(from, to time.Time, defaultValue string) *DateConstraint {
	return &DateConstraint{
		from:         from,
		to:           to,
		defaultValue: defaultValue,
	}
}

func (c *DateConstraint) Run(v string) (string, error) {
	if v == "" {
		return c.defaultValue, nil
	}

	date, err := time.Parse("2006-01-02", v)
	if err != nil {
		return v, ErrInvalidDateValue
	}

	if date.Before(c.from) || date.After(c.to) {
		return v, ErrValueNotMatchingRegExp
	}

	return v, nil
}
