package util

import (
	"fmt"
	"strings"
)

func TrimText(t string) string {
	t = strings.TrimSpace(t)
	l := len(t)

	if l > 25 {
		t = fmt.Sprint(t[:22], "...")
	}

	return t
}
