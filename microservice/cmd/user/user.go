package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"

	"os/signal"

	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user"
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user/endpoints"
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user/transport"
	"github.com/go-kit/kit/log"
	"github.com/oklog/run"
)

const (
	defaultHTTPPort = "8080"
)

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
	)

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8080", "caller", log.DefaultCaller)

	var (
		db, _       = user.OpenDbConnection()
		repo        = user.NewSqliteRepository(db)
		service     = user.NewService(repo)
		e           = endpoints.NewEndpoints(service)
		httpHandler = transport.NewHttpHandler(e, logger)
	)

	logger.Log("msg", "HTTP", "addr", "8080")
	logger.Log("err", http.ListenAndServe(":8080", httpHandler))

	var g run.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
