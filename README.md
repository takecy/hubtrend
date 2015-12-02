# hubtrend
hubtrend is simple command-line tool for [GithubTrends](https://github.com/trending) by golang.

![](https://img.shields.io/badge/golang-1.5.1-blue.svg?style=flat)
[![GoDoc](https://godoc.org/github.com/takecy/hubtrend?status.svg)](https://godoc.org/github.com/takecy/hubtrend)

## Usage
```shell
$ go get github.com/takecy/hubtrend
```

Print usage.
```shell
$ hubtrend
```

Print supported languages.
```shell
$ hubtrend ls
```

Print trend repos.  
example) `golang` `daily`
```shell
$ hubtrend -l go -p d show
```
