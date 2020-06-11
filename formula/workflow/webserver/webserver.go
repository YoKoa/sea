package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/YoKoa/sea/formula/workflow/constant"
	"github.com/YoKoa/sea/formula/workflow/telemetry"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// WebServer handles the webserver configuration
type WebServer struct {
	router        *mux.Router
}

// NewWebserver returns a new instance of *WebServer
func NewWebServer(router *mux.Router) *WebServer {
	ws := &WebServer{
		router:        router,
	}

	return ws
}

// Test if the service is working
func (webserver *WebServer) pingHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("pong"))
}

// Helper function for encoding things for returning from REST calls
func (webserver *WebServer) encode(data interface{}, writer http.ResponseWriter) {
	writer.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(writer)
	err := enc.Encode(data)
	// Problems encoding
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (webserver *WebServer) metricsHandler(writer http.ResponseWriter, _ *http.Request) {
	telem := telemetry.NewSystemUsage()

	webserver.encode(telem, writer)

	return
}
func (webserver *WebServer) versionHandler(writer http.ResponseWriter, _ *http.Request) {
	type Version struct {
		Version    string `json:"version"`
	}
	version := Version{
		Version:    "0.1",
	}
	webserver.encode(version, writer)

	return
}

// AddRoute enables support to leverage the existing webserver to add routes.
func (webserver *WebServer) AddRoute(route string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	webserver.router.HandleFunc(route, handler).Methods(methods...)
}

// ConfigureStandardRoutes loads up some default routes
func (webserver *WebServer) ConfigureStandardRoutes() {

	// Ping Resource
	webserver.router.HandleFunc(constant.ApiPingRoute, webserver.pingHandler).Methods(http.MethodGet)

	// Metrics
	webserver.router.HandleFunc(constant.ApiMetricsRoute, webserver.metricsHandler).Methods(http.MethodGet)

	// Version
	webserver.router.HandleFunc(constant.ApiVersionRoute, webserver.versionHandler).Methods(http.MethodGet)
}

// SetupTriggerRoute adds a route to handle trigger pipeline from HTTP request
func (webserver *WebServer) SetupTriggerRoute(handlerForTrigger func(http.ResponseWriter, *http.Request)) {
	webserver.router.HandleFunc(constant.ApiTriggerRoute, handlerForTrigger).Methods("POST")
}

// StartHTTPServer starts the http server
func (webserver *WebServer) StartHTTPServer(errChannel chan error) {
	go func() {
		p := fmt.Sprintf(":%d", 7474)
		errChannel <- http.ListenAndServe(p, http.TimeoutHandler(webserver.router, time.Millisecond*time.Duration(constant.BootTimeoutDefault), "Request timed out"))
	}()
}
