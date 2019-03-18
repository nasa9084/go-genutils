package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

type Interface struct {
	DocComment string
	Name       string
	Methods    Methods
}

type Method struct {
	Name    string
	Params  Params
	Results Results
}

type Methods []Method

type Param struct {
	Name string
	Type string
}
type Params []Param

type Result struct {
	Name string
	Type string
}
type Results []Result

func LoadInterfacesFromFile(filename string) ([]Interface, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return LoadInterfaces(f)
}

func LoadInterfaces(f *ast.File) ([]Interface, error) {
	var ret []Interface
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok { // means the decl is FuncDecl
			continue
		}
		if genDecl.Tok != token.TYPE { // means the decl is not type spec
			continue
		}
		for _, spec := range genDecl.Specs {
			typ, ok := spec.(*ast.TypeSpec)
			// when token.TYPE, spec is always *ast.TypeSpec
			if !ok { // so this block is not executed
				return nil, errors.New("Tok is token.TYPE but spec is not TypeSpec")
			}
			if iface, ok := typ.Type.(*ast.InterfaceType); ok {
				ret = append(ret, NewInterface(genDecl.Doc.Text(), typ.Name.Name, iface))
			}
		}
	}
	return ret, nil
}

func NewInterface(comment, name string, iface *ast.InterfaceType) Interface {
	i := Interface{
		DocComment: comment,
		Name:       name,
	}
	for _, af := range iface.Methods.List {
		m := Method{}
		if len(af.Names) != 0 {
			m.Name = af.Names[0].Name
		}
		if fn, ok := af.Type.(*ast.FuncType); ok {
			for _, p := range fn.Params.List {
				param := Param{}
				if t, ok := p.Type.(*ast.Ident); ok {
					param.Type = t.Name
				}
				for _, name := range p.Names {
					param.Name = name.Name
					m.Params = append(m.Params, param)
				}
				if len(p.Names) == 0 {
					m.Params = append(m.Params, param)
				}
			}
			for _, r := range fn.Results.List {
				result := Result{}
				if len(r.Names) != 0 {
					result.Name = r.Names[0].Name
				}
				if t, ok := r.Type.(*ast.Ident); ok {
					result.Type = t.Name
				}
				m.Results = append(m.Results, result)
			}
		}
		i.Methods = append(i.Methods, m)
	}
	return i
}

func (iface Interface) String() string {
	var buf bytes.Buffer
	if iface.DocComment != "" {
		lines := strings.SplitAfter(iface.DocComment, "\n")
		for _, line := range lines[:len(lines)-1] {
			fmt.Fprintf(&buf, "// %s", line)
		}
	}
	fmt.Fprintf(&buf, "type %s interface {", iface.Name)
	for _, method := range iface.Methods {
		fmt.Fprintf(&buf, "\n%s(%s)", method.Name, method.Params.String())
		switch len(method.Results) {
		case 0:
			continue
		case 1:
			fmt.Fprintf(&buf, " %s", method.Results.String())
		case 2:
			fmt.Fprintf(&buf, " (%s)", method.Results.String())
		}

	}
	fmt.Fprint(&buf, "\n}")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(src)
}

func (param Param) String() string {
	return fmt.Sprintf("%s %s", param.Name, param.Type)
}

func (params Params) String() string {
	var paramList []string
	for i, param := range params {
		if i+1 < len(params) && params[i+1].Type == param.Type {
			paramList = append(paramList, param.Name)
		} else {
			paramList = append(paramList, param.String())
		}
	}
	return strings.Join(paramList, ", ")
}

func (result Result) String() string {
	if result.Name != "" {
		return result.Name + " " + result.Type
	}
	return result.Type
}

func (results Results) String() string {
	var resultList []string
	for _, result := range results {
		resultList = append(resultList, result.String())
	}
	return strings.Join(resultList, ", ")
}
