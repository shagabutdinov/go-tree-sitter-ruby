package ruby

import (
	"github.com/smacker/go-tree-sitter"
	"log"
	"testing"
)

func parse(code string) string {
	parser := sitter.NewParser()
	parser.SetLanguage(GetLanguage())
	tree := parser.Parse([]byte(code))
	defer tree.Delete()
	return tree.RootNode().String()
}

func TestSimpleProgramIsParsed(test *testing.T) {
	actual := parse("variable = 'value'")
	if actual != "(program (assignment (identifier) (string)))" {
		test.Fail()
	}
}
