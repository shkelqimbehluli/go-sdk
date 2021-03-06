package r2

import (
	"net/http"
)

// OptHeaders sets the request headers.
func OptHeaders(headers http.Header) Option {
	return func(r *Request) error {
		r.Header = headers
		return nil
	}
}

// OptHeaderValue adds or sets a header value.
func OptHeaderValue(key, value string) Option {
	return func(r *Request) error {
		if r.Header == nil {
			r.Header = http.Header{}
		}
		r.Header.Set(key, value)
		return nil
	}
}
