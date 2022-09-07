package multi

import (
	"flag"
	"testing"
)

func TestMultiString(t *testing.T) {
	t.Skip()
}

func TestMultiCSVString(t *testing.T) {

	var m MultiCSVString

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Var(&m, "test", "...")

	fs.Set("test", "hello,world")

	if len(m) != 2 {
		t.Fatalf("Unexpected length: %d", len(m))
	}
}
