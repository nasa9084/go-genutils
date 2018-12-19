package gen_test

import (
	"bytes"
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
		var buf bytes.Buffer
		if err := gen.Imports(&buf, c.pkgs); err != nil {
			t.Fatal(err)
		}
		output := buf.String()
		if output != c.expect {
			t.Errorf("%s != %s", output, c.expect)
			return
		}
	}
}
