package main

import (
	"bytes"
	"testing"
)

func TestAddEmptyButton(t *testing.T) {
	expected := []byte("<button class=\"old-style\">")
	result := addClassToTag([]byte("<button>"), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestAddButtonWithClass(t *testing.T) {
	expected := []byte("<button class=\"old-style some-class good\">")
	result := addClassToTag([]byte("<button class=\"some-class good\">"), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestAddButtonWithClassAndNewline(t *testing.T) {
	expected := []byte("<button\n\n\t\tclass=\"old-style some-class good\"")
	result := addClassToTag([]byte("<button\n\n\t\tclass=\"some-class good\""), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestAddInputWithClassAndNewline(t *testing.T) {
	expected := []byte("<input\n\n\t\tclass=\"old-style some-class good\"")
	result := addClassToTag([]byte("<input\n\n\t\tclass=\"some-class good\""), "input", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}
