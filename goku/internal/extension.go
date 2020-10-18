package internal

import (
	"bytes"
	"fmt"
	"strings"
)

// 获取不带双斜杠的注释
func (g *Generator) makeOriginComments(path string) (string, bool) {
	loc, ok := g.file.comments[path]
	if !ok {
		return "", false
	}
	w := new(bytes.Buffer)
	nl := ""
	for _, line := range strings.Split(strings.TrimSuffix(loc.GetLeadingComments(), "\n"), "\n") {
		fmt.Fprintf(w, "%s%s", nl, line)
		nl = "\n"
	}
	return w.String(), true
}
