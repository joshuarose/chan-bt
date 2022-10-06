package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	equal := Same(tree.New(1), tree.New(1))
	if !equal {
		t.Error("binary trees should be equal")
	}
	notequal := Same(tree.New(1), tree.New(2))
	if notequal {
		t.Error("binary trees should not be equal")
	}
}
