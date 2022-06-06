package main

type TextualConstraint struct {
	MinLength uint
	MaxLength uint
	Default   string
}

func NewTextualConstraint(minLength, maxLength uint, defaultValue string) *TextualConstraint {
	return &TextualConstraint{
		MinLength: minLength,
		MaxLength: maxLength,
		Default:   defaultValue,
	}
}

func (c *TextualConstraint) Run(v string) (interface{}, error) {
	if v == "" {
		return c.Default, nil
	}

	l := len(v)
	if l < int(c.MinLength) || int(c.MaxLength) < l {
		return v, ErrValueOutOfRange
	}

	return v, nil
}
