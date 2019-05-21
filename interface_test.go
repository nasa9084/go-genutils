package gen_test

import (
	"go/parser"
	"go/token"
	"strconv"
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
	t.Run("len(Interfaces)", func(t *testing.T) {
		if len(got) != len(expected) {
			t.Errorf("length of interfaces invalid: %d != %d", len(got), len(expected))
			return
		}
	})
	for i := range got {
		t.Run("Interfaces["+strconv.Itoa(i)+"]", func(t *testing.T) {
			assertEqualInterfaces(t, got[i], expected[i])
		})
	}
}

func assertEqualInterfaces(t *testing.T, got gen.Interface, expected gen.Interface) {
	t.Run("DocComment", func(t *testing.T) {
		if got.DocComment != expected.DocComment {
			t.Errorf("%s != %s", got.DocComment, expected.DocComment)
			return
		}
	})
	t.Run("Name", func(t *testing.T) {
		if got.Name != expected.Name {
			t.Errorf("%s != %s", got.Name, expected.Name)
			return
		}
	})
	assertEqualInterfaceLists(t, got.Interfaces, expected.Interfaces)
	t.Run("len(Methods)", func(t *testing.T) {
		if len(got.Methods) != len(expected.Methods) {
			t.Errorf("length of methods in %s invalid: %d != %d", got.Name, len(got.Methods), len(expected.Methods))
			return
		}
	})
	for j := range got.Methods {
		t.Run("Methods["+strconv.Itoa(j)+"]", func(t *testing.T) {
			gm := got.Methods[j]
			em := expected.Methods[j]
			t.Run("Method.Name", func(t *testing.T) {
				if gm.Name != em.Name {
					t.Errorf("%s != %s", gm.Name, em.Name)
					return
				}
			})
			t.Run("len(Params)", func(t *testing.T) {
				if len(gm.Params) != len(em.Params) {
					t.Errorf("length of parameters in %s is invalid: %d != %d", gm.Name, len(gm.Params), len(em.Params))
					return
				}
			})
			for k := range gm.Params {
				t.Run("Params["+strconv.Itoa(k)+"]", func(t *testing.T) {
					t.Run("Param.Name", func(t *testing.T) {
						if gm.Params[k].Name != em.Params[k].Name {
							t.Errorf("%s != %s", gm.Params[k].Name, em.Params[k].Name)
						}
					})
					t.Run("Param.Type", func(t *testing.T) {
						if gm.Params[k].Type != em.Params[k].Type {
							t.Errorf("%s != %s", gm.Params[k].Type, em.Params[k].Type)
							return
						}
					})
				})
			}
			t.Run("len(Results)", func(t *testing.T) {
				if len(gm.Results) != len(em.Results) {
					t.Errorf("length of results in %s is invalid: %d != %d", gm.Name, len(gm.Results), len(em.Results))
					return
				}
			})
			for k := range gm.Results {
				t.Run("Results["+strconv.Itoa(k)+"]", func(t *testing.T) {
					t.Run("Result.Name", func(t *testing.T) {
						if gm.Results[k].Name != em.Results[k].Name {
							t.Errorf("%s != %s", gm.Results[k].Name, em.Results[k].Name)
							return
						}
					})
					t.Run("Result.Type", func(t *testing.T) {
						if gm.Results[k].Type != em.Results[k].Type {
							t.Errorf("%s != %s", gm.Results[k].Type, em.Results[k].Type)
							return
						}
					})
				})
			}
		})
	}
}
