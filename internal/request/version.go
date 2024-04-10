package request

import (
	"context"
	"encoding/json"
	"net/http"
)

// VersionResponse is to decode the response from the version endpoint.
type VersionResponse struct {
	Version    string `json:"Version"`
	ApiVersion string `json:"ApiVersion"`
	Os         string `json:"Os"`
	Arch       string `json:"Arch"`
}

// Version is to get the version information from the docker unix socket.
func Version() (decode VersionResponse, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := config{
		Path:    "version",
		Method:  http.MethodGet,
		Context: ctx,
	}

	response, err := c.send()
	if err != nil {
		return VersionResponse{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return VersionResponse{}, err
	}

	return decode, nil

}
