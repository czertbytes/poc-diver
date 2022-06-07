package main

import "regexp"

type RegexpConstraint struct {
	Re      *regexp.Regexp
	Default string
}

func NewRegexpConstraint(re string, defaultValue string) *RegexpConstraint {
	return &RegexpConstraint{
		Re:      regexp.MustCompile(re),
		Default: defaultValue,
	}
}

func (c *RegexpConstraint) Run(v string) (string, error) {
	if v == "" {
		return c.Default, nil
	}

	if !c.Re.MatchString(v) {
		return v, ErrValueNotMatchingRegExp
	}

	return v, nil
}
