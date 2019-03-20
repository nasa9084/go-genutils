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
	ifaces, err := gen.LoadInterfaces(f)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ifaces) != 1 {
		t.Errorf("len(ifaces) is not expected valud: %d != 1", len(ifaces))
		return
	}
	i := ifaces[0]
	if i.Name != "Anything" {
		t.Errorf("%s != Anything", i.Name)
		return
	}
	if len(i.Methods) != 2 {
		t.Errorf("len(iface.Fields) is invalid: %d != 2", len(i.Methods))
		return
	}
	m1 := i.Methods[0]
	if m1.Name != "Foo" {
		t.Errorf("%s != Foo", m1.Name)
		return
	}
	if len(m1.Params) != 2 {
		t.Errorf("%d != 2", len(m1.Params))
		return
	}
	if m1.Params[0].Type != "*string" {
		t.Errorf("%s != *string", m1.Params[0].Type)
		return
	}
	m2 := i.Methods[1]
	if m2.Results[0].Type != "*Bar" {
		t.Errorf("%s != *Bar", m2.Results[0].Type)
		return
	}
	// need more tests
	t.Log(i)
}
