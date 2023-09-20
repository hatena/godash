# go-lo

hatena/go-lo is a minimal wrapper of https://github.com/samber/lo.

## Motivation

https://github.com/samber/lo is a nice library, but in some cases teams do not want to have all of the helper functions available. Some helper functions defined in samber/lo are actually already available as part of the official [slices package](https://pkg.go.dev/slices).

hatena/go-lo is meant to only provide a small portion of samber/lo's functions:
- functions that are not available via Go's standard library
- functions that are not "too much", such as "Conditional helpers" or "Error handling"
