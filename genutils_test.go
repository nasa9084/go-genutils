package gen_test

import (
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
	}
	for _, c := range candidates {
		output := gen.NewImports(c.pkgs).String()
		if output != c.expect {
			t.Errorf("%s != %s", output, c.expect)
			return
		}
	}
}
