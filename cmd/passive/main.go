package main

import (
	"fmt"
	"os"

	"passive/internal/cli"
	"passive/internal/core"
	"passive/internal/output"
	"passive/internal/services/fullname"
	"passive/internal/services/ip"
	"passive/internal/services/username"
)

func main() {
	opts, showHelp, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		cli.PrintHelp(os.Stdout)
		os.Exit(2)
	}
	if showHelp {
		cli.PrintHelp(os.Stdout)
		return
	}

	var (
		res core.Result
		runErr error
	)

	switch opts.Mode {
	case cli.ModeFullName:
		res, runErr = fullname.Run(opts.Query)
	case cli.ModeIP:
		res, runErr = ip.Run(opts.Query)
	case cli.ModeUsername:
		res, runErr = username.Run(opts.Query)
	default:
		fmt.Fprintln(os.Stderr, "Error: no mode selected (-fn, -ip, -u)")
		cli.PrintHelp(os.Stdout)
		os.Exit(2)
	}

	// Always print (so errors show up)
	cli.PrintResult(os.Stdout, res)

	// If there was an error, do NOT write to file
	if runErr != nil {
		os.Exit(1)
	}

	filename, err := output.NextResultFilename(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if err := output.WriteResult(filename, res); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Saved in %s\n", filename)
}
