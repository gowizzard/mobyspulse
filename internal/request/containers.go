package request

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"time"
)

// ContainersResponse  is to decode the response from the containers endpoint.
type ContainersResponse struct {
	Id string `json:"Id"`
}

// ContainerResponse  is to decode the response from the container endpoint.
type ContainerResponse struct {
	Created      time.Time `json:"Created"`
	Id           string    `json:"Id"`
	Image        string    `json:"Image"`
	Name         string    `json:"Name"`
	RestartCount int       `json:"RestartCount"`
	State        struct {
		StartedAt time.Time `json:"StartedAt"`
		Status    string    `json:"Status"`
	} `json:"State"`
}

// Containers are to get the containers information from the docker unix socket.
func Containers() (decode []ContainersResponse, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := config{
		Path:    path.Join("containers", "json"),
		Method:  http.MethodGet,
		Context: ctx,
	}

	response, err := c.send()
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	return decode, nil

}

// Container is to get the container information from the docker unix socket.
func Container(id string) (decode ContainerResponse, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := config{
		Path:    path.Join("containers", id, "json"),
		Method:  http.MethodGet,
		Context: ctx,
	}

	response, err := c.send()
	if err != nil {
		return ContainerResponse{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ContainerResponse{}, err
	}

	return decode, nil

}
