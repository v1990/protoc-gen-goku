package dump

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

type Thing struct {
	ThingID int
}

type Animal struct {
	*Thing
	HP int
}

func TestDump2(t *testing.T) {
	obj := &Animal{
		Thing: &Thing{ThingID: 12345},
		HP:    1e4,
	}

	t.Logf("%#v", obj)
	fmt.Println(Dump2(obj))

}

func TestDump22(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset
	file, err := parser.ParseFile(fset, "dump.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	type T struct {
		MA   map[int]interface{}
		MB   map[string]interface{}
		File *ast.File
	}

	tt := &T{
		MA: map[int]interface{}{
			1: 100,
		},
		MB: map[string]interface{}{
			"a": "aaa",
		},
		File: file,
	}

	//t.Logf("%#v", file)
	fmt.Printf("%#v\n", file)
	fmt.Println("====")
	fmt.Println(Dump2(tt))

}
