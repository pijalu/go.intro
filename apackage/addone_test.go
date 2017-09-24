package apackage

import "testing"

func TestAddone(t *testing.T) {
	if 1+1 != Addone(1) {
		t.Error("Add one didn't work as expected !")
	}
}
