package main

import (
	"fmt"
	"os"
	"time"

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

	start := time.Now()

	var res core.Result
	switch opts.Mode {
	case cli.ModeFullName:
		r, e := fullname.Run(opts.Query)
		if e != nil {
			fmt.Fprintln(os.Stderr, "Error:", e)
			os.Exit(1)
		}
		res = r

	case cli.ModeIP:
		r, e := ip.Run(opts.Query)
		if e != nil {
			fmt.Fprintln(os.Stderr, "Error:", e)
			os.Exit(1)
		}
		res = r

	case cli.ModeUsername:
		r, e := username.Run(opts.Query)
		if e != nil {
			fmt.Fprintln(os.Stderr, "Error:", e)
			os.Exit(1)
		}
		res = r

	default:
		fmt.Fprintln(os.Stderr, "Error: no mode selected (-fn, -ip, -u)")
		cli.PrintHelp(os.Stdout)
		os.Exit(2)
	}

	cli.PrintResult(os.Stdout, res)

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
	_ = start // keep for future: timings, verbose mode, etc.
}
