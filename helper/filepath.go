package helper

import (
	"os"
	"path/filepath"
	"strings"
)

func init() {
	RegisterFuncMap(map[string]interface{}{
		"baseFile":     filepath.Base,
		"baseFileName": BaseFileName,
		"fileExists":   FileExists,
	})
}

// BaseFileName returns the last path element of the name, with the last dotted suffix removed.
// a/b/c.d => c
func BaseFileName(name string) string {
	// First, find the last element
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	// Now drop the suffix
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[0:i]
	}
	return name
}

// 判断文件是否存在
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
