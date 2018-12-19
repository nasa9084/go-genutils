go-genutils
===========
[![GoDoc](https://godoc.org/github.com/nasa9084/go-genutils?status.svg)](https://godoc.org/github.com/nasa9084/go-genutils)
[![Build Status](https://travis-ci.org/nasa9084/go-genutils.svg?branch=master)](https://travis-ci.org/nasa9084/go-genutils)

Package `gen` provides some utility functions for go code generation.


## SYNOPSIS

``` go
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
```
