package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

type Index struct {
	Name        string
	ServiceData [3]Service
}

type Service struct {
	Name string
	Url  string
	Icon string
}

var serviceData = [3]Service{
	{
		Name: "Jellyfin",
		Url:  "jellyfin.local",
		Icon: "https://jellyfin.org/images/logo.svg",
	},
	{
		Name: "Kavita",
		Url:  "kavita.local",
		Icon: "https://raw.githubusercontent.com/Kareadita/Kavita/develop/Logo/kavita.svg",
	},
	{
		Name: "Gitea",
		Url:  "gitea.local",
		Icon: "https://about.gitea.com/gitea-text-dark.svg",
	},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(500)
		slog.Error("Unable to parse template", "Error", err)
		return
	}

	data := Index{Name: "Lyn", ServiceData: serviceData}

	if err := indexTemplate.Execute(w, data); err != nil {
		w.WriteHeader(500)
		slog.Error("Unable to build html", "Error", err)
		return
	}
}
