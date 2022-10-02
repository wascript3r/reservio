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

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	"github.com/wascript3r/reservio/cmd/app/registry"
)

const (
	// Database
	DatabaseDriver = "postgres"
)

var (
	Cfg *registry.Config
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	Cfg, err = registry.LoadConfig()
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

	// Registries
	globalReg := registry.NewGlobal(Cfg, dbConn)
	userReg := registry.NewUser(globalReg)
	companyReg := registry.NewCompany(globalReg)
	serviceReg := registry.NewService(globalReg)
	reservationReg := registry.NewReservation(globalReg)

	globalReg.
		SetUserReg(userReg).
		SetCompanyReg(companyReg).
		SetServiceReg(serviceReg).
		SetReservationReg(reservationReg).
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

	httpServer := &http.Server{
		Addr:    ":" + Cfg.HTTP.Port,
		Handler: httpRouter,
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
