# unusepram

[![pkg.go.dev][gopkg-badge]][gopkg]

`unusepram` finds a unused parameter but its name is not `_`.

```go
package a

func _()                   {}        // OK
func _(_ int)              {}        // OK
func _(_, _ int)           {}        // OK
func _(_, _ int, _ string) {}        // OK
func _(n int)              { _ = n } // OK
func _(_, n int)           { _ = n } // OK
func _(n int)              {}        // want "n is unused parameter"
func _(n, m int)           {}        // want "n is unused parameter" "m is unused parameter"
```

## Install

You can get `unuseparam` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/gostaticanalysis/unuseparam/cmd/unuseparam@latest
```

## How to use

`unuseparam` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which unuseparam) ./...
```
## Analyze with golang.org/x/tools/go/analysis

You can use [unuseparam.Analyzer](https://pkg.go.dev/github.com/gostaticanalysis/unuseparam/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

## Fix unused parameters

`fixunusepram` command check and fix unused parameter to `_`.

```sh
$ go install github.com/gostaticanalysis/unusepram/cmd/fixunusepram@latest
$ fixunusepram ./...
```
<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/unusepram
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/unusepram?status.svg
