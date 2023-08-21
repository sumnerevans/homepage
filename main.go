package main

import (
	"embed"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed templates/*
var TemplateFS embed.FS

func main() {
	logger := log.Output(os.Stdout)
	if os.Getenv("LOG_CONSOLE") != "" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	logger.Info().Msg("backend starting...")

	log.Logger = logger

	app := NewApplication()

	listen := os.Getenv("HOMEPAGE_LISTEN_ADDR")
	app.Start(listen)
}
