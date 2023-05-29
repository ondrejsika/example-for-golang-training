package hello_test

import (
	"hello-world/pkg/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Hello("Dela")
	want := "Hello Dela!"
	if got != want {
		t.Errorf(`Hello("Dela") = "%s"; want "%s"`, got, want)
	}
}
