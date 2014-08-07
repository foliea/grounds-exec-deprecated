package utils

import (
	"fmt"
	"testing"
)

func TestFormatImageName(t *testing.T) {
	expected := fmt.Sprintf("42/%s-ruby", imagePrefix)
	if imageName := FormatImageName("42", "ruby"); imageName != expected {
		t.Fatalf("Expected image name '%s'. Got '%s'.", expected, imageName)
	}
}

func TestFormatImageNameEmptyRegistry(t *testing.T) {
	expected := fmt.Sprintf("%s-ruby", imagePrefix)
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
	code := "puts \"Hello world\\n\"\r\n\t"
	expected := "puts \"Hello world\\\\n\"\\r\\n\\t"
	if formated := FormatCode(code); formated != expected {
		t.Fatalf("Expected formated code '%s'. Got '%s'.", expected, formated)
	}
}
