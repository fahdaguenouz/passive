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
		fn   string
		ip   string
		u    string
		help bool
	)

	fs.StringVar(&fn, "fn", "", "Search with full-name")
	fs.StringVar(&ip, "ip", "", "Search with ip address")
	fs.StringVar(&u, "u", "", "Search with username")

	// support both -h and --help
	fs.BoolVar(&help, "h", false, "Show help")
	fs.BoolVar(&help, "help", false, "Show help")

	if err := fs.Parse(args); err != nil {
		return Options{}, false, err
	}
	if help {
		return Options{}, true, nil
	}

	rest := fs.Args()

	selected := 0
	mode := ModeNone
	query := ""

	// helper: joins flag value + remaining args (important for full name)
	joinValueAndRest := func(val string) string {
		parts := []string{}
		if strings.TrimSpace(val) != "" {
			parts = append(parts, val)
		}
		if len(rest) > 0 {
			parts = append(parts, rest...)
		}
		return strings.TrimSpace(strings.Join(parts, " "))
	}

	if fn != "" {
		selected++
		mode = ModeFullName
		query = joinValueAndRest(fn)
	}
	if ip != "" {
		selected++
		mode = ModeIP
		query = joinValueAndRest(ip)
	}
	if u != "" {
		selected++
		mode = ModeUsername
		query = joinValueAndRest(u)
	}

	// also allow: passive JohnDoe  (no flags) -> show help (for now)
	if selected == 0 {
		return Options{}, true, nil
	}
	if selected > 1 {
		return Options{}, false, errors.New("choose only one option: -fn OR -ip OR -u")
	}

	if strings.TrimSpace(query) == "" {
		return Options{}, false, fmt.Errorf("missing value for selected option")
	}

	return Options{Mode: mode, Query: query}, false, nil
}

func PrintHelp(w io.Writer) {
	fmt.Fprintln(w, "passive v1.0.0")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  passive -fn \"First Last\"")
	fmt.Fprintln(w, "  passive -ip 8.8.8.8")
	fmt.Fprintln(w, "  passive -u username")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -fn       Search with full-name")
	fmt.Fprintln(w, "  -ip       Search with ip address")
	fmt.Fprintln(w, "  -u        Search with username")
	fmt.Fprintln(w, "  -h, --help  Show help")
}
