package r2

import (
	"net/http"
	"time"
)

// WithTimeout sets the client timeout.
func WithTimeout(d time.Duration) Option {
	return func(r *Request) {
		if r.Client == nil {
			r.Client = &http.Client{}
		}
		r.Client.Timeout = d
	}
}