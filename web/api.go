package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sem-onyalo/wimbyai-api/service"
	"github.com/sem-onyalo/wimbyai-api/service/request"
)

// TODO: move to config service
const (
	webAPIVersion  = "0.1"
	defaultAPIPort = 5001
)

// API is a service to the api application
type API struct {
	Config service.Config
	Port   int
}

// apiInfo represents the web api info
type apiInfo struct {
	Name    string
	Version string
}

// NewAPI returns a reference to the web api service
func NewAPI(config service.Config) *API {
	return &API{Config: config, Port: defaultAPIPort}
}

// Start runs the api
func (a API) Start(request.StartAPI) {
	http.HandleFunc("/", a.rootHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s", webAPIVersion), a.apiHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", a.Port), nil)
}

// rootHandler is the http handler for the root path
func (a API) rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("/api/v%s", webAPIVersion), http.StatusFound)
}

// apiHandler is the http handler for the api
func (a API) apiHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(apiInfo{Name: "WimbyAI API", Version: webAPIVersion})
}
