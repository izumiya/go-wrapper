package wrapper

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	body := "Hello %v"
	label := "foo"
	w := New(fmt.Sprintf(body, "{{"+label+"}}"))
	if w.Labels[0] != label {
		t.Errorf("want %v but %v", label, w.Labels[0])
	}
	wrapped := fmt.Sprintf(body, placeholder)
	if w.Body != wrapped {
		t.Errorf("want %v but %v", wrapped, w.Body)
	}
}

func TestWrapper_Extract(t *testing.T) {
	m := Map{Key: "foo", Value: "bar"}
	w := New("Hello {{foo}}")
	e := w.Extract("Hello bar")
	if e[0] != m {
		t.Errorf("want %v but %v", m, e[0])
	}
}
