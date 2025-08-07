package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Hello world!")
	exp := 2
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("line1\nline2\nline3")

	exp := 3
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Hello world!")
	exp := 12
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d\n", exp, res)
	}
}
