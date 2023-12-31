package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

func TestCountBytes(t *testing.T) {
	data := []byte("Hello, World!")
	expected := 13
	result := countBytes(data)
	if result != expected {
		t.Errorf("Expected %d bytes, but got %d", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	data := []byte("Line 1\nLine 2\nLine 3")
	expected := 3
	result := countLines(data)
	if result != expected {
		t.Errorf("Expected %d lines, but got %d", expected, result)
	}
}

func TestCountWords(t *testing.T) {
	data := []byte("This is a test sentence.")
	expected := 5
	result := countWords(data)
	if result != expected {
		t.Errorf("Expected %d words, but got %d", expected, result)
	}
}

func TestPrintCounts(t *testing.T) {
	// Test data
	fileName := "WC_ToolText.txt"

	// Redirect standard output for testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fileData, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	// Run the function with the test data
	printCounts(fileData, fileName)

	// Reset standard output
	w.Close()
	os.Stdout = old

	expectedOutput := "342190 7145 58164 WC_ToolText.txt\n" // Update this with the expected output based on your test data
	actualOutput := readOutput(r)
	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, actualOutput)
	}
}

// Helper function to read output from a pipe
func readOutput(r io.Reader) string {
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
