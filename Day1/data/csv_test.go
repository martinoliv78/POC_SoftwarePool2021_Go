package data

import "testing"

func TestLineToCSV(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in:  "jean,34,france",
			out: []string{"jean", "34", "france"},
			err: false,
		},
		{
			in:  "j,j,j,j",
			out: nil,
			err: true,
		},
		{
			in:  "j",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := LineToCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in:  "../go.mod",
			out: []string{"module SofwareGoDay1", "go 1.15"},
			err: false,
		},
		{
			in:  "lalala",
			out: nil,
			err: true,
		},
		{
			in:  "pouic",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := ReadFile(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error at [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error at [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}
