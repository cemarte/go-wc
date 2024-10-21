package main

import (
	"fmt"
	"testing"
	"testing/fstest"
)

func TestCountLines(t *testing.T) {
	fsys := fstest.MapFS{
		"test.txt": {Data: []byte("Hello, World!")},
	}
	wants := "1 test.txt"
	expect, err := CountLines(fsys, []string{"test.txt"})

	if err != nil || wants != expect[0] {
		t.Errorf("wants: %s, got: %s\n", wants, expect[0])
	}
}

func TestCountBytes(t *testing.T) {
	content := []byte("Hello, World!")
	fsys := fstest.MapFS{
		"test.txt": {Data: content},
	}
	wants := fmt.Sprintf("%d test.txt", len(content))
	expect, err := CountBytes(fsys, []string{"test.txt"})

	if err != nil || wants != expect[0] {
		t.Errorf("wants: %s, got: %s\n", wants, expect[0])
	}
}

type TestResult struct {
	file   string
	result string
}

func TestCountWords(t *testing.T) {
	fsys := fstest.MapFS{
		"test.txt":  {Data: []byte("Hello, World!")},
		"test2.txt": {Data: []byte("Hello,  World!")},
	}
	wants := map[string]string{
		"test.txt":  fmt.Sprintf("%d test.txt", 2),
		"test2.txt": fmt.Sprintf("%d test2.txt", 2),
	}
	for name := range fsys {
		expect, err := CountWords(fsys, []string{name})

		if err != nil || wants[name] != expect[0] {
			t.Errorf("wants: %s, got: %s\n", wants[name], expect[0])
		}

	}
}

func TestProcessLine(t *testing.T) {
	line := "This    s  a line !"
	wants := 5
	expect := processLine(line)
	if wants != expect {
		t.Errorf("wants: %d, got: %d\n", wants, expect)
	}

}

func TestCountCharacters(t *testing.T) {
	fsys := fstest.MapFS{
		"test.txt":  {Data: []byte("Hello, World!")},
		"test2.txt": {Data: []byte("Hello,  World!")},
	}
	wants := map[string]string{
		"test.txt":  fmt.Sprintf("%d test.txt", 14),
		"test2.txt": fmt.Sprintf("%d test2.txt", 15),
	}
	for name := range fsys {
		expect, err := CountCharacters(fsys, []string{name})

		if err != nil || wants[name] != expect[0] {
			t.Errorf("wants: %s, got: %s\n", wants[name], expect[0])
		}

	}
}
