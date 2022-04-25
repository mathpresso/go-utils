# errorutil

This module aims to provide an error wrapper that can be commonly used in mathpresso go projects.

---

## Installation

```bash
$ go env -w GOPRIVATE=github.com/mathpresso/*
$ go get https://github.com/mathpresso/go-utils/errorutil
```

## Usage

```go
// If you want to just wrap error with stack trace, simply wrap your error with .Wrap()
return errorutil.Wrap(err)

// If you want to set some cause-error for your error, simply use `.FromCause()` option
if err != nil {
	return errorutil.Wrap(ErrSomeStaticStuff, errorutil.FromCause(err))
}
```

## API

- `errorutil.Wrap(err error, opts ...wrapOpt) error`
  - Wrap wraps the error with provided opts.
- `errorutil.AutoStackTrace() wrapOpt`
  - AutoStackTrace automatically bind caller's stacktrace to error. This makes some error-capturing module (like [sentry-go](https://github.com/getsentry/sentry-go)) can extract proper stacktrace of your error.
  - For convenience, this option is enabled by default even if you don't include it.
- `errorutil.FromCause(err error) wrapOpt`
  - FromCause wrap the error with provided cause. If you Unwrap this error, provided cause will be extracted.