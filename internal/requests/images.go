package request

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
)

// ImageResponse is to decode the response from the image endpoint.
type ImageResponse struct {
	RepoTags []string `json:"RepoTags"`
}

// Image is to get the image from the docker unix socket.
func Image(id string) (decode ImageResponse, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := config{
		Path:    path.Join("images", id, "json"),
		Method:  http.MethodGet,
		Context: ctx,
	}

	response, err := c.send()
	if err != nil {
		return ImageResponse{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ImageResponse{}, err
	}

	return decode, nil

}
