package gen_test

import (
	"reflect"
	"testing"

	gen "github.com/nasa9084/go-genutils"
)

func TestLoadPackagesFromPath(t *testing.T) {
	pkgs, err := gen.LoadPackagesFromPath("internal/tests")
	if err != nil {
		t.Fatal(err)
	}
	expected := []gen.Package{
		gen.Package{
			Name: "test",
			Structs: []gen.Struct{
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
				gen.Struct{
					Name: "Foo",
				},
				gen.Struct{
					Name: "Bar",
				},
			},
			Interfaces: []gen.Interface{
				gen.Interface{
					Name: "Anything",
					Methods: gen.Methods{
						gen.Method{
							Name: "Foo",
							Params: gen.Params{
								gen.Param{
									Name: "hoge",
									Type: "*string",
								},
								gen.Param{
									Name: "fuga",
									Type: "bool",
								},
							},
							Results: gen.Results{
								gen.Result{
									Type: "[]*Foo",
								},
								gen.Result{
									Type: "error",
								},
							},
						},
						gen.Method{
							Name: "Bar",
							Params: gen.Params{
								gen.Param{
									Name: "ctx",
									Type: "context.Context",
								},
								gen.Param{
									Name: "baz",
									Type: "string",
								},
								gen.Param{
									Name: "qux",
									Type: "string",
								},
							},
							Results: gen.Results{
								gen.Result{
									Type: "*Bar",
								},
								gen.Result{
									Type: "error",
								},
							},
						},
					},
				},
				gen.Interface{
					Name: "Nothing",
					Interfaces: []gen.Interface{
						gen.Interface{
							Name: "Anything",
						},
					},
					Methods: gen.Methods{
						gen.Method{
							Name: "Baz",
							Params: gen.Params{
								gen.Param{
									Name: "args",
									Type: "...interface{}",
								},
							},
						},
					},
				},
			},
		},
	}
	for i := range pkgs {
		got := pkgs[i]
		want := expected[i]
		if got.Name != want.Name {
			t.Errorf("name not expected: %s != %s", got.Name, want.Name)
			return
		}
		if !reflect.DeepEqual(got.Structs, want.Structs) {
			t.Errorf("structs not expected: %+v != %+v", got.Structs, want.Structs)
			return
		}
		if !reflect.DeepEqual(got.Interfaces, want.Interfaces) {
			t.Errorf("interfaces not expected: %+v != %+v", got.Interfaces, want.Interfaces)
			return
		}
	}
}
