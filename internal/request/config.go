// Package request provides the request to the docker socket.
package request

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"
)

// config is to configure the request.
type config struct {
	Path, Method string
	Context      context.Context
}

// socket is to define the docker socket.
// scheme is to define the scheme for the request.
// host is to define the host for the request.
// timeout is to define the timeout for the request.
const (
	socket  = "/var/run/docker.sock"
	scheme  = "http"
	host    = "localhost"
	timeout = 5 * time.Second
)

// Client is to define the http client.
var (
	Client *http.Client
)

// init is to initialize the http client.
func init() {

	Client = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout("unix", socket, timeout)
			},
		},
	}

}

// send is to send the request to the docker unix socket.
func (c *config) send() (response *http.Response, err error) {

	result := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   c.Path,
	}

	request, err := http.NewRequestWithContext(c.Context, c.Method, result.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
