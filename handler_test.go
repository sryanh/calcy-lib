package main

import (
	"errors"
	"os"
	"testing"

	"github.com/smartystreets/training-cohort/calcy-lib/calc"
)

func TestMissingInput(t *testing.T) {
	var values []string
	handle := NewHandler(os.Stdout, calc.Addition{})
	result, err := handle.Handle(values)
	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
	if !errors.Is(err, errTooFewArgs) {
		t.Errorf("Expected %s", errTooFewArgs)
	}

}
func TestMalformedInput(t *testing.T) {
	var values []string
	handle := NewHandler(os.Stdout, calc.Addition{})
	result, err := handle.Handle(values)
	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
	if !errors.Is(err, errTooFewArgs) {
		t.Errorf("Expected %s", errMalformedInput)
	}
}
func TestCorrectInput(t *testing.T) {
	values := []string{"1", "2"}
	handle := NewHandler(os.Stdout, calc.Addition{})
	result, err := handle.Handle(values)
	if result != 0 {
		t.Error("Expected 3, got ", result)
	}
	if !errors.Is(err, nil) {
		t.Errorf("Expected %s", "")
	}
}
