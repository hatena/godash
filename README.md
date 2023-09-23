# godash

[![Go Reference](https://pkg.go.dev/badge/github.com/hatena/godash.svg)](https://pkg.go.dev/github.com/hatena/godash)

hatena/godash is a minimal wrapper of https://github.com/samber/lo.

## Motivation

https://github.com/samber/lo is a nice library, but in some cases teams do not want to have all of the helper functions available. Some helper functions defined in samber/lo are actually already available as part of the official [slices package](https://pkg.go.dev/slices).

hatena/godash is meant to only provide a small portion of samber/lo's functions:
- functions that are not available via Go's standard library
- functions that are not "too much", such as "Conditional helpers" or "Error handling"
