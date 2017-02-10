package app

import (
	"net/http"
	"github.com/go-zoo/bone"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
)

func Setup() (*App, error) {
	flag.Parse()

	// create some configurations
	external := External()
	mux := external.newMux()
	app := NewApp(nil, mux)

	// setup the app's routers
	mux.GetFunc("/healthcheck", app.HealthCheck)

	return app, nil
}

type Mux interface {
	GetFunc(string, http.HandlerFunc) *bone.Route
	http.Handler
}

func NewApp(config *Config, mux Mux) *App {
	return &App {
		config,
		mux,
	}
}

type App struct {
	*Config
	Mux Mux
}

func (app *App) CheckHealth() (*HealthCheck, int) {
	return &HealthCheck{
		Status: OK,
	}, 200
}

func (app *App) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "application/json; charset=utf-8")

	healthCheck, statusCode := app.CheckHealth()
	hcJson, err := json.Marshal(healthCheck)
	if err != nil {
		glog.Errorf("Error processing /healthcheck: %v", err)
		writer.WriteHeader(500)
		writer.Write([]byte(""))
		return
	}

	writer.WriteHeader(statusCode)
	writer.Write(hcJson)
}

func (app *App) Run() error {
	// run stuff forever in here
	return http.ListenAndServe(":9000", app.Mux)
}