package gen_test

import (
	"go/parser"
	"go/token"
	"testing"

	gen "github.com/nasa9084/go-genutils"
)

func TestInterface(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "internal/tests/test_interfaces.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	got, err := gen.LoadInterfaces(f)
	if err != nil {
		t.Error(err)
		return
	}
	expected := []gen.Interface{
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
	}
	assertEqualInterfaceLists(t, got, expected)
}

func assertEqualInterfaceLists(t *testing.T, got []gen.Interface, expected []gen.Interface) {
	if len(got) != len(expected) {
		t.Errorf("length of interfaces invalid: %d != %d", len(got), len(expected))
		return
	}
	for i := range got {
		assertEqualInterfaces(t, got[i], expected[i])
	}
}

func assertEqualInterfaces(t *testing.T, got gen.Interface, expected gen.Interface) {
	if got.DocComment != expected.DocComment {
		t.Errorf("%s != %s", got.DocComment, expected.DocComment)
		return
	}
	if got.Name != expected.Name {
		t.Errorf("%s != %s", got.Name, expected.Name)
		return
	}
	assertEqualInterfaceLists(t, got.Interfaces, expected.Interfaces)
	if len(got.Methods) != len(expected.Methods) {
		t.Errorf("length of methods in %s invalid: %d != %d", got.Name, len(got.Methods), len(expected.Methods))
		return
	}
	for j := range got.Methods {
		gm := got.Methods[j]
		em := expected.Methods[j]
		if gm.Name != em.Name {
			t.Errorf("%s != %s", gm.Name, em.Name)
			return
		}
		if len(gm.Params) != len(em.Params) {
			t.Errorf("length of parameters in %s is invalid: %d != %d", gm.Name, len(gm.Params), len(em.Params))
			return
		}
		for k := range gm.Params {
			if gm.Params[k].Name != em.Params[k].Name {
				t.Errorf("%s != %s", gm.Params[k].Name, em.Params[k].Name)
			}
			if gm.Params[k].Type != em.Params[k].Type {
				t.Errorf("%s != %s", gm.Params[k].Type, em.Params[k].Type)
				return
			}
		}
		if len(gm.Results) != len(em.Results) {
			t.Errorf("length of results in %s is invalid: %d != %d", gm.Name, len(gm.Results), len(em.Results))
			return
		}
		for k := range gm.Results {
			if gm.Results[k].Name != em.Results[k].Name {
				t.Errorf("%s != %s", gm.Results[k].Name, em.Results[k].Name)
				return
			}
			if gm.Results[k].Type != em.Results[k].Type {
				t.Errorf("%s != %s", gm.Results[k].Type, em.Results[k].Type)
				return
			}
		}
	}
}
