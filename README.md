# Berr - Better-Errors

Better-Errors is a simple Golang library for better error handling.

<!-- vim-markdown-toc GFM -->

* [📦 Importing](#-importing)
* [🧩 Usage](#-usage)
    * [Value Example](#value-example)
    * [Error Example](#error-example)

<!-- vim-markdown-toc -->

# 📦 Importing

Run `go get github.com/bbfh-dev/berr.go` in the project directory.

# 🧩 Usage

For functions that return a `(value, error)` use `berr.Get(...)`.

It will return a pointer to a struct that contains methods `ValueOr(...)`, `ValueOrPanic()`, `ValueOrFatal()`.

## Value Example

```go
// In this example divide() returns the result and a possible error:

berr.Get(divide(5, 0)).ValueOr(0)  // Will return 0 if err != nil
berr.Get(divide(10, 5)).ValueOrFatal()  // Will log.Fatal if err != nil
berr.Get(divide(10, 5)).ValueOrPanic()  // Will log.Panic if err != nil
```

For functions that only return a possible error use functions `Try()`, `TryOrPanic()`, `TryOrFatal`.

## Error Example

```go
// In this example divideErrOnly() returns a possible error:

berr.Try(divideErrOnly(5, 0), func(err error) {
    // Handle error, only exists for consistency.
    // Same as if divideErrOnly(5, 0) != nil {}
    // Preferebly only use it to share error handling between different calls.
})

berr.TryOrPanic(divideErrOnly(5, 0))  // Will log.Fatal if err != nil
berr.TryOrFatal(divideErrOnly(5, 0))  // Will log.Panic if err != nil
```
