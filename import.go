//go:generate go run mkstdlib.go

package gen

import (
	"sort"
	"strings"
)

// Import represents an `import` declaration in go struct.
type Import struct {
	PackageName string
	ImportPath  string
}

// String returns an `import` declaration state.
func (imp Import) String() string {
	var buf strings.Builder
	buf.Grow(len(imp.PackageName) + len(imp.ImportPath) + 13)
	buf.WriteString("\n\nimport ")
	if imp.PackageName != "" {
		buf.WriteString(imp.PackageName + " ")
	}
	buf.WriteString(`"` + imp.ImportPath + `"`)
	return buf.String()
}

// Imports represents a list of imports.
type Imports struct {
	BasePath string
	imports  []Import
}

// NewImports returns a new import list from import path string slice.
func NewImports(pkgs []string) Imports {
	var imps Imports
	for _, pkg := range pkgs {
		imps.imports = append(imps.imports, Import{ImportPath: pkg})
	}
	return imps
}

func (imps Imports) Import(pkgName, importPath string) {
	imps.imports = append(imps.imports, Import{
		PackageName: pkgName,
		ImportPath:  importPath,
	})
}

// String generates an `import` declaration states.
// If the number of imports is 1, this is same as Import.String().
// Returned string is not formatted. If you need, use format.Source() to format returned code..
func (imps Imports) String() string {
	switch len(imps.imports) {
	case 0:
		return ""
	case 1:
		return imps.imports[0].String()
	default:
		importSets := [3][]Import{}
		for _, imp := range imps.imports {
			if _, ok := stdlib[imp.ImportPath]; ok {
				importSets[0] = append(importSets[0], imp)
				continue
			}
			if imps.BasePath != "" && strings.HasPrefix(imp.ImportPath, imps.BasePath) {
				importSets[2] = append(importSets[2], imp)
				continue
			}
			importSets[1] = append(importSets[1], imp)
		}

		var buf strings.Builder

		var blocks []string
		for _, importSet := range importSets {
			buf.Reset()
			buf.Grow(32 * len(imps.imports)) // 32 is nearly equal len of 1 line (avg) from heuristic

			sort.Slice(importSet, func(i, j int) bool {
				a := importSet[i].ImportPath
				if importSet[i].PackageName != "" {
					a = importSet[i].PackageName
				}
				b := importSet[j].ImportPath
				if importSet[j].PackageName != "" {
					b = importSet[j].PackageName
				}
				return a < b
			})
			for _, imp := range importSet {
				buf.WriteString("\n")
				if imp.PackageName != "" {
					buf.WriteString(imp.PackageName + " ")
				}
				buf.WriteString(`"` + imp.ImportPath + `"`)
			}
			block := buf.String()
			if block != "" {
				blocks = append(blocks, block)
			}
		}
		buf.Reset()
		buf.WriteString("\n\nimport (")
		buf.WriteString(strings.Join(blocks, "\n"))
		buf.WriteString("\n)")
		return buf.String()
	}
}
