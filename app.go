package joute

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// A must be loaded to start Joute
type App struct {
	Port int
}

func (app *App) Run() error {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", app.Port), nil); err != nil {
		return err
	}

	return nil
}

func LoadApp() (*App, error) {
	return LoadAppWithConfigFrom(WorkingDirectory{})
}

func LoadAppWithConfigFrom[S ConfigFileSource](source S) (*App, error) {

	var app App
	if reader, err := source.Reader(); err == nil {
		err = json.NewDecoder(reader).Decode(&app)
		return &app, err
	} else {
		return nil, err
	}
}
