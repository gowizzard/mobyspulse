package handler

import (
	"github.com/gowizzard/mobyspulse/internal/metric"
	request "github.com/gowizzard/mobyspulse/internal/requests"
	"github.com/gowizzard/mobyspulse/internal/write"
	"net/http"
	"strconv"
	"strings"
)

// Metrics is to return the metrics for prometheus to scrape. It gets the version information from the docker unix socket.
// It writes a comment to the buffer to let the user know that the system is being monitored. It gets the container information
// from the docker unix socket. It gets the container information from the docker unix socket. It gets the image information from
// the docker unix socket. It writes the container restart count metric to the buffer.
func Metrics(w http.ResponseWriter, _ *http.Request) {

	version, err := request.Version()
	if err != nil {
		write.Logger.Error("Get the version information from docker unix socket.", "err", err)
	}

	err = metric.Comment(w, "Moby's Pulse - Docker metrics exporter for Prometheus")
	if err != nil {
		write.Logger.Error("Write the comment metric.", "err", err)
	}

	err = metric.Comment(w, "This is a custom exporter for Docker metrics. Your system running Docker version "+version.Version+" is being monitored.")
	if err != nil {
		write.Logger.Error("Write the comment metric.", "err", err)
	}

	containers, err := request.Containers()
	if err != nil {
		write.Logger.Error("Get the container information from docker unix socket.", "err", err)
	}

	for _, value := range containers {

		container, err := request.Container(value.Id)
		if err != nil {
			write.Logger.Error("Get the container information from docker unix socket.", "err", err)
		}

		image, err := request.Image(container.Image)
		if err != nil {
			write.Logger.Error("Get the image information from docker unix socket.", "err", err)
		}

		var tag string
		if len(image.RepoTags) > 0 {
			tag = image.RepoTags[0]
		}

		err = metric.Counter(w, metric.Config{
			Name: "container_restart_count",
			Labels: map[string]string{
				"id":         container.Id,
				"name":       strings.Replace(container.Name, "/", "", 1),
				"image":      tag,
				"created":    strconv.FormatInt(container.Created.Unix(), 10),
				"started_at": strconv.FormatInt(container.State.StartedAt.Unix(), 10),
			},
			Value: container.RestartCount,
		})
		if err != nil {
			write.Logger.Error("Write the container restart count metric.", "err", err)
		}

	}

}
