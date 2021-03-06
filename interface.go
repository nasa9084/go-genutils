package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

type Interface struct {
	DocComment string
	Name       string
	Interfaces []Interface
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
	if iface == nil {
		return i
	}
	for _, af := range iface.Methods.List {
		m := Method{}
		if len(af.Names) != 0 {
			m.Name = af.Names[0].Name
		} else {
			i.Interfaces = append(i.Interfaces, NewInterface("", af.Type.(*ast.Ident).Name, nil))
			continue
		}
		if fn, ok := af.Type.(*ast.FuncType); ok {
			if fn.Params != nil {
				for _, p := range fn.Params.List {
					param := Param{
						Type: getType(p.Type),
					}
					for _, name := range p.Names {
						param.Name = name.Name
						m.Params = append(m.Params, param)
					}
					if len(p.Names) == 0 {
						m.Params = append(m.Params, param)
					}
				}
			}
			if fn.Results != nil {
				for _, r := range fn.Results.List {
					result := Result{
						Type: getType(r.Type),
					}
					if len(r.Names) != 0 {
						result.Name = r.Names[0].Name
					}
					m.Results = append(m.Results, result)
				}
			}
		}
		i.Methods = append(i.Methods, m)
	}
	return i
}

func getType(expr ast.Expr) string {
	switch t := expr.(type) { // need generarize param and result
	case *ast.Ident: // primitive
		return t.Name
	case *ast.StarExpr: // pointer
		switch typ := t.X.(type) {
		case *ast.Ident: // pointer of primitive
			return "*" + typ.Name
		case *ast.SelectorExpr: // pointer of pkg.type
			return "*" + typ.X.(*ast.Ident).Name + "." + typ.Sel.Name
		}
	case *ast.SelectorExpr: // pkg.type
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.ArrayType: // array
		var len string
		if t.Len != nil {
			len = t.Len.(*ast.BasicLit).Value
		}
		return "[" + len + "]" + getType(t.Elt)
	case *ast.InterfaceType: // interface{}
		return "interface{}"
	case *ast.Ellipsis: // ...type
		return "..." + getType(t.Elt)
	default:
		log.Printf("unknown expr type: %s", reflect.TypeOf(expr))
		os.Exit(1)
	}
	return ""
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
		fmt.Fprint(&buf, method.String())
	}
	fmt.Fprint(&buf, "\n}")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(src)
}

func (method Method) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "\n%s(%s)", method.Name, method.Params.String())
	if len(method.Results) == 1 {
		fmt.Fprintf(&buf, " %s", method.Results.String())
	} else if len(method.Results) > 1 {
		fmt.Fprintf(&buf, " (%s)", method.Results.String())
	}
	return buf.String()
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
