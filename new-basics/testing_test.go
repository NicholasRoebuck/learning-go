package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 1.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expected-data", 10.0, 10.0, 1.0, false},
	{"expected-10", 100.0, 10.0, 10.0, false},
	{"expected-error", -100.0, 0, 0.0, true},
}

func testDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Errorf("did not get expected error")
			}
		} else {
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}
		}

		if got != tt.expected {
			t.Errorf("got %f, want %f", got, tt.expected)
		}

	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestDivideFail(t *testing.T) {
	_, err := divide(10.0, 0)

	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
