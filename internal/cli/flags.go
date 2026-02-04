package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

type Mode int

const (
	ModeNone Mode = iota
	ModeFullName
	ModeIP
	ModeUsername
)

type Options struct {
	Mode  Mode
	Query string
}

func ParseArgs(args []string) (Options, bool, error) {
	fs := flag.NewFlagSet("passive", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // we print our own help

	var (
		fn = fs.String("fn", "", "Search with full-name")
		ip = fs.String("ip", "", "Search with ip address")
		u  = fs.String("u", "", "Search with username")
		h  = fs.Bool("help", false, "Show help")
	)

	if err := fs.Parse(args); err != nil {
		// if user typed something like -h (unknown), show our help
		return Options{}, false, err
	}

	// handle "--help" / "-help"
	if *h {
		return Options{}, true, nil
	}

	// allow user to pass query without quotes (we just join remaining args)
	rest := strings.TrimSpace(strings.Join(fs.Args(), " "))

	// Determine selected mode (must be exactly one)
	var mode Mode
	var query string
	selected := 0

	if *fn != "" || rest != "" && hasFlag(args, "-fn") {
		selected++
		mode = ModeFullName
		query = firstNonEmpty(*fn, rest)
	}
	if *ip != "" || rest != "" && hasFlag(args, "-ip") {
		selected++
		mode = ModeIP
		query = firstNonEmpty(*ip, rest)
	}
	if *u != "" || rest != "" && hasFlag(args, "-u") {
		selected++
		mode = ModeUsername
		query = firstNonEmpty(*u, rest)
	}

	if selected == 0 {
		// no mode: show help
		return Options{}, true, nil
	}
	if selected > 1 {
		return Options{}, false, errors.New("choose only one option: -fn OR -ip OR -u")
	}

	query = strings.TrimSpace(query)
	if query == "" {
		return Options{}, false, fmt.Errorf("missing value for selected option")
	}

	return Options{Mode: mode, Query: query}, false, nil
}

func PrintHelp(w io.Writer) {
	fmt.Fprintln(w, "Welcome to passive v1.0.0")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "OPTIONS:")
	fmt.Fprintln(w, "    -fn         Search with full-name")
	fmt.Fprintln(w, "    -ip         Search with ip address")
	fmt.Fprintln(w, "    -u          Search with username")
	fmt.Fprintln(w, "    --help      Show help")
}

func firstNonEmpty(a, b string) string {
	if strings.TrimSpace(a) != "" {
		return a
	}
	return b
}

func hasFlag(args []string, name string) bool {
	for _, a := range args {
		if a == name {
			return true
		}
	}
	return false
}
