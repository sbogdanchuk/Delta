package main

import (
	"testing"
)

func TestConv(t *testing.T) {
	expected := "1:09"
	actual := conv(69)
	if actual != expected {
		t.Errorf("Convert was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestConvNegative(t *testing.T) {
	expected := "-(1:09)"
	actual := conv(-69)
	if actual != expected {
		t.Errorf("Convert was incorrect, got: %v, want: %v.", actual, expected)
	}
}
