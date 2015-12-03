package trender

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
)

// http://github-trends.ryotarai.info/
var (
	rssURLFormat = "http://github-trends.ryotarai.info/rss/github_trends_%s_%s.rss"
	periodM      = map[string]string{
		"d": "daily",
		"w": "weekly",
		"m": "monthly",
	}
)

var (
	lang   string
	period string
	min    bool
)

var helpers = template.FuncMap{
	"magenta": color.MagentaString,
	"yellow":  color.YellowString,
	"green":   color.GreenString,
	"black":   color.BlackString,
	"white":   color.WhiteString,
	"blue":    color.BlueString,
	"cyan":    color.CyanString,
	"red":     color.RedString,
}

const itemTmpl = `
{{range .}}
 - {{.Title | cyan}}
    {{.Description}}
{{end}}
`

const itemMinTmpl = `
{{range .}}
  {{.Title | cyan}} {{end}}
`

// NewRss is init Rss
func NewRss(l, p string, m bool) error {
	_p, ok := periodM[p]
	if !ok {
		return errors.New("bad.period -> [d]or[w]or[m]")
	}

	lang = l
	period = _p
	min = m

	return nil
}

// Rss is fetch RSS
func Rss() (err error) {
	fmt.Fprintf(os.Stdout, "\n Lang:%s Period:%s", lang, period)
	err = fetchFeed(fmt.Sprintf(rssURLFormat, lang, period), 5, nil)
	return
}

func fetchFeed(uri string, timeout int, cr xmlx.CharsetFunc) (err error) {
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	err = feed.Fetch(uri, cr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
		return
	}
	return
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for i := range newitems {
		newitems[i].Description = strings.Replace(newitems[i].Description, "\n", " ", -1)
	}

	err := t().Execute(os.Stdout, newitems)
	if err != nil {
		panic(err)
	}
}

func t() *template.Template {
	tm := itemTmpl
	if min {
		tm = itemMinTmpl
	}
	return template.Must(template.New("item").Funcs(helpers).Parse(tm))
}
