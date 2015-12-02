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
)

const usage = `
Usage:
  hubtrend                                  Print usage.
  hubtrend -l <language> -p <period> show   Print Trend repos.
                                            period is [d|w|m]. daily, weekly and monthly.
  hubrrend ls                               Print supported languages.
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

	if flag.Args()[flag.NArg()-1] == "show" {
		err := trender.Rss(*l, *p)
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
