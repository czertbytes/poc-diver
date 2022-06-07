package main

type EnumConstraint[T Numeric | Textual] struct {
	Values 	  []T
	Default   T
}

func NewEnumConstraint[T Numeric | Textual](values []T, defaultValue T) *EnumConstraint[T] {
	return &EnumConstraint{
		Values:    values,
		Default:   defaultValue,
	}
}

func (c *EnumConstraint) Run(v string) (string, error) {
	if v == "" {
		return c.Default, nil
	}

	var s string
	for _, e := range c.Values {
		s = fmt.Sprintf("%s", e)
		if s == v {
			return v, nil
		}
	}

	return v, ErrInvalidEnumValue
}
