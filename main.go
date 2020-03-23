// go-cmd-template is an template repository for simple go commands
// with example makefiles, modules and github actions.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	version = "v0.0.0-dev"

	versionFlag = flag.Bool("v", false, "print version and exit")
)

func main() {
	flag.Usage = printUsage
	os.Exit(realMain(os.Args[1:], os.Stdout))
}

func realMain(args []string, stdout io.Writer) int {
	flag.CommandLine.Parse(args)
	if *versionFlag {
		fmt.Fprintln(stdout, version)
		return 0
	}

	for _, arg := range flag.Args() {
		fmt.Printf("%s: arg %s\n", os.Args[0], arg)
	}
	return 0
}

func usageError(msg string) int {
	fmt.Fprintln(os.Stderr, msg)
	printUsage()
	return 2
}

func fatal(format string, args ...interface{}) int {
	format = os.Args[0] + ": " + format + "\n"
	fmt.Fprintf(os.Stderr, format, args...)
	return 1
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: %s [options] [args]

  TODO: Describe the command

Options:

`, os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
}
