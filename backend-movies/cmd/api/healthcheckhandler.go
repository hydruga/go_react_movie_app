package main

import (
	"net/http"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := AppStatus{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}

	err := app.writeJSON(w, 200, data, "HealthCheck")
	if err != nil {
		app.logger.Println("Error encounted with writeJSON", err)
	}
}
