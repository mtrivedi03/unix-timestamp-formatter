# unix-timestamp-formatter

A small Go library for converting between Unix timestamps in seconds, milliseconds, microseconds, and nanoseconds and `time.Time` values, with helpers for formatting and parsing RFC 3339 strings.

## Usage

```go
package main

import (
	"fmt"
	"time"

	utf "unix-timestamp-formatter"
)

func main() {
	t := utf.FromMilliseconds(1609459200123)
	fmt.Println(t.UTC()) // 2021-01-01 00:00:00.123 +0000 UTC

	ms := utf.ToMilliseconds(t)
	fmt.Println(ms) // 1609459200123

	s := utf.FormatNanoseconds(1500000000)
	fmt.Println(s) // 1970-01-01T00:00:01.5Z

	ns, err := utf.ParseNanoseconds(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(ns) // 1500000000

	_ = time.Time{}
}
```

## Why this library exists

Go's `time.Time` uses a pair of `(seconds, nanoseconds)` and its `Unix`/`UnixNano` methods are not enough to move between the four common Unix timestamp precisions without leaking negative-number bugs. For example, converting a negative timestamp in milliseconds to `time.Time` with `time.Unix(ms/1000, (ms%1000)*1e6)` is correct only if the remainder is handled as a signed quantity; naive implementations use `math.Mod` or floor division and end up one second off. This library centralises that arithmetic and provides symmetric formatting and parsing.

The trade-off: it is a thin wrapper, not a general-purpose date library. It does not handle time zones other than UTC for formatting and parsing, and it does not attempt to interpret timestamps outside the range of `int64` nanoseconds. Values that overflow `int64` nanoseconds are silently wrapped by Go's arithmetic, which is the same behaviour as `time.Time.UnixNano`.

## The awkward edge

Negative timestamps with sub-second components are the common trap. `FromMilliseconds(-1)` must be `1969-12-31T23:59:59.999Z`, not `1969-12-31T23:59:59.001Z` and not `1970-01-01T00:00:00.001Z`. The implementation keeps the remainder signed and lets `time.Unix` normalise it, which produces the expected result. The same applies to microseconds and nanoseconds.

## API

- `FromSeconds(sec int64) time.Time`
- `FromMilliseconds(ms int64) time.Time`
- `FromMicroseconds(us int64) time.Time`
- `FromNanoseconds(ns int64) time.Time`
- `ToSeconds(t time.Time) int64`
- `ToMilliseconds(t time.Time) int64`
- `ToMicroseconds(t time.Time) int64`
- `ToNanoseconds(t time.Time) int64`
- `FormatSeconds(sec int64) string`
- `FormatMilliseconds(ms int64) string`
- `FormatMicroseconds(us int64) string`
- `FormatNanoseconds(ns int64) string`
- `ParseSeconds(s string) (int64, error)`
- `ParseMilliseconds(s string) (int64, error)`
- `ParseMicroseconds(s string) (int64, error)`
- `ParseNanoseconds(s string) (int64, error)`
