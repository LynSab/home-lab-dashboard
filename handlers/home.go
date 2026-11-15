package handlers

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const labelPrefix = "home-lab-dashboard"

type Index struct {
	Name        string
	ServiceData []Service
}

type Service struct {
	Name        string
	Description string
	Url         string
	Icon        string
}

// var serviceData = [3]Service{
// 	{
// 		Name: "Jellyfin",
// 		Url:  "jellyfin.local",
// 		Icon: "https://jellyfin.org/images/logo.svg",
// 	},
// 	{
// 		Name: "Kavita",
// 		Url:  "kavita.local",
// 		Icon: "https://raw.githubusercontent.com/Kareadita/Kavita/develop/Logo/kavita.svg",
// 	},
// 	{
// 		Name: "Gitea",
// 		Url:  "gitea.local",
// 		Icon: "https://about.gitea.com/gitea-text-dark.svg",
// 	},
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	services, err := GetDockerServices()
	if err != nil {
		w.WriteHeader(500)
		slog.Error("Unable to fetch docker service details", "Error", err)
		return
	}

	indexTemplate, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(500)
		slog.Error("Unable to parse template", "Error", err)
		return
	}

	data := Index{Name: "Lyn", ServiceData: services}

	if err := indexTemplate.Execute(w, data); err != nil {
		w.WriteHeader(500)
		slog.Error("Unable to build html", "Error", err)
		return
	}
}

func GetDockerServices() (services []Service, err error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: false, Filters: filters.NewArgs(filters.KeyValuePair{Key: "label", Value: labelPrefix + ".name"})})
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		service := Service{
			Name:        container.Labels[labelPrefix+".name"],
			Description: container.Labels[labelPrefix+".desciption"],
			Url:         container.Labels[labelPrefix+".url"],
			Icon:        container.Labels[labelPrefix+".icon"],
		}

		services = append(services, service)
	}

	return services, nil
}
