package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	oob "github.com/jessemillar/oob/pkg"
)

const version = "1.1.0"

type argConfig struct {
	version  bool
	fileMode bool
	args     []string // args are the positional (non-flag) command-line arguments
}

// parseFlags parses the command-line arguments provided to the program
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'
// Returns the Config in case parsing succeeded, or an error
// In any case, the output of the flag.Parse is returned in output
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message
func parseFlags(progname string, args []string) (config *argConfig, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)

	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf argConfig

	flags.BoolVar(&conf.version, "v", false, "Print the current oob version")
	flags.BoolVar(&conf.fileMode, "f", false, "Read a file instead of STDIN")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}

	conf.args = flags.Args()

	return &conf, buf.String(), nil
}

func main() {
	// Parse command line flags
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:\n", output)
		os.Exit(1)
	}

	// Print the version number and exit if that's what's asked for
	if conf.version {
		fmt.Fprintf(os.Stdout, version+"\n")
		os.Exit(0)
	}

	if conf.fileMode {
		dat, err := ioutil.ReadFile(conf.args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Print(oob.Oob(string(dat)))
	} else {
		fmt.Print(oob.Oob(strings.Join(conf.args, " ")))
	}
}
