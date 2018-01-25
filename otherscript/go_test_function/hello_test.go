package main

/**
Testing for hello.go
to test these file create new file and named hello_test.go
you must add _test.go if you want to test file
*/

import "testing"

func TestName(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Response from getName is enuxpected value")
	}
}
