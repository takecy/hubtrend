package trender

import (
	"errors"
	"fmt"
	"os"
	"text/template"

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

const itemTmpl = `
TrendRepos:
--------------------------------------------------------------------------------------------
{{range .}}
--------------------------------------------------------------------------------------------
|  {{.Title}}
|  {{.Description}}
{{end}}
--------------------------------------------------------------------------------------------
`

// Rss is fetch RSS
func Rss(lang, period string) (err error) {
	p, ok := periodM[period]
	if !ok {
		return errors.New("bad.period -> [d]or[w]or[m]")
	}
	err = fetchFeed(fmt.Sprintf(rssURLFormat, lang, p), 5, nil)
	return
}

func fetchFeed(uri string, timeout int, cr xmlx.CharsetFunc) (err error) {
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	err = feed.Fetch(uri, cr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s\n", uri, err)
		return
	}
	return
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Fprintf(os.Stdout, "%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	//	fmt.Fprintf(os.Stdout, "%d new item(s) in %s\n", len(newitems), feed.Url)

	t := template.New("item")
	template.Must(t.Parse(itemTmpl))
	err := t.Execute(os.Stdout, newitems)
	if err != nil {
		panic(err)
	}
}
