package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/cmd/app/config"
	"github.com/wascript3r/reservio/cmd/app/registry"
	corsMid "github.com/wascript3r/reservio/pkg/cors/delivery/http/middleware"
)

const (
	// Database
	DatabaseDriver = "postgres"
)

var (
	Cfg *config.Config
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	Cfg, err = config.LoadConfig(false)
	if err != nil {
		fatalError(err)
	}

	os.Setenv("TZ", "UTC")
}

func fatalError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// Logging
	loggerReg := registry.NewLogger(Cfg)
	logger := loggerReg.Usecase()

	// Startup message
	logger.Info("... Starting app ...")

	// Database connection
	dbConn, err := openDatabase(DatabaseDriver, Cfg.Database.Postgres.DSN)
	if err != nil {
		fatalError(err)
	}
	logger.Info("... Connected to database ...")

	mapper := httpjson.NewCodeMapper()

	// Registries
	globalReg := registry.NewGlobal(Cfg, dbConn, mapper)
	userReg := registry.NewUser(globalReg)
	companyReg := registry.NewCompany(globalReg)
	serviceReg := registry.NewService(globalReg)
	reservationReg := registry.NewReservation(globalReg)
	clientReg := registry.NewClient(globalReg)
	tokenReg := registry.NewToken(globalReg)

	globalReg.
		SetUserReg(userReg).
		SetCompanyReg(companyReg).
		SetServiceReg(serviceReg).
		SetReservationReg(reservationReg).
		SetClientReg(clientReg).
		SetTokenReg(tokenReg).
		SetLoggerReg(loggerReg)

	// Graceful shutdown
	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// HTTP server
	httpRouter := httprouter.New()
	httpRouter.MethodNotAllowed = MethodNotAllowedHnd
	httpRouter.NotFound = NotFoundHnd

	if Cfg.HTTP.EnablePprof {
		// pprof
		httpRouter.Handler(http.MethodGet, "/debug/pprof/*item", http.DefaultServeMux)
	}

	userReg.ServeHTTP(httpRouter)
	companyReg.ServeHTTP(httpRouter)
	serviceReg.ServeHTTP(httpRouter)
	reservationReg.ServeHTTP(httpRouter)
	clientReg.ServeHTTP(httpRouter)
	tokenReg.ServeHTTP(httpRouter)

	httpServer := &http.Server{
		Addr:    ":" + Cfg.HTTP.Port,
		Handler: httpRouter,
	}
	if Cfg.HTTP.CORS != nil {
		httpServer.Handler = corsMid.NewHTTPMiddleware(
			Cfg.HTTP.CORS.AllowOrigin,
		).EnableCors(httpRouter)
	} else {
		httpServer.Handler = httpRouter
	}

	// Graceful shutdown
	gracefulShutdown := func() {
		if err := httpServer.Shutdown(context.Background()); err != nil {
			logger.Error("Cannot shutdown HTTP server: %s", err)
		}

		logger.Info("... Exited ...")
		os.Exit(0)
	}

	go func() {
		<-stopSig
		logger.Info("... Received stop signal ...")
		gracefulShutdown()
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Println(err)
			gracefulShutdown()
		}
	}

	var c chan struct{}
	<-c
}
