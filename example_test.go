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

func ExampleLoadStructs() {
	structs, _ := gen.LoadStructsFromFile("./internal/tests/test_structs.go")
	for _, s := range structs {
		fmt.Println(s.Name)
	}
	// Output:
	// Something
	// Nested
	// Parent
	// Child
}
