package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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

	input, err := os.Open("test_data/l50k.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}
	defer input.Close()

	exp, err := ioutil.ReadFile("test_data/l50k.csv")
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
		WithConstraint("UNIT_MULT", NewNumericConstraint[int8](0, 100, 1)),
		//WithConstraint("OBS_VALUE", NewNumericConstraint[float64](-math.MaxFloat64, math.MaxFloat64, 0)),
		WithConstraint("ACCOUNTING_ENTRY", NewEnumConstraint([]string{"A", "B", "C", "D", "FI", "FO", "L", "N", "NI", "NE", "NO"}, "A")),
	)

	input, err := os.Open("test_data/l4m.csv")
	if err != nil {
		t.Fatalf("cannot open test file: %s", err)
	}
	defer input.Close()

	test_result, err := os.Create("test_data/l4m_test_result.csv")
	if err != nil {
		t.Fatalf("storing test result failed %s", err)
	}

	bufSize := 64 * 1024 * 1024
	err = d.Run(bufio.NewReaderSize(input, bufSize), bufio.NewWriterSize(test_result, bufSize))
	if err != nil {
		t.Fatalf("expected no err got %s", err)
	}
	test_result.Sync()
	test_result.Close()

	if !hasSameContent("test_data/l4m.csv", "test_data/l4m_test_result.csv") {
		t.Fatalf("exp and res files are different")
	}
}

func hasSameContent(p1, p2 string) bool {
	hashes := make([]string, 2)
	for i, p := range []string{p1, p2} {
		f, err := os.Open(p)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		hashes[i] = md5hash(f)
	}

	return hashes[0] == hashes[1]
}

func md5hash(r io.Reader) string {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
