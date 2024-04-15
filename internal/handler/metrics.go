package handler

import (
	"fmt"
	"github.com/gowizzard/mobyspulse/internal/request"
	"github.com/gowizzard/mobyspulse/internal/write"
	"net/http"
	"strconv"
	"strings"
)

// Metric is to define the metric struct. It has the name, labels, and value fields.
type Metric struct {
	Name   string
	Labels []Labels
	Value  int
}

// Labels is to define the labels' struct. It has the key and value fields.
type Labels struct {
	Key   string
	Value string
}

// Metrics is to return the metrics for prometheus to scrape. It gets the version information from the docker unix socket.
// It writes a comment to the buffer to let the user know that the system is being monitored. It gets the container information
// from the docker unix socket. It gets the container information from the docker unix socket. It gets the image information from
// the docker unix socket. It writes the container restart count metric to the buffer.
func Metrics(w http.ResponseWriter, _ *http.Request) {

	version, err := request.Version()
	if err != nil {
		write.Logger.Error("Get the version information from docker unix socket.", "err", err)
	}

	_, err = fmt.Fprintln(w, "# Moby's Pulse - Docker metrics exporter for Prometheus.")
	if err != nil {
		write.Logger.Error("Write the comment metric.", "err", err)
	}

	_, err = fmt.Fprintf(w, "# This is a custom exporter for Docker metrics. Your system running Docker version \"%s\" and API version \"%s\" is being monitored.\n", version.Version, version.ApiVersion)
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

		metric := Metric{
			Name: "container_restart_count",
			Labels: []Labels{
				{
					Key:   "id",
					Value: container.Id,
				},
				{
					Key:   "name",
					Value: strings.Replace(container.Name, "/", "", 1),
				},
				{
					Key:   "image",
					Value: tag,
				},
				{
					Key:   "status",
					Value: container.State.Status,
				},
				{
					Key:   "created",
					Value: strconv.FormatInt(container.Created.Unix(), 10),
				},
				{
					Key:   "started_at",
					Value: strconv.FormatInt(container.State.StartedAt.Unix(), 10),
				},
			},
			Value: container.RestartCount,
		}

		var b strings.Builder

		b.WriteString(metric.Name)
		b.WriteString("{")

		for index, value := range metric.Labels {

			_, err = fmt.Fprintf(&b, `%s="%s"`, value.Key, value.Value)
			if err != nil {
				write.Logger.Error("Write the container restart count metric.", "err", err)
			}

			if index < len(metric.Labels)-1 {
				b.WriteString(",")
			}

		}

		b.WriteString("}")
		b.WriteString(fmt.Sprintf(" %v\n", metric.Value))

		_, err = fmt.Fprintf(w, b.String())
		if err != nil {
			write.Logger.Error("Write the container restart count metric.", "err", err)
		}

	}

}
