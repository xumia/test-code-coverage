package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(1, 2), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}

func TestCoverred(t *testing.T) {
	if got, want := Covered(), 1; got != want {
        t.Errorf("Covered method produced wrong result. expected: %d, got: %d", want, got)
    }
}
