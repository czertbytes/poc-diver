package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Diver struct {
	constraints map[string]DataConstraintRunner
}

func NewDiver(options ...func(*Diver)) *Diver {
	diver := &Diver{
		constraints: map[string]DataConstraintRunner{},
	}
	for _, o := range options {
		o(diver)
	}

	return diver
}

func WithConstraint(column string, dc DataConstraintRunner) func(*Diver) {
	return func(d *Diver) {
		d.constraints[column] = dc
	}
}

func (d *Diver) Run(data io.Reader, output io.Writer) error {
	// TODO: validate input data
	r := csv.NewReader(data)
	w := csv.NewWriter(output)
	w.UseCRLF = false

	defer w.Flush()

	headers, err := r.Read()
	if err != nil {
		fmt.Println(err)
		return err
	}

	constraints := make([]DataConstraintRunner, len(headers))
	for i, h := range headers {
		if c, found := d.constraints[h]; found {
			constraints[i] = c
		} else {
			constraints[i] = NewNopConstraint()
		}
	}
	w.Write(headers)

	var line int64
	row := make([]string, len(headers))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		for i, v := range record {
			row[i], err = constraints[i].Run(v)
			if err != nil {
				fmt.Printf("diver: validating line %d column %d failed: %s", line, i, err)
				return err
			}
		}

		w.Write(row)
		line++
	}

	return nil
}
