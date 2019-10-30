# etcddump

[![Go Report Card](https://goreportcard.com/badge/github.com/jinmukeji/etcddump?style=flat-square)](https://goreportcard.com/report/github.com/jinmukeji/etcddump)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/jinmukeji/etcddump)
[![Releases](https://img.shields.io/github/release/jinmukeji/etcddump/all.svg?style=flat-square)](https://github.com/jinmukeji/etcddump/releases)
[![LICENSE](https://img.shields.io/github/license/jinmukeji/etcddump.svg?style=flat-square)](https://github.com/jinmukeji/etcddump/blob/master/LICENSE)


A tool used to dump / restore etcd KV.

### Install

```sh
go get github.com/jinmukeji/etcddump
```

Or download a compiled version at [Release](https://github.com/jinmukeji/etcddump/releases) page

### Usage

```sh
# help
etcddump -h

# dump
etcddump dump \
	--address=127.0.0.1:2379 \
	--prefix="/micro/config/jm" \
	--output=test.out

# restore
etcddump restore \
	--address=127.0.0.1:2379 \
	--file=test.out
```

