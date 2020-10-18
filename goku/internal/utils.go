package internal

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func (g *Generator) Debug(format string, args ...interface{}) {
	if g != nil && !g.debug {
		return
	}

	var prefix string
	//_, file, line, ok := runtime.Caller(2)
	//if ok {
	//	prefix = fmt.Sprintf("[%s:%d]", file, line)
	//}

	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(os.Stderr, prefix, msg)
}

func (g *Generator) GetFileDescriptor(protoFilename string) *FileDescriptor {
	return g.fileByName(protoFilename)
}

func (d *FileDescriptor) GetMessageDescriptor(protoMessageName string) *Descriptor {
	for _, descriptor := range d.desc {
		if descriptor.GetName() == protoMessageName {
			return descriptor
		}
	}
	return nil
}
func (d *FileDescriptor) PrintAllComments() string {
	var lines []string
	for path, loc := range d.comments {
		lines = append(lines, fmt.Sprintf("[%s] :\n%s", path, loc.GetLeadingComments()))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n\n\n\n")

}
