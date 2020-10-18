package main

import (
	"github.com/v1990/protoc-gen-goku/goku"
	"os"

	// builtin plugins
	_ "github.com/v1990/protoc-gen-goku/plugins"
)

func main() {
	goku.Run(os.Stdin, os.Stdout, os.Stderr)
}
