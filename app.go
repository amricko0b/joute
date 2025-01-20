package joute

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// A must be loaded to start Joute
type App struct {
	Config      *AppConfig
	Downstreams DownstreamMap
	Endpoints   EndpointMap
}

type AppConfig struct {
	Port        int
	Downstreams DownstreamConfigMap
	Endpoints   EndpointConfigMap
}

func (app *App) Run() error {

	for path, endpoint := range app.Endpoints {
		http.HandleFunc(path, endpoint.MakeHandlerFunc(app))
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), nil); err != nil {
		return err
	}

	return nil
}

func LoadApp() (*App, error) {
	return LoadAppWithConfigFrom(WorkingDirectory{})
}

func LoadAppWithConfigFrom[S ConfigFileSource](source S) (*App, error) {

	app := App{}
	if reader, err := source.Reader(); err == nil {
		err = json.NewDecoder(reader).Decode(&app.Config)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	app.Downstreams = make(DownstreamMap, len(app.Config.Downstreams))
	for name, cfg := range app.Config.Downstreams {
		app.Downstreams[name] = &Downstream{cfg}
	}

	app.Endpoints = make(EndpointMap, len(app.Config.Endpoints))
	for name, cfg := range app.Config.Endpoints {
		app.Endpoints[name] = &Endpoint{cfg}
	}

	return &app, nil
}
