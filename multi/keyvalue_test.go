package multi

import (
	"flag"
	"testing"
)

func TestMultiKeyValuet(t *testing.T) {

	kv := new(KeyValueStringFlag)
	err := kv.Set("foo=bar")

	if err != nil {
		t.Fatalf("Failed to set KeyValueStringFlag, %v", err)
	}

	v := kv.Value().(string)

	if v != "bar" {
		t.Fatalf("Invalid value for KeyValueStringFlag")
	}

	multi_kv := new(KeyValueString)
	err = multi_kv.Set("bar=foo")

	if err != nil {
		t.Fatalf("Failed to set multi KeyValueString, %v", err)
	}

	v2 := multi_kv.Get().(KeyValueString)

	if len(v2) != 1 {
		t.Fatal("Invalid count for multi KeyValueString")
	}

	multi_v := v2[0].Value().(string)

	if multi_v != "foo" {
		t.Fatalf("Invalid value for multi KeyValueString")
	}

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
