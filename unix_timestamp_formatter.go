// Package unix_timestamp_formatter converts between Unix timestamps in seconds,
// milliseconds, microseconds, and nanoseconds and time.Time values.
//
// The zero value of time.Time represents year 1, and its Unix methods return
// values far outside the range of int64 nanoseconds, so all conversions here
// use int64 and preserve the full range that time.Time can represent.
package unix_timestamp_formatter

import (
	"fmt"
	"time"
)

// FromSeconds returns the time.Time corresponding to a Unix timestamp in seconds.
func FromSeconds(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// FromMilliseconds returns the time.Time corresponding to a Unix timestamp in milliseconds.
func FromMilliseconds(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*1e6)
}

// FromMicroseconds returns the time.Time corresponding to a Unix timestamp in microseconds.
func FromMicroseconds(us int64) time.Time {
	return time.Unix(us/1e6, (us%1e6)*1e3)
}

// FromNanoseconds returns the time.Time corresponding to a Unix timestamp in nanoseconds.
func FromNanoseconds(ns int64) time.Time {
	return time.Unix(0, ns)
}

// ToSeconds returns the Unix timestamp in seconds for t.
func ToSeconds(t time.Time) int64 {
	return t.Unix()
}

// ToMilliseconds returns the Unix timestamp in milliseconds for t.
func ToMilliseconds(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond())/1e6
}

// ToMicroseconds returns the Unix timestamp in microseconds for t.
func ToMicroseconds(t time.Time) int64 {
	return t.Unix()*1e6 + int64(t.Nanosecond())/1e3
}

// ToNanoseconds returns the Unix timestamp in nanoseconds for t.
func ToNanoseconds(t time.Time) int64 {
	return t.Unix()*1e9 + int64(t.Nanosecond())
}

// FormatSeconds formats a Unix timestamp in seconds as an RFC 3339 string in UTC.
func FormatSeconds(sec int64) string {
	return time.Unix(sec, 0).UTC().Format(time.RFC3339)
}

// FormatMilliseconds formats a Unix timestamp in milliseconds as an RFC 3339 string in UTC.
func FormatMilliseconds(ms int64) string {
	return FromMilliseconds(ms).UTC().Format(time.RFC3339Nano)
}

// FormatMicroseconds formats a Unix timestamp in microseconds as an RFC 3339 string in UTC.
func FormatMicroseconds(us int64) string {
	return FromMicroseconds(us).UTC().Format(time.RFC3339Nano)
}

// FormatNanoseconds formats a Unix timestamp in nanoseconds as an RFC 3339 string in UTC.
func FormatNanoseconds(ns int64) string {
	return FromNanoseconds(ns).UTC().Format(time.RFC3339Nano)
}

// ParseSeconds parses an RFC 3339 string and returns the Unix timestamp in seconds.
func ParseSeconds(s string) (int64, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return 0, fmt.Errorf("parse seconds: %w", err)
	}
	return ToSeconds(t), nil
}

// ParseMilliseconds parses an RFC 3339 string and returns the Unix timestamp in milliseconds.
func ParseMilliseconds(s string) (int64, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return 0, fmt.Errorf("parse milliseconds: %w", err)
	}
	return ToMilliseconds(t), nil
}

// ParseMicroseconds parses an RFC 3339 string and returns the Unix timestamp in microseconds.
func ParseMicroseconds(s string) (int64, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return 0, fmt.Errorf("parse microseconds: %w", err)
	}
	return ToMicroseconds(t), nil
}

// ParseNanoseconds parses an RFC 3339 string and returns the Unix timestamp in nanoseconds.
func ParseNanoseconds(s string) (int64, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return 0, fmt.Errorf("parse nanoseconds: %w", err)
	}
	return ToNanoseconds(t), nil
}
