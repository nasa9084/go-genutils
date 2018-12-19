package gen

import (
	"fmt"
	"io"
	"strconv"

	errwriter "github.com/nasa9084/go-errwriter"
)

// Imports generates an `import` declaration states from a string slice which
// contains package names you want to import.
// error is returned only if w cannot be written.
func Imports(w io.Writer, pkgs []string) error {
	ew := errwriter.New(w)

	switch len(pkgs) {
	case 0:
		return nil
	case 1:
		fmt.Fprintf(ew, "\n\nimport %s", strconv.Quote(pkgs[0]))
	default:
		fmt.Fprint(ew, "\n\nimport (")
		for _, pkg := range pkgs {
			fmt.Fprintf(ew, "\n%s", strconv.Quote(pkg))
		}
		fmt.Fprint(ew, "\n)")
	}

	if _, err := ew.Write(nil); err != nil {
		return err
	}
	return nil
}
