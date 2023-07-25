package util

import (
	"fmt"
	"strings"
)

func TrimText(t string) string {
	t = strings.TrimSpace(t)
	l := len(t)

	if l > 73 {
		t = fmt.Sprint(t[:70], "...")
	}

	return t
}
