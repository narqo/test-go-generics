# Go 1.18 generics use cases

## What are generics?

See [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

## How to run the examples?

As of today, [`gotip`][1] is the simplest way to run the examples in the repo.

After installing `gotip`, you can run the examples as usual Go code:

```
$ gotip test -v ./...
=== RUN   ExampleEqual
--- PASS: ExampleEqual (0.00s)
PASS
ok  	github.com/narqo/test-go-generics/assert	0.123s
=== RUN   ExampleIterator_sliceIterator
--- PASS: ExampleIterator_sliceIterator (0.00s)
=== RUN   ExampleIterator_csvIterator
--- PASS: ExampleIterator_csvIterator (0.00s)
PASS
ok  	github.com/narqo/test-go-generics/iter	0.197s
```

## License

MIT

[1]: https://godoc.org/golang.org/dl/gotip
