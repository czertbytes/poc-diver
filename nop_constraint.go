package main

type NopConstraint struct {
}

func NewNopConstraint() *NopConstraint {
	return &NopConstraint{}
}

func (c *NopConstraint) Run(v string) (string, error) {
	return v, nil
}
