package gen

import (
	"strconv"
	"testing"
)

func TestImportsSplit(t *testing.T) {
	imports := NewImports([]string{
		"net/http",
		"github.com/nasa9084/go-genutils/gen",
		"io/ioutil",
		"github.com/foo/bar",
		"fmt",
		"github.com/nasa9084/go-openapi/openapi",
	})
	imports.BasePath = "github.com/nasa9084/go-genutils"
	got := imports.String()
	want := "\n\nimport (\n\"fmt\"\n\"io/ioutil\"\n\"net/http\"\n\n\"github.com/foo/bar\"\n\"github.com/nasa9084/go-openapi/openapi\"\n\n\"github.com/nasa9084/go-genutils/gen\"\n)"
	if got != want {
		t.Errorf("%s != %s", got, want)
		return
	}
}

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
			output := NewImports(c.pkgs).String()
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
			got: Imports{
				imports: []Import{
					{ImportPath: "sync"},
					{ImportPath: "encoding/json"},
				},
			}.String(),
			want: "\n\nimport (\n\"encoding/json\"\n\"sync\"\n)",
		},
		{
			got: Imports{
				imports: []Import{
					{ImportPath: "log", PackageName: "stdlog"},
					{ImportPath: "encoding/json"},
				},
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
