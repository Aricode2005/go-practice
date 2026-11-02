package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Counts
	}{
		{"empty", "", Counts{0, 0, 0}},
		{"one line", "hello world\n", Counts{1, 2, 12}},
		{"no trailing newline", "a b", Counts{0, 2, 3}},
		{"multiple spaces", "a   b\n", Counts{1, 2, 6}},
		{"unicode", "héllo\n", Counts{1, 1, 7}},
		{"tab and CRLF", "a\tb\r\n", Counts{1, 2, 5}},
		{"only whitespace", "  \n\n", Counts{2, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Count(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestCountFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "x.txt")
	if err := os.WriteFile(path, []byte("one two\nthree\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := countFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := Counts{Lines: 2, Words: 3, Bytes: 14}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestCountFileMissing(t *testing.T) {
	_, err := countFile(filepath.Join(t.TempDir(), "nope.txt"))
	if err == nil {
		t.Fatal("expected an error for a missing file")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected ErrNotExist, got %v", err)
	}
}
