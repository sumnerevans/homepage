package main

import (
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

type HomeTemplateData struct {
}

func (a *Application) Home() http.HandlerFunc {
	template, err := template.ParseFS(TemplateFS, "templates/home.html")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse template")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		templateData := HomeTemplateData{}
		if err := template.ExecuteTemplate(w, "home.html", templateData); err != nil {
			log.Err(err).Msg("Failed to execute the template")
		}
	}
}

func (a *Application) Start(listen string) {
	log.Info().Msg("Starting router")

	http.HandleFunc("/", a.Home())

	log.Info().Str("listen", listen).Msg("Starting server")
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
