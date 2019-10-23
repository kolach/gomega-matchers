# Gomega Matchers

A lib where I put my custom matchers for [Gomega](https://onsi.github.io/gomega/).

## BelongsTo

Checks actual value is within expectationn array.

```go
Ω(RandomMonth()).Should(BelongTo("Jan", "Feb", "Mar", "Apr", "May",
  "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"))

```

## CausedBy

This matcher is designed to work with [errors](https://godoc.org/github.com/pkg/errors) library.
Suppose we have sample allocator that throws NoMemoryError if too much bytes is required:

```go
var NoMemoryError = errors.New("No memory")

func doAllocate() error {
  if size > 512 {
    return NoMemoryError
  }
  // do allocate
  return nil
}

func allocate(size int) error {
  if err := doAllocate(); err != nil {
    return errors.Wrapf(err, "failed to allocate %s bytes", size)
  }
}
```

and we need to test the cause of the failure:

```go
It("should fail on allocation bigger than 512", func() {
  err := allocate(1024)
  Ω(err).Should(BeCausedBy(NoMemoryError))
})
```

## PanicWith

Checks the value that panic returns. Here is an example:


```go
It("should panic with foo", func() {
  Ω(func() { panic("foo") }).Should(PanicWith("foo"))
})
```
