package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Struct represents a struct type.
type Struct struct {
	DocComments []string
	Name        string
	Fields      Fields
}

// Field represents a field of struct.
type Field struct {
	Comments []string
	Name     string
	Type     string
	Tags     map[string]string
}

// Fields is a list of Fields.
type Fields []Field

// LoadStructsFromFile load struct types from file.
func LoadStructsFromFile(filename string) ([]Struct, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return LoadStructs(f)
}

// LoadStructs load struct types from *ast.File.
func LoadStructs(f *ast.File) ([]Struct, error) {
	var ret []Struct
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
			if st, ok := typ.Type.(*ast.StructType); ok {
				ret = append(ret, NewStruct(genDecl.Doc.Text(), typ.Name.Name, st))
			}
		}
	}
	return ret, nil
}

// NewStruct creates a new Struct object from *ast.StructType.
func NewStruct(comment string, name string, st *ast.StructType) Struct {
	s := Struct{
		Name: name,
	}
	if comment != "" {
		s.DocComments = strings.FieldsFunc(comment, func(r rune) bool { return r == '\n' })
	}
	for _, af := range st.Fields.List {
		f := Field{}
		if af.Doc != nil {
			f.Comments = strings.FieldsFunc(strings.Trim(af.Doc.Text(), " 	"), func(r rune) bool { return r == '\n' })
		}
		if len(af.Names) != 0 {
			f.Name = af.Names[0].Name
		}
		switch t := af.Type.(type) {
		case *ast.Ident:
			f.Type = t.Name
		case *ast.StructType:
			s := NewStruct("", "", t)
			var buf strings.Builder
			buf.WriteString("struct {")
			buf.WriteString(s.Fields.String())
			buf.WriteString("\n}")
			f.Type = buf.String()
		default:
			panic(reflect.TypeOf(t))
		}
		if af.Tag != nil {
			f.Tags = map[string]string{}
			tags := strings.Fields(strings.Trim(af.Tag.Value, "`"))
			for _, tag := range tags {
				t := strings.Split(tag, ":")
				f.Tags[t[0]] = strings.Trim(t[1], `"`)
			}
		}
		s.Fields = append(s.Fields, f)
	}
	return s
}

// String returns formatted source-form struct type definition.
func (s Struct) String() string {
	var buf bytes.Buffer
	if s.DocComments != nil && len(s.DocComments) != 0 {
		for _, line := range s.DocComments {
			if line != "" {
				fmt.Fprintf(&buf, "// %s\n", line)
			}
		}
	}
	fmt.Fprintf(&buf, "type %s struct {", s.Name)
	fmt.Fprint(&buf, s.Fields.String())
	fmt.Fprint(&buf, "\n}")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(src)
	return buf.String()
}

// String returns a struct field line as string.
func (field Field) String() string {
	var buf strings.Builder
	if field.Comments != nil {
		for _, line := range field.Comments {
			if line != "" {
				fmt.Fprintf(&buf, "// %s\n", line)
			}
		}
	}
	fmt.Fprintf(&buf, "%s %s", field.Name, field.Type)
	var tagKeys []string
	for k := range field.Tags {
		tagKeys = append(tagKeys, k)
	}
	sort.Strings(tagKeys)
	var tags []string
	for _, k := range tagKeys {
		tags = append(tags, fmt.Sprintf("%s: %s", k, strconv.Quote(field.Tags[k])))
	}
	if len(tags) > 0 {
		fmt.Fprintf(&buf, " `%s`", strings.Join(tags, " "))
	}
	return buf.String()

}

// String returns a struct field lines as string.
func (fields Fields) String() string {
	var buf strings.Builder
	for _, f := range fields {
		buf.WriteString("\n" + f.String())
	}
	return buf.String()
}
