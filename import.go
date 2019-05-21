package gen

import (
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
type Imports []Import

// NewImports returns a new import list from import path string slice.
func NewImports(pkgs []string) Imports {
	var imps Imports
	for _, pkg := range pkgs {
		imps = append(imps, Import{ImportPath: pkg})
	}
	return imps
}

// String generates an `import` declaration states.
// If the number of imports is 1, this is same as Import.String().
// Returned string is not formatted. If you need, use format.Source() to format returned code..
func (imps Imports) String() string {
	switch len(imps) {
	case 0:
		return ""
	case 1:
		return imps[0].String()
	default:
		var buf strings.Builder
		buf.Grow(32 * len(imps)) // 32 is nearly equal len of 1 line (avg) from heuristic
		buf.WriteString("\n\nimport (")
		for _, imp := range imps {
			buf.WriteString("\n")
			if imp.PackageName != "" {
				buf.WriteString(imp.PackageName + " ")
			}
			buf.WriteString(`"` + imp.ImportPath + `"`)
		}
		buf.WriteString("\n)")
		return buf.String()
	}
}
