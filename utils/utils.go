package utils

import (
	"fmt"
	"strings"
)

const (
	imagePrefix = "exec"
)

var codeReplacements = [][]string{
	{"\\", "\\\\"},
	{"\n", "\\n"},
	{"\r", "\\r"},
	{"\t", "\\t"},
}

func FormatImageName(registry, language string) string {
	if language == "" {
		return ""
	}
	if registry == "" {
		return fmt.Sprintf("%s-%s", imagePrefix, language)
	}
	return fmt.Sprintf("%s/%s-%s", registry, imagePrefix, language)
}

func FormatCode(code string) string {
	for _, replacement := range codeReplacements {
		code = strings.Replace(code, replacement[0], replacement[1], -1)
	}
	return code
}
