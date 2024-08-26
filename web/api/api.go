package api

import (
	"context"
	"desafio-backend/internal/gyroscope"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	port = ":7001"
)

type API struct {
	// Dependencies
	gyroscopeMain gyroscope.UseCases
}

func NewAPI(
	gyroscopeMain gyroscope.UseCases,
) *mux.Router {
	api := API{
		gyroscopeMain: gyroscopeMain,
	}
	router := mux.NewRouter()
	api.health(router)

	s := router.PathPrefix("/api").Subrouter()
	s.Use(TraceIDMiddleware, LogMiddleware, RecoverMiddleware)

	api.newV1Api(s)

	return router
}

func (api *API) health(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(":)"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

func Start(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// run http server
	go func() {
		logrus.Infof("listenning and serving on port %s", port)
		printEndpoints(router)
		// service connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalln(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	// greacefull shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	logrus.Info("starting shutdown server")
	defer logrus.Info("finishing shutdown server")

	const timeout = 5
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Error("could not shutdown server greacefully")

	}
	<-ctx.Done()
}

func printEndpoints(r *mux.Router) {
	logrus.Info("api routes")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		logrus.Info(fmt.Sprintf("%v %s", methods, path))
		return nil
	})
}
