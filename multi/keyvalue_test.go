package multi

import (
	"flag"
	"testing"
)

func TestMultiKeyValuet(t *testing.T) {
	t.Skip()
}

func TestKeyValueCSVString(t *testing.T) {

	pairs := "hello=world,foo=bar"
	var kv KeyValueCSVString

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Var(&kv, "test", "...")

	fs.Set("test", pairs)

	if len(kv) != 2 {
		t.Fatalf("Unexpected length, %d", len(kv))
	}

	if kv.String() != pairs {
		t.Fatalf("Unexpected string for pairs: %s", kv.String())
	}
}
