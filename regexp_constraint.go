package main

import "regexp"

type RegexpConstraint struct {
	re           *regexp.Regexp
	defaultValue string
}

func NewRegexpConstraint(re string, defaultValue string) *RegexpConstraint {
	return &RegexpConstraint{
		re:           regexp.MustCompile(re),
		defaultValue: defaultValue,
	}
}

func (c *RegexpConstraint) Run(v string) (string, error) {
	if v == "" {
		return c.defaultValue, nil
	}

	if !c.re.MatchString(v) {
		return v, ErrValueNotMatchingRegExp
	}

	return v, nil
}
