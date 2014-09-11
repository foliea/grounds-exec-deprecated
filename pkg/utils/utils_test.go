package utils

import (
	"fmt"
	"testing"
)

func TestFormatImageName(t *testing.T) {
	expected := fmt.Sprintf("42/%s-ruby", imagesPrefix)
	if imageName := FormatImageName("42", "ruby"); imageName != expected {
		t.Fatalf("Expected image name '%s'. Got '%s'.", expected, imageName)
	}
}

func TestFormatImageNameEmptyRegistry(t *testing.T) {
	expected := fmt.Sprintf("%s-ruby", imagesPrefix)
	if imageName := FormatImageName("", "ruby"); imageName != expected {
		t.Fatalf("Expected image name '%s'. Got '%s'.", expected, imageName)
	}
}

func TestFormatImageNameEmptyLanguage(t *testing.T) {
	expected := ""
	if imageName := FormatImageName("42", ""); imageName != expected {
		t.Fatalf("Expected image name '%s'. Got '%s'.", expected, imageName)
	}
}

func TestFormatCode(t *testing.T) {
	var (
		code     = "puts \"Hello world\\n\\r\\t\"\r\n\t"
		expected = "puts \"Hello world\\\\n\\\\r\\\\t\"\\r\\n\\t"
	)
	if formated := FormatCode(code); formated != expected {
		t.Fatalf("Expected formated code '%s'. Got '%s'.", expected, formated)
	}
}

func TestFormatStatus(t *testing.T) {
	var (
		statusTable	   = []int{0, 1, 254, 255}
		statusExpected = []int{0, 1, -2, -1}
	)

	for i := range statusTable {
		var (
			formated = FormatStatus(statusTable[i])
			expected = statusExpected[i]
		)
		if formated != expected {
			t.Fatalf("Expected formated status '%s'. Got '%s'.", expected, formated)
		}
	}
}
