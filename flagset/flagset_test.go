package flagset

import (
	"flag"
	"os"
	"testing"
)

func TestFlagSet(t *testing.T) {
	t.Skip()
}

func TestFlagNameToEnvVar(t *testing.T) {

	prefix := "debug"

	tests := map[string]string{
		"enable":      "DEBUG_ENABLE",
		"hello-world": "DEBUG_HELLO_WORLD",
	}

	for name, expected := range tests {

		env_var := FlagNameToEnvVar(prefix, name)

		if env_var != expected {
			t.Fatalf("Unexpected env var for '%s'. Expected '%s' but got '%s'", name, expected, env_var)
		}
	}
}

func TestSetFlagsFromEnvVarsWithFeedback(t *testing.T) {

	err := os.Setenv("DEBUG_HELLO", "world")

	if err != nil {
		t.Fatalf("Failed to set environment variables, %v", err)
	}

	fs := NewFlagSet("test")
	fs.String("hello", "", "An example flag")

	err = SetFlagsFromEnvVars(fs, "DEBUG")

	if err != nil {
		t.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	v := fs.Lookup("hello")

	if v == nil {
		t.Fatalf("Failed to locate -hello flag in flagset")
	}

	// Go is weird...
	i := v.Value.(flag.Getter).Get()

	if i.(string) != "world" {
		t.Fatalf("Unexpected -hello flag: %s", i.(string))
	}

}
