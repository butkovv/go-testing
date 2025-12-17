package even

import "testing"

func TestIsEven(t *testing.T) {
	if got, want := IsEven(10), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := IsEven(-10), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := IsEven(5), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := IsEven(-5), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := IsEven(0), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
