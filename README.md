# godash

[![Go Reference](https://pkg.go.dev/badge/github.com/hatena/godash.svg)](https://pkg.go.dev/github.com/hatena/godash)

hatena/godash is a minimal wrapper of https://github.com/samber/lo.

**Though this repository is still maintained, we strongly recommend to use iter instead.**

## Motivation

https://github.com/samber/lo is a nice library, but in some cases teams do not want to have all of the helper functions available. Some helper functions defined in samber/lo are actually already available as part of the official [slices package](https://pkg.go.dev/slices).

hatena/godash is meant to only provide a small portion of samber/lo's functions:
- functions that are not available via Go's standard library
- functions that are not "too much", such as "Conditional helpers" or "Error handling"

## Usage
Install the package to your repo with `go get`.

```sh
go get -u github.com/hatena/godash
```

Sample usage:
```go
package main

import (
	"fmt"

	"github.com/hatena/godash"
)

func main() {
	integers := []int{0, 1, 1, 3, 2, 1, 1, 0}
	fmt.Println(godash.Chunk(integers, 3))
}
```

Check out more examples at [godash package - github.com/hatena/godash - Go Packages](https://pkg.go.dev/github.com/hatena/godash).

### Recommended golangci-lint configuration
Since this package aims to only present a subset of functions from `samber/lo`, it is recommended to add linter settings to deny importing directly from `samber/lo`. (If you want to use `samber/lo` directly, then there is no meaning to use `hatena/godash`). 

Below is a minimal `.golangci.yml` configuration to deny `samber/lo` imports.

```yml
linters:
  enable:
    - depguard

linters-settings:
  depguard:
    rules:
      deny-samber-lo:
        deny:
          - pkg: "github.com/samber/lo"
            desc: Use github.com/hatena/godash instead.
```

Once you have `.golangci.yml`, make sure to run `golangci-lint` via CI. Below is an example for GitHub Actions.

```yml
name: CI

on:
  push:
    branches: main
  pull_request:

permissions:
  contents: read
  pull-requests: read

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version-file: "go.mod"
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.54.2
          only-new-issues: true
```

See [Introduction | golangci-lint](https://golangci-lint.run/) for more details.
