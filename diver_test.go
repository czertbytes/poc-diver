package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDiver(t *testing.T) {
	tests := []struct {
		data  io.Reader
		diver *Diver
		exp   string
		err   error
	}{
		{
			data: strings.NewReader(`c1,c2,c3
1,2,3
3,4,5
5,6,`),
			diver: NewDiver(
				WithConstraint("c1", NewNumericConstraint[int32](0, 10, 5)),
				WithConstraint("c3", NewNumericConstraint[int32](0, 10, 2)),
			),
			exp: `c1,c2,c3
1,2,3
3,4,5
5,6,2
`,
			err: nil,
		},
	}

	for _, test := range tests {
		var output strings.Builder
		err := test.diver.Run(test.data, &output)
		if err != test.err {
			t.Fatalf("expected err %s got %s", test.err, err)
		}
		if test.err == nil {
			res := output.String()
			if res != test.exp {
				t.Fatalf("expected %s got %s", test.exp, res)
			}
		}
	}
}

func TestDiverSmallFile(t *testing.T) {
	d := NewDiver(
		WithConstraint("DATA_TYPE_BOP", NewNumericConstraint[int8](0, 5, 2)),
	)

	input, err := os.Open("test_data/test.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}
	defer input.Close()

	exp, err := ioutil.ReadFile("test_data/test.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}

	var output strings.Builder
	err = d.Run(input, &output)
	if err != nil {
		t.Fatalf("expected no err got %s", err)
	}

	res := output.String()
	if res != string(exp) {
		t.Fatalf("expected %s got %s", string(exp), res)
	}
}

func TestDiverLargeFile(t *testing.T) {
	d := NewDiver(
		WithConstraint("CB_REP_SECTOR", NewNumericConstraint[int8](1, 100, 1)),
	)

	input, err := os.Open("test_data/test_large.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}
	defer input.Close()

	exp, err := ioutil.ReadFile("test_data/test_large.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}

	var res bytes.Buffer
	err = d.Run(input, &res)
	if err != nil {
		t.Fatalf("expected no err got %s", err)
	}

	if !reflect.DeepEqual(res.Bytes(), exp) {
		t.Fatalf("it's different")
	}
}
