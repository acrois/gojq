package cli

import "os"

// Run gojq.
func Run() int {
	return (&Cli{
		inStream:  os.Stdin,
		outStream: os.Stdout,
		errStream: os.Stderr,
	}).RunWithExitCode(os.Args[1:])
}
