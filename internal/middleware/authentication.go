// Package middleware is to define the different middleware for the http server.
package middleware

import (
	"net/http"
	"reflect"
)

// config is to define the configuration for the basic auth.
type config struct {
	Username string
	Password string
	Ok       bool
}

// Config is to define the configuration for the basic auth.
var (
	Config config
)

// Authentication is to authenticate the request with basic auth. If the environment variables BASIC_AUTH_USERNAME and
// BASIC_AUTH_PASSWORD are set, then the request will be authenticated with basic auth. If the environment variables are
// not set, then the request will be passed to the next handler.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if Config.Ok {

			username, password, ok := r.BasicAuth()
			if ok {

				if reflect.DeepEqual(Config.Username, username) && reflect.DeepEqual(Config.Password, password) {
					next.ServeHTTP(w, r)
					return
				}

			}

			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return

		}

		next.ServeHTTP(w, r)

	})
}
