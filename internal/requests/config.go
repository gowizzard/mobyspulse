// Package request provides the request to the docker socket.
package request

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"
)

// socket is to define the docker socket.
// scheme is to define the scheme for the request.
// host is to define the host for the request.
// timeout is to define the timeout for the request.
// retries is to define the retries for the request.
const (
	socket  = "/var/run/docker.sock"
	scheme  = "http"
	host    = "localhost"
	timeout = 5 * time.Second
	retries = 3
)

// config is to configure the request.
type config struct {
	Path, Method string
	Context      context.Context
}

// send is to send the request to the docker unix socket.
func (c *config) send() (response *http.Response, err error) {

	result := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   c.Path,
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(context context.Context, network, addr string) (netConn net.Conn, err error) {
				return net.DialTimeout("unix", socket, timeout)
			},
		},
	}

	request, err := http.NewRequestWithContext(c.Context, c.Method, result.String(), nil)
	if err != nil {
		return nil, err
	}

	for i := range retries {

		response, err = client.Do(request)
		if err != nil {
			return nil, err
		}

		switch response.StatusCode {
		case http.StatusTooManyRequests:
			time.Sleep(time.Second * time.Duration(2^i))
			continue
		default:
			return response, nil
		}

	}

	return response, nil

}
