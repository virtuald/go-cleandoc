package cleandoc_test

import (
	"fmt"
	"github.com/virtuald/go-cleandoc"
)

func ExampleCleandoc() {
	cleaned := cleandoc.Cleandoc("An\n    indented\n    docstring.")
	fmt.Println(cleaned)
	// Output: An
	// indented
	// docstring.
}

func ExampleCleandoc_rawstring() {
	s := `
	An
	indented

	rawstring.
	`

	cleaned := cleandoc.Cleandoc(s)
	fmt.Println(cleaned)
	// Output: An
	// indented
	//
	// rawstring.
}
