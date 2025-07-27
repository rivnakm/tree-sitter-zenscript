package tree_sitter_zenscript_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_zenscript "github.com/rivnakm/tree-sitter-zenscript/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_zenscript.Language())
	if language == nil {
		t.Errorf("Error loading Zenscript grammar")
	}
}
