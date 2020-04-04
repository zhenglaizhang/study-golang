package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 10 + 5
	expected := 15
	if got != expected {
		t.Errorf("Dit not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSubstraction(t *testing.T) {
	got := 10 - 5
	expected := 5
	if got != expected {
		t.Errorf("Dit not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}
