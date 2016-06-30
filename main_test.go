package main_test

import "testing"

func Test_pass(t *testing.T) {
	if false {
		t.Fatal("Should never reach here.")
	}
}
