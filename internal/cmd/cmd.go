package cmd

import (
	"fmt"
	"io"

	"github.com/yoheimuta/protolint/internal/cmd/subcmds/lint"
	"github.com/yoheimuta/protolint/internal/cmd/subcmds/list"
	"github.com/yoheimuta/protolint/internal/osutil"
)

const (
	help = `
Protocol Buffer Linter Command.

Usage:
	protolint <command> [arguments]

The commands are:
	lint     lint protocol buffer files
	list     list all current lint rules being used
	version  print protolint version
`
)

const (
	subCmdLint    = "lint"
	subCmdList    = "list"
	subCmdVersion = "version"
)

var (
	version  = "master"
	revision = "latest"
)

// Do runs the command logic.
func Do(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	switch {
	case len(args) == 0:
		_, _ = fmt.Fprint(stderr, help)
		return osutil.ExitFailure
	default:
		return doSub(
			args,
			stdout,
			stderr,
		)
	}
}

func doSub(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	switch args[0] {
	case subCmdLint:
		return doLint(args[1:], stdout, stderr)
	case subCmdList:
		return doList(stdout, stderr)
	case subCmdVersion:
		return doVersion(stdout)
	default:
		return doLint(args, stdout, stderr)
	}
}

func doLint(
	args []string,
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	if len(args) < 1 {
		_, _ = fmt.Fprintln(stderr, "protolint lint requires at least one argument. See Usage.")
		_, _ = fmt.Fprint(stderr, help)
		return osutil.ExitFailure
	}

	flags, err := lint.NewFlags(args)
	if err != nil {
		_, _ = fmt.Fprint(stderr, err)
		return osutil.ExitFailure
	}
	if len(flags.Args()) < 1 {
		_, _ = fmt.Fprintln(stderr, "protolint lint requires at least one argument. See Usage.")
		_, _ = fmt.Fprint(stderr, help)
		return osutil.ExitFailure
	}

	subCmd, err := lint.NewCmdLint(
		flags,
		stdout,
		stderr,
	)
	if err != nil {
		_, _ = fmt.Fprint(stderr, err)
		return osutil.ExitFailure
	}
	return subCmd.Run()
}

func doList(
	stdout io.Writer,
	stderr io.Writer,
) osutil.ExitCode {
	subCmd := list.NewCmdList(
		stdout,
		stderr,
	)
	return subCmd.Run()
}

func doVersion(
	stdout io.Writer,
) osutil.ExitCode {
	_, _ = fmt.Fprintln(stdout, "protolint version "+version+"("+revision+")")
	return osutil.ExitSuccess
}
