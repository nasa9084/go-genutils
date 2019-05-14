package gen_test

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	gen "github.com/nasa9084/go-genutils"
)

func TestLoadStructs(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "internal/tests/test_structs.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	got, err := gen.LoadStructs(f)
	if err != nil {
		t.Error(err)
		return
	}
	expected := []gen.Struct{
		gen.Struct{
			DocComment: "Something is a test struct.\nsecond line of doc comment.\n",
			Name:       "Something",
			Fields: gen.Fields{
				gen.Field{
					Comment: "Foo is something foo\n",
					Name:    "Foo",
					Type:    "string",
				},
				gen.Field{
					Comment: "Bar is not a bar\nbut bar\n",
					Name:    "Bar",
					Type:    "string",
					Tags: map[string]string{
						"json": "barbar",
					},
				},
				gen.Field{
					Name: "Baz",
					Type: "bool",
				},
				gen.Field{
					Comment: "qux is qux\n",
					Name:    "Qux",
					Type:    "int",
				},
			},
		},
		gen.Struct{
			Name: "Nested",
			Fields: gen.Fields{
				gen.Field{
					Name: "Foo",
					Type: "struct {\nBar string\n}",
				},
			},
		},
		gen.Struct{
			Name: "Parent",
			Fields: gen.Fields{
				gen.Field{
					Name: "Child",
					Type: "Child",
				},
			},
		},
		gen.Struct{
			Name: "Child",
			Fields: gen.Fields{
				gen.Field{
					Name: "Value",
					Type: "int64",
				},
			},
		},
	}
	if len(got) != len(expected) {
		t.Errorf("length of structs invalid: %d != %d", len(got), len(expected))
		return
	}
	for i := range got {
		if got[i].DocComment != expected[i].DocComment {
			t.Errorf("%s != %s", got[i].DocComment, expected[i].DocComment)
			return
		}
		if got[i].Name != expected[i].Name {
			t.Errorf("%s != %s", got[i].Name, expected[i].DocComment)
			return
		}
		if len(got[i].Fields) != len(expected[i].Fields) {
			t.Errorf("length of fields invalid: %d != %d", len(got[i].Fields), len(expected[i].Fields))
			return
		}
		for j := range got[i].Fields {
			if got[i].Fields[j].Comment != expected[i].Fields[j].Comment {
				t.Errorf("%s != %s", got[i].Fields[j].Comment, expected[i].Fields[j].Comment)
				return
			}
			if got[i].Fields[j].Type != expected[i].Fields[j].Type {
				t.Errorf("%s != %s", got[i].Fields[j].Type, expected[i].Fields[j].Type)
				return
			}
			if !reflect.DeepEqual(got[i].Fields[j].Tags, expected[i].Fields[j].Tags) {
				t.Errorf("%v != %v", got[i].Fields[j].Tags, expected[i].Fields[j].Tags)
				return
			}
		}
	}
	candidatesStr := []struct {
		expected string
	}{
		{"// Something is a test struct.\n// second line of doc comment.\ntype Something struct {\n	// Foo is something foo\n	Foo string\n	// Bar is not a bar\n	// but bar\n	Bar string `json: \"barbar\"`\n	Baz bool\n	// qux is qux\n	Qux int\n}"},
		{`type Nested struct {
	Foo struct {
		Bar string
	}
}`},
		{`type Parent struct {
	Child Child
}`},
		{`type Child struct {
	Value int64
}`},
	}
	for i := range got {
		gotStr := got[i].String()
		expectedStr := candidatesStr[i].expected
		if gotStr != expectedStr {
			t.Errorf("the result of String() is invalid: %s != %s", expectedStr, gotStr)
			return
		}
	}
}
