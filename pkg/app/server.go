package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

type Server struct {
	*fiber.App
	Cfg *Config
}

func New(config ...*Config) *Server {
	var cfg *Config
	if len(config) > 0 {
		cfg = config[0]
	}
	srv := fiber.New()
	return &Server{App: srv, Cfg: cfg}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func (srv *Server) StartServerWithGracefulShutdown() {
	// Create channel for idleChannel connections.
	idleChannel := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := srv.Shutdown(); err != nil {
			logrus.WithFields(logrus.Fields{
				"topic":  "Shutdown failure",
				"reason": fmt.Sprintf("%v", err),
			}).Error("Unable to shutdown server.")
		}

		close(idleChannel)
	}()

	// Run server.
	addr := srv.Cfg.Host + ":" + srv.Cfg.Port
	if err := srv.Listen(addr); err != nil {
		logrus.WithFields(logrus.Fields{
			"topic":  "Server status",
			"status": "Not running",
			"reason": fmt.Sprintf("%v", err),
		}).Fatal("Server is not running.")
	}

	<-idleChannel
}

// StartServer func for starting a simple server.
func (srv *Server) StartServer() {
	// Run server.
	addr := srv.Cfg.Host + ":" + srv.Cfg.Port
	if err := srv.Listen(addr); err != nil {
		logrus.WithFields(logrus.Fields{
			"topic":  "Server status",
			"status": "Not running",
			"reason": fmt.Sprintf("%v", err),
		}).Fatal("Server is not running.")
	}
}
