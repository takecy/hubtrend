package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/takecy/hubtrend/trender"
)

var (
	l = flag.String("l", "go", "")
	p = flag.String("p", "d", "")
	m = flag.Bool("m", false, "")
)

const usage = `
hubtrend is Simple command-line tool for GithubTrend.
GithugTrend: https://github.com/trending
more info:   https://github.com/takecy/hubtrend#readme

Usage:
  hubtrend [options] <command>

Commands:
  help    Print usage.
  show    Print Trend repos.
  ls      Print supported languages.

Options:
  show:
    -l   Specific language.
         Supported language are by [hubtrend ls]
         Default is [go]
    -p   Specific period. daily->d, weekly->w, monthly->m
         Default is [d].
    -m   Print result with minimal layout.
`

const lsTmpl = `
Languages:{{range .}}
  - {{.}}{{end}}

`

func main() {
	flag.Usage = func() { usageAndExit() }
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if flag.Args()[flag.NArg()-1] == "help" {
		flag.Usage()
		return
	}

	if flag.Args()[flag.NArg()-1] == "show" {
		err := trender.NewRss(*l, *p, *m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		err = trender.Rss()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		return
	}

	if flag.Args()[flag.NArg()-1] == "ls" {
		showLs()
		return
	}

	fmt.Fprintf(os.Stderr, "Error: Bad Command\n")
	flag.Usage()
}

func showLs() {
	t := template.New("ls")
	template.Must(t.Parse(lsTmpl))
	err := t.Execute(os.Stdout, trender.Langs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
}

func usageAndExit() {
	fmt.Fprintf(os.Stderr, "%s\n", usage)
	os.Exit(1)
}
