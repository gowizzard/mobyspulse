// Package redirect is to redirect the request to the correct path.
package redirect

import "net/http"

// Metrics is to redirect the request to the /metrics path.
func Metrics(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
}
