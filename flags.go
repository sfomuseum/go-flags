package flags

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func Parse(fs *flag.FlagSet) {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-h" {
		fs.Usage()
		os.Exit(0)
	}

	fs.Parse(args)
}

func SetFlagsFromEnvVars(fs *flag.FlagSet, prefix string) error {

	prefix = normalize(prefix)

	fs.VisitAll(func(fl *flag.Flag) {

		name := fl.Name
		env := name

		env = normalize(env)
		env = fmt.Sprintf("%s_%s", prefix, env)

		val, ok := os.LookupEnv(env)

		if ok {
			log.Printf("set -%s flag (%s) from %s environment variable\n", name, val, env)
			fs.Set(name, val)
		}
	})

	return nil
}

func NewFlagSet(name string) *flag.FlagSet {

	fs := flag.NewFlagSet(name, flag.ExitOnError)

	fs.Usage = func() {
		fs.PrintDefaults()
	}

	return fs
}

func normalize(raw string) string {

	new := raw

	new = strings.ToUpper(new)
	new = strings.Replace(new, "-", "_", -1)

	return new
}
