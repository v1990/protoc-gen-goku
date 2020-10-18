package helper

import (
	"strings"
)

func InStrings(x string, list ...string) bool {
	for _, s := range list {
		if x == s {
			return true
		}
	}
	return false
}

type eachLineFunc func(line string) string

func EachLine(s string, f eachLineFunc) string {
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		lines = append(lines, f(line))
	}
	return strings.Join(lines, "\n")
}

type eachLine1Func func(line, arg string) string

func eachLine1(s, arg string, f eachLine1Func) string {
	return EachLine(s, func(line string) string {
		return f(line, arg)
	})
}

func TrimLinePrefix(prefix string, s string) string {
	return eachLine1(s, prefix, strings.TrimPrefix)
}
func TrimLineSuffix(suffix string, s string) string {
	return eachLine1(s, suffix, strings.TrimSuffix)

}
func TrimLineLeft(cutset string, s string) string {
	return eachLine1(s, cutset, strings.TrimLeft)

}
func TrimLineRight(cutset string, s string) string {
	return eachLine1(s, cutset, strings.TrimRight)
}
func WithLinePrefix(prefix string, s string) string {
	return EachLine(s, func(line string) string {
		return prefix + line
	})
}
func WithLineSuffix(suffix string, s string) string {
	return EachLine(s, func(line string) string {
		return line + suffix
	})
}

func WithLineSlash(s string) string {
	s = TrimLineLeft("/ ", s)
	s = TrimLinePrefix("*", s)
	s = strings.TrimSpace(s)
	s = WithLinePrefix("// ", s)
	return s
}

func Join(seq string, args_ ...interface{}) string {
	var items []string
	args := Args(args_).Flatten()
	for _, arg := range args {
		items = append(items, ToString(arg))
	}

	return strings.Join(items, seq)
}
