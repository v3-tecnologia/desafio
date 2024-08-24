package webserver

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
	BasePath      string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
		BasePath:      "/api/v1/telemetry",
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	fullPath := "/api/v1/telemetry" + path
	if path == "/docs/*" {
		fullPath = path
	}
	if s.Handlers[fullPath] == nil {
		s.Handlers[fullPath] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[fullPath][method] = handler
}

func (s *WebServer) Start() error {
	s.Router.Use(middleware.Logger)

	for path, methodHandlers := range s.Handlers {
		for method, handler := range methodHandlers {
			switch method {
			case http.MethodPost:
				s.Router.Post(path, handler)
			case http.MethodGet:
				s.Router.Get(path, handler)
			}
		}
	}
	return http.ListenAndServe(s.WebServerPort, s.Router)
}
