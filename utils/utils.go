package utils

import "fmt"

const imagePrefix = "exec"

func FormatImageName(registry, language string) string {
	if (language == "") {
		return ""
	}
	if (registry == "") {
		return fmt.Sprintf("%s-%s", imagePrefix, language)
	}
	return fmt.Sprintf("%s/%s-%s", registry, imagePrefix, language)
}