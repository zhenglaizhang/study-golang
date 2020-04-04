package messages 
// same package with test targeting functions

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestDepart(t *testing.T) {
	// whitebox, same package, so we can access internal function
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestFailureTypes(t *testing.T) {
	t.Error("Error signals a failed test, but doesn't stop the rest of the test from executing")
	t.Fatal("Fatal will fail the test and stop its execution")
	t.Error("This will never print since it is preceeded by an immediate failure")
}