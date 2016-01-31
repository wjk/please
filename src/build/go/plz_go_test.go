// plz_go_test implements a code templater for go tests.
// This is essentially equivalent to what 'go test' does but it lifts some restrictions
// on file organisation and allows the code to be instrumented for coverage as separate
// build targets rather than having to repeat it for every test.
package main

import (
	"os"
	
	"github.com/op/go-logging"

	"buildgo"
	"output"
)

var log = logging.MustGetLogger("plz_go_test")

var opts struct {
	Dir        string   `short:"d" long:"dir" description:"Directory to search for Go package files for coverage"`
	Verbosity  int      `short:"v" long:"verbose" default:"1" description:"Verbosity of output (higher number = more output, default 1 -> warnings and errors only)"`
	Exclude    []string `short:"x" long:"exclude" default:"third_party/go" description:"Directories to exclude from search"`
	Sources    []string `short:"s" long:"source" description:"Sources to parse for tests" required:"true"`
	Output     string   `short:"o" long:"output" description:"Output filename" required:"true"`
}

func main() {
	output.ParseFlagsOrDie("plz_go_test", &opts)
	output.InitLogging(opts.Verbosity, "", 0)
	coverVars, err := buildgo.FindCoverVars(opts.Dir, opts.Exclude)
	if err != nil {
		log.Fatalf("Error scanning for coverage: %s", err)
	}
	if err = buildgo.WriteTestMain(opts.Sources, opts.Output, coverVars); err != nil {
		log.Fatalf("Error writing test main: %s", err)
	}
	os.Exit(0)
}
