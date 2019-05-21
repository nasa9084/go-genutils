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
