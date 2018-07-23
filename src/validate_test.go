package downit

import (
	"strings"
	"testing"
)

func TestURLToName(t *testing.T) {
	tests := []struct {
		URL  string
		Name string
	}{
		{"http://google.com", "google.com"},
		{"https://https://dl.google.com/go/go1.10.3.src.tar.gz", "go1.10.3.src.tar.gz"},
		{"http://github.com/crazcalm/down-it", "down-it"},
	}

	for _, test := range tests {
		result := URLToName(test.URL)
		if !strings.EqualFold(test.Name, result) {
			t.Errorf("For %s, expected %s, but got %s", test.URL, test.Name, result)
		}
	}
}

func TestAddHTTPProtocol(t *testing.T) {
	tests := []struct {
		Name   string
		Answer string
	}{
		{"google.com", "http://google.com"},
	}

	for _, test := range tests {
		result := AddHTTPProtocol(test.Name)
		if !strings.EqualFold(result, test.Answer) {
			t.Errorf("For %s, expected %s, but got %s", test.Name, test.Answer, result)
		}
	}
}

func TestFileName(t *testing.T) {
	tests := []struct {
		Name   string
		Answer string
	}{
		{"hello.go", "hello.go"},
		{"validate.go", "validate.go.1"},
	}

	for _, test := range tests {
		result := FileName(test.Name)
		if !strings.EqualFold(result, test.Answer) {
			t.Errorf("For %s, expected %s, but got %s", test.Name, test.Answer, result)

		}
	}
}

func TestFileExist(t *testing.T) {
	tests := []struct {
		Name   string
		Answer bool
	}{
		{"validate.go", true},
		{"DoesNotExist", false},
	}

	for _, test := range tests {
		result := FileExist(test.Name)
		if result != test.Answer {
			t.Errorf("For %s, expected %t, but got %t", test.Name, test.Answer, result)
		}
	}
}

func TestValidateHTTPProtocol(t *testing.T) {
	tests := []struct {
		URL    string
		Answer bool
	}{
		{"http://google.com", true},
		{"https://google.com", true},
		{"google.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := ValidateHTTPProtocol(test.URL)
		if result != test.Answer {
			t.Errorf("For %s, expected %t, but got %t", test.URL, test.Answer, result)
		}
	}
}
