package helper

import (
	"strings"
	"testing"
)

func Test_Line_(t *testing.T) {
	s := `/* 0
	// a
	/* b
	* c
	# d
	`
	t.Log(TrimLineLeft("/*", s))
	t.Log(TrimLinePrefix("//", TrimLinePrefix(" ", `
	// a
	// b
	`)))

	t.Log(strings.TrimLeft("/*/*//#* a", "/*#"))
	t.Log(strings.TrimPrefix("/*/*//#* a", "/*"))

	t.Log(WithLineSlash(`// a
/* b`))
}
