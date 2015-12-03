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
$ hubtrend help
```

Print supported languages.
```shell
$ hubtrend ls
```

Print trend repos.
```shell
hubtrend -l <language> -p <period> [-m] show
```

<br/>
### Example
Specific `golang` `daily`
```shell
$ hubtrend -l go -p d show
```
will print like this.
```
 Language:go Period:daily

 - dgraph-io/dgraph (#1 - Go - Daily)
    Scalable, Distributed, Low Latency Graph Database (Go)

 - davidlazar/vuvuzela (#2 - Go - Daily)
    Private messaging system that hides metadata (Go)

 - influxdb/kapacitor (#3 - Go - Daily)
    Open source framework for processing, monitoring, and alerting on time series data (Go)

////
```

with minimal layout
```shell
$ hubtrend -l go -p d -m show
```
will print like this.
```
 Lang:go Period:daily

  dgraph-io/dgraph (#1 - Go - Daily)
  davidlazar/vuvuzela (#2 - Go - Daily)
  influxdb/kapacitor (#3 - Go - Daily)

///
```

<br/>
## License
MIT
