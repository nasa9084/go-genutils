package gen_test

import (
	"go/parser"
	"go/token"
	"testing"

	gen "github.com/nasa9084/go-genutils"
)

func TestLoadStructs(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "internal/tests/test_structs.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	ss, err := gen.LoadStructs(f)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ss) != 1 {
		t.Errorf("len(ss) is not expected value: %d != 1", len(ss))
		return
	}
	s := ss[0]
	if s.Name != "Something" {
		t.Errorf("%s != Something", s.Name)
		return
	}
	if len(s.Fields) != 3 {
		t.Errorf("len(s.Fields) is invalid: %d != 3", len(s.Fields))
		return
	}
	if f := s.Fields[0]; f.Name != "Foo" {
		t.Errorf("%s != Foo", f.Name)
		return
	}
	if f := s.Fields[0]; f.Type != "string" {
		t.Errorf("%s != string", f.Type)
		return
	}
	t.Log(s)
}
