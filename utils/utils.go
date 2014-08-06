package utils

import (
	"fmt"
	"strconv"
)

const imagePrefix = "exec"

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
	fmt.Printf("%v", strconv.Quote(code))
	return strconv.Quote(code)
}
