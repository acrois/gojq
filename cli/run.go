package cli

import "os"

// Run gojq.
func Run() int {
	return (&Cli{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}).RunWithExitCode(os.Args[1:])
}
