package gen_test

import (
	"fmt"

	gen "github.com/nasa9084/go-genutils"
)

func ExampleImports_singleImport() {
	pkgs := []string{"fmt"}
	imports := gen.NewImports(pkgs)
	fmt.Print(imports.String())
	// Output:
	//
	// import "fmt"
}

func ExampleImports() {
	pkgs := []string{
		"fmt",
		"net/http",
	}
	imports := gen.NewImports(pkgs)
	fmt.Print(imports.String())
	// Output:
	//
	// import (
	// "fmt"
	// "net/http"
	// )
}
