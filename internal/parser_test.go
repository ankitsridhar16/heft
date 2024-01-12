package internal

import (
	"errors"
	"testing"
)

func TestParseURL(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want error
	}{
		{
			name: "Valid URL",
			in:   "https://www.google.com",
			want: nil,
		},
		{
			name: "Valid URL with port",
			in:   "https://www.google.com:8080",
			want: nil,
		},
		{
			name: "Valid URL with path",
			in:   "https://www.google.com/search",
			want: nil,
		},
		{
			name: "Valid URL with path and query",
			in:   "https://www.google.com/search?q=golang",
			want: nil,
		},
		{
			name: "Valid URL with path, query and fragment",
			in:   "https://www.google.com/search?q=golang#top",
			want: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ParseURL(c.in)
			if !errors.Is(got, c.want) {
				t.Fatalf("want %q but got %q", c.want, got)
			}
		})
	}

}
