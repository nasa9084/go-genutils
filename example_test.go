package gen_test

import (
	"os"

	gen "github.com/nasa9084/go-genutils"
)

func ExampleImportsSingleImport() {
	pkgs := []string{"fmt"}
	if err := gen.Imports(os.Stdout, pkgs); err != nil {
		// some error handling
	}
	// Output:
	//
	// import "fmt"
}

func ExampleImports() {
	pkgs := []string{
		"fmt",
		"net/http",
	}
	if err := gen.Imports(os.Stdout, pkgs); err != nil {
		// some error handling
	}
	// Output:
	//
	// import (
	// "fmt"
	// "net/http"
	// )
}
