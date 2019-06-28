package gen_test

import (
	"strconv"
	"testing"

	gen "github.com/nasa9084/go-genutils"
)

func TestImports(t *testing.T) {
	candidates := []struct {
		pkgs   []string
		expect string
	}{
		{
			pkgs:   nil,
			expect: "",
		},
		{
			pkgs:   []string{},
			expect: "",
		},
		{
			pkgs:   []string{"fmt"},
			expect: "\n\nimport \"fmt\"",
		},
		{
			pkgs:   []string{"fmt", "net/http"},
			expect: "\n\nimport (\n\"fmt\"\n\"net/http\"\n)",
		},
		{
			pkgs:   []string{"sync", "encoding/json"},
			expect: "\n\nimport (\n\"encoding/json\"\n\"sync\"\n)",
		},
	}
	for i, c := range candidates {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output := gen.NewImports(c.pkgs).String()
			if output != c.expect {
				t.Errorf("%s != %s", output, c.expect)
				return
			}
		})
	}
}

func TestImportsIsSorted(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got: gen.Imports{
				gen.Import{ImportPath: "sync"},
				gen.Import{ImportPath: "encoding/json"},
			}.String(),
			want: "\n\nimport (\n\"encoding/json\"\n\"sync\"\n)",
		},
		{
			got: gen.Imports{
				gen.Import{ImportPath: "log", PackageName: "stdlog"},
				gen.Import{ImportPath: "encoding/json"},
			}.String(),
			want: "\n\nimport (\n\"encoding/json\"\nstdlog \"log\"\n)",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("unexpected:\ngot:  %s\nwant: %s", tt.got, tt.want)
				return
			}
		})
	}
}
