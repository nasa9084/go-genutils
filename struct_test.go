package gen_test

import (
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
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
			DocComments: []string{
				"Something is a test struct.",
				"second line of doc comment.",
			},
			Name: "Something",
			Fields: gen.Fields{
				gen.Field{
					Comments: []string{"Foo is something foo"},
					Name:     "Foo",
					Type:     "string",
				},
				gen.Field{
					Comments: []string{"Bar is not a bar", "but bar"},
					Name:     "Bar",
					Type:     "string",
					Tags: map[string]string{
						"json": "barbar",
					},
				},
				gen.Field{
					Name: "Baz",
					Type: "bool",
				},
				gen.Field{
					Comments: []string{"qux is qux"},
					Name:     "Qux",
					Type:     "int",
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
	t.Run("len(Structs)", func(t *testing.T) {
		if len(got) != len(expected) {
			t.Errorf("length of structs invalid: %d != %d", len(got), len(expected))
			return
		}
	})
	for i := range got {
		t.Run("Structs["+strconv.Itoa(i)+"]", func(t *testing.T) {
			t.Run("DocComments", func(t *testing.T) {
				if !reflect.DeepEqual(got[i].DocComments, expected[i].DocComments) {
					t.Errorf("doc comments not expected: %+v != %+v", got[i].DocComments, expected[i].DocComments)
					return
				}
			})
			t.Run("Name", func(t *testing.T) {
				if got[i].Name != expected[i].Name {
					t.Errorf("name not expected: %s != %s", got[i].Name, expected[i].Name)
					return
				}
			})
			t.Run("len(Fields)", func(t *testing.T) {
				if len(got[i].Fields) != len(expected[i].Fields) {
					t.Errorf("length of fields not expected: %d != %d", len(got[i].Fields), len(expected[i].Fields))
					return
				}
			})
			for j := range got[i].Fields {
				t.Run("Fields["+strconv.Itoa(j)+"]", func(t *testing.T) {
					t.Run("Comments", func(t *testing.T) {
						if !reflect.DeepEqual(got[i].Fields[j].Comments, expected[i].Fields[j].Comments) {
							t.Log(got[i].Fields[j].Name)
							t.Errorf("comments not expected: %+v != %+v", got[i].Fields[j].Comments, expected[i].Fields[j].Comments)
							return
						}
					})
					t.Run("Type", func(t *testing.T) {
						if got[i].Fields[j].Type != expected[i].Fields[j].Type {
							t.Errorf("type not expected: %s != %s", got[i].Fields[j].Type, expected[i].Fields[j].Type)
							return
						}
					})
					t.Run("Tags", func(t *testing.T) {
						if !reflect.DeepEqual(got[i].Fields[j].Tags, expected[i].Fields[j].Tags) {
							t.Errorf("tags not expected: %v != %v", got[i].Fields[j].Tags, expected[i].Fields[j].Tags)
							return
						}
					})
				})
			}
		})
	}
	candidatesStr := []struct {
		expected string
	}{
		{"// Something is a test struct.\n// second line of doc comment.\ntype Something struct {\n	// Foo is something foo\n	Foo string\n	// Bar is not a bar\n	// but bar\n	Bar string `json:\"barbar\"`\n	Baz bool\n	// qux is qux\n	Qux int\n}"},
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

func TestFieldString(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got:  gen.Field{}.String(),
			want: "",
		},
		{
			got: gen.Field{
				Name: "Foo",
				Type: "string",
			}.String(),
			want: "Foo string",
		},
		{
			got: gen.Field{
				Name: "Foo",
				Type: "int",
				Tags: map[string]string{"json": "foo"},
			}.String(),
			want: "Foo int `json:\"foo\"`",
		},
		{
			got: gen.Field{
				Name:     "Foo",
				Type:     "float64",
				Comments: []string{"Foo is example."},
			}.String(),
			want: "// Foo is example.\nFoo float64",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("unexpected:\n  got:  %s\n  want: %s", tt.got, tt.want)
				return
			}
		})
	}
}
