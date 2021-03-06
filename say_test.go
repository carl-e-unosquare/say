package say_test

import (
	"github.com/carl-e-unosquare/say"
	"testing"
)

func TestHello(t *testing.T) {
	n := "Bob"
	expected := "Hello Bob."
	actual := say.Hello(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}

func TestBye(t *testing.T) {
	n := "Bob"
	expected := "Good bye Bob."
	actual := say.Bye(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}
