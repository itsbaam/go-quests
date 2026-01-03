package selecttimeout

import (
	"bytes"
	"os"
	"testing"
)

func TestFunctionOrdered(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	FunctionOrdered()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	want := "from c1\nfrom c4\nfrom c2\nfrom c5\nfrom c3\n"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}

}
