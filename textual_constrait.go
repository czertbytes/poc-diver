package main

type TextualConstraint struct {
	minLength    uint
	maxLength    uint
	defaultValue string
}

func NewTextualConstraint(minLength, maxLength uint, defaultValue string) *TextualConstraint {
	return &TextualConstraint{
		minLength:    minLength,
		maxLength:    maxLength,
		defaultValue: defaultValue,
	}
}

func (c *TextualConstraint) Run(v string) (string, error) {
	if v == "" {
		return c.defaultValue, nil
	}

	l := len(v)
	if l < int(c.minLength) || int(c.maxLength) < l {
		return v, ErrValueOutOfRange
	}

	return v, nil
}
