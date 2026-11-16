package unix_timestamp_formatter

import (
	"testing"
	"time"
)

func TestFromSeconds(t *testing.T) {
	tests := []struct {
		sec  int64
		want time.Time
	}{
		{0, time.Unix(0, 0)},
		{1, time.Unix(1, 0)},
		{-1, time.Unix(-1, 0)},
		{1609459200, time.Unix(1609459200, 0)},
	}
	for _, tt := range tests {
		got := FromSeconds(tt.sec)
		if !got.Equal(tt.want) {
			t.Errorf("FromSeconds(%d) = %v, want %v", tt.sec, got, tt.want)
		}
	}
}

func TestFromMilliseconds(t *testing.T) {
	tests := []struct {
		ms   int64
		want time.Time
	}{
		{0, time.Unix(0, 0)},
		{1, time.Unix(0, 1e6)},
		{-1, time.Unix(-1, 999e6)},
		{1000, time.Unix(1, 0)},
		{1500, time.Unix(1, 500e6)},
		{-1500, time.Unix(-2, 500e6)},
	}
	for _, tt := range tests {
		got := FromMilliseconds(tt.ms)
		if !got.Equal(tt.want) {
			t.Errorf("FromMilliseconds(%d) = %v, want %v", tt.ms, got, tt.want)
		}
	}
}

func TestFromMicroseconds(t *testing.T) {
	tests := []struct {
		us   int64
		want time.Time
	}{
		{0, time.Unix(0, 0)},
		{1, time.Unix(0, 1e3)},
		{-1, time.Unix(-1, 999999e3)},
		{1000000, time.Unix(1, 0)},
		{1500000, time.Unix(1, 500000e3)},
		{-1500000, time.Unix(-2, 500000e3)},
	}
	for _, tt := range tests {
		got := FromMicroseconds(tt.us)
		if !got.Equal(tt.want) {
			t.Errorf("FromMicroseconds(%d) = %v, want %v", tt.us, got, tt.want)
		}
	}
}

func TestFromNanoseconds(t *testing.T) {
	tests := []struct {
		ns   int64
		want time.Time
	}{
		{0, time.Unix(0, 0)},
		{1, time.Unix(0, 1)},
		{-1, time.Unix(-1, 999999999)},
		{1000000000, time.Unix(1, 0)},
		{1500000000, time.Unix(1, 500000000)},
		{-1500000000, time.Unix(-2, 500000000)},
	}
	for _, tt := range tests {
		got := FromNanoseconds(tt.ns)
		if !got.Equal(tt.want) {
			t.Errorf("FromNanoseconds(%d) = %v, want %v", tt.ns, got, tt.want)
		}
	}
}

func TestToSeconds(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int64
	}{
		{time.Unix(0, 0), 0},
		{time.Unix(1, 0), 1},
		{time.Unix(-1, 0), -1},
		{time.Unix(1609459200, 0), 1609459200},
		{time.Unix(1, 500e6), 1},
	}
	for _, tt := range tests {
		got := ToSeconds(tt.t)
		if got != tt.want {
			t.Errorf("ToSeconds(%v) = %d, want %d", tt.t, got, tt.want)
		}
	}
}

func TestToMilliseconds(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int64
	}{
		{time.Unix(0, 0), 0},
		{time.Unix(0, 1e6), 1},
		{time.Unix(0, 999e6), 999},
		{time.Unix(1, 0), 1000},
		{time.Unix(1, 500e6), 1500},
		{time.Unix(-1, 0), -1000},
		{time.Unix(-2, 500e6), -1500},
	}
	for _, tt := range tests {
		got := ToMilliseconds(tt.t)
		if got != tt.want {
			t.Errorf("ToMilliseconds(%v) = %d, want %d", tt.t, got, tt.want)
		}
	}
}

func TestToMicroseconds(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int64
	}{
		{time.Unix(0, 0), 0},
		{time.Unix(0, 1e3), 1},
		{time.Unix(0, 999999e3), 999999},
		{time.Unix(1, 0), 1000000},
		{time.Unix(1, 500000e3), 1500000},
		{time.Unix(-1, 0), -1000000},
		{time.Unix(-2, 500000e3), -1500000},
	}
	for _, tt := range tests {
		got := ToMicroseconds(tt.t)
		if got != tt.want {
			t.Errorf("ToMicroseconds(%v) = %d, want %d", tt.t, got, tt.want)
		}
	}
}

func TestToNanoseconds(t *testing.T) {
	tests := []struct {
		t    time.Time
		want int64
	}{
		{time.Unix(0, 0), 0},
		{time.Unix(0, 1), 1},
		{time.Unix(0, 999999999), 999999999},
		{time.Unix(1, 0), 1000000000},
		{time.Unix(1, 500000000), 1500000000},
		{time.Unix(-1, 0), -1000000000},
		{time.Unix(-2, 500000000), -1500000000},
	}
	for _, tt := range tests {
		got := ToNanoseconds(tt.t)
		if got != tt.want {
			t.Errorf("ToNanoseconds(%v) = %d, want %d", tt.t, got, tt.want)
		}
	}
}

func TestFormatSeconds(t *testing.T) {
	tests := []struct {
		sec  int64
		want string
	}{
		{0, "1970-01-01T00:00:00Z"},
		{1609459200, "2021-01-01T00:00:00Z"},
		{-1, "1969-12-31T23:59:59Z"},
	}
	for _, tt := range tests {
		got := FormatSeconds(tt.sec)
		if got != tt.want {
			t.Errorf("FormatSeconds(%d) = %q, want %q", tt.sec, got, tt.want)
		}
	}
}

func TestFormatMilliseconds(t *testing.T) {
	tests := []struct {
		ms   int64
		want string
	}{
		{0, "1970-01-01T00:00:00Z"},
		{1, "1970-01-01T00:00:00.001Z"},
		{1500, "1970-01-01T00:00:01.5Z"},
		{-1, "1969-12-31T23:59:59.999Z"},
	}
	for _, tt := range tests {
		got := FormatMilliseconds(tt.ms)
		if got != tt.want {
			t.Errorf("FormatMilliseconds(%d) = %q, want %q", tt.ms, got, tt.want)
		}
	}
}

func TestFormatMicroseconds(t *testing.T) {
	tests := []struct {
		us   int64
		want string
	}{
		{0, "1970-01-01T00:00:00Z"},
		{1, "1970-01-01T00:00:00.000001Z"},
		{1500000, "1970-01-01T00:00:01.5Z"},
		{-1, "1969-12-31T23:59:59.999999Z"},
	}
	for _, tt := range tests {
		got := FormatMicroseconds(tt.us)
		if got != tt.want {
			t.Errorf("FormatMicroseconds(%d) = %q, want %q", tt.us, got, tt.want)
		}
	}
}

func TestFormatNanoseconds(t *testing.T) {
	tests := []struct {
		ns   int64
		want string
	}{
		{0, "1970-01-01T00:00:00Z"},
		{1, "1970-01-01T00:00:00.000000001Z"},
		{1500000000, "1970-01-01T00:00:01.5Z"},
		{-1, "1969-12-31T23:59:59.999999999Z"},
	}
	for _, tt := range tests {
		got := FormatNanoseconds(tt.ns)
		if got != tt.want {
			t.Errorf("FormatNanoseconds(%d) = %q, want %q", tt.ns, got, tt.want)
		}
	}
}

func TestParseSeconds(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"1970-01-01T00:00:00Z", 0},
		{"2021-01-01T00:00:00Z", 1609459200},
		{"1969-12-31T23:59:59Z", -1},
	}
	for _, tt := range tests {
		got, err := ParseSeconds(tt.s)
		if err != nil {
			t.Errorf("ParseSeconds(%q) returned error: %v", tt.s, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseSeconds(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestParseMilliseconds(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"1970-01-01T00:00:00Z", 0},
		{"1970-01-01T00:00:00.001Z", 1},
		{"1970-01-01T00:00:01.5Z", 1500},
		{"1969-12-31T23:59:59.999Z", -1},
	}
	for _, tt := range tests {
		got, err := ParseMilliseconds(tt.s)
		if err != nil {
			t.Errorf("ParseMilliseconds(%q) returned error: %v", tt.s, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseMilliseconds(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestParseMicroseconds(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"1970-01-01T00:00:00Z", 0},
		{"1970-01-01T00:00:00.000001Z", 1},
		{"1970-01-01T00:00:01.5Z", 1500000},
		{"1969-12-31T23:59:59.999999Z", -1},
	}
	for _, tt := range tests {
		got, err := ParseMicroseconds(tt.s)
		if err != nil {
			t.Errorf("ParseMicroseconds(%q) returned error: %v", tt.s, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseMicroseconds(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestParseNanoseconds(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"1970-01-01T00:00:00Z", 0},
		{"1970-01-01T00:00:00.000000001Z", 1},
		{"1970-01-01T00:00:01.5Z", 1500000000},
		{"1969-12-31T23:59:59.999999999Z", -1},
	}
	for _, tt := range tests {
		got, err := ParseNanoseconds(tt.s)
		if err != nil {
			t.Errorf("ParseNanoseconds(%q) returned error: %v", tt.s, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseNanoseconds(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestParseErrors(t *testing.T) {
	if _, err := ParseSeconds("not a time"); err == nil {
		t.Error("ParseSeconds(\"not a time\") expected error, got nil")
	}
	if _, err := ParseMilliseconds("not a time"); err == nil {
		t.Error("ParseMilliseconds(\"not a time\") expected error, got nil")
	}
	if _, err := ParseMicroseconds("not a time"); err == nil {
		t.Error("ParseMicroseconds(\"not a time\") expected error, got nil")
	}
	if _, err := ParseNanoseconds("not a time"); err == nil {
		t.Error("ParseNanoseconds(\"not a time\") expected error, got nil")
	}
}
