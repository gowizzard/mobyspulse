// Package middleware is to define the different middleware for the http server.
package middleware

import (
	"github.com/gowizzard/mobyspulse/internal/write"
	"net/http"
	"os"
	"reflect"
)

// config is to define the configuration for the basic auth.
type config struct {
	Username string
	Password string
	Ok       bool
}

// environment is to define the configuration for the basic auth.
var (
	environment config
)

// init is to initialize the basic auth configuration. If the environment variables BASIC_AUTH_USERNAME and
// BASIC_AUTH_PASSWORD are set, then the basic auth is enabled. If the environment variables are not set, then the
// basic auth is disabled.
func init() {

	environment.Username, environment.Ok = os.LookupEnv("BASIC_AUTH_USERNAME")
	environment.Password, environment.Ok = os.LookupEnv("BASIC_AUTH_PASSWORD")

	if environment.Ok {
		write.Logger.Info("Basic auth is enabled.")
	}

}

// Authentication is to authenticate the request with basic auth. If the environment variables BASIC_AUTH_USERNAME and
// BASIC_AUTH_PASSWORD are set, then the request will be authenticated with basic auth. If the environment variables are
// not set, then the request will be passed to the next handler.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if environment.Ok {

			username, password, ok := r.BasicAuth()
			if ok {

				if reflect.DeepEqual(environment.Username, username) && reflect.DeepEqual(environment.Password, password) {
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
