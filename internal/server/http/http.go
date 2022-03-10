package http

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:           "0.0.0.0:" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}

	logrus.WithFields(logrus.Fields{
		"package": "internal.server",
		"layer":   "server",
	}).Infof("server listening at http://localhost:%s", port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
