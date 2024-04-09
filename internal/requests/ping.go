package request

import (
	"context"
	"net/http"
)

// Ping is to ping the docker socket and return the status code.
func Ping() (status int, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := config{
		Path:    "_ping",
		Method:  http.MethodGet,
		Context: ctx,
	}

	response, err := c.send()
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	return response.StatusCode, nil

}
