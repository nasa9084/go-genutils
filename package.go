package gen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Package struct {
	Name       string
	Structs    []Struct
	Interfaces []Interface
}

func LoadPackagesFromPath(path string) ([]Package, error) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return LoadPackages(pkgs)
}

func LoadPackages(pkgs map[string]*ast.Package) ([]Package, error) {
	ret := make([]Package, len(pkgs))
	var i int
	for _, astpkg := range pkgs {
		pkg, err := LoadPackage(astpkg)
		if err != nil {
			return nil, err
		}
		ret[i] = pkg
		i++
	}
	return ret, nil
}

func LoadPackage(astpkg *ast.Package) (Package, error) {
	pkg := Package{
		Name: astpkg.Name,
	}
	for _, file := range astpkg.Files {
		structs, err := LoadStructs(file)
		if err != nil {
			return Package{}, err
		}
		pkg.Structs = append(pkg.Structs, structs...)
		interfaces, err := LoadInterfaces(file)
		if err != nil {
			return Package{}, err
		}
		pkg.Interfaces = append(pkg.Interfaces, interfaces...)
	}
	return pkg, nil
}
