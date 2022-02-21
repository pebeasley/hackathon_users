package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"

	"github.com/pebeasley/users/database"
	"github.com/pebeasley/users/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DB.AutoMigrate(&users.User{})
}

func main() {
	fmt.Println("hello from users")
	var httpAddress = flag.String("http", ":9010", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	initDatabase()

	flag.Parse()
	ctx := context.Background()
	var srv users.Service
	{
		srv = users.NewService()
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := users.MakeEndpoints(srv)

	go func() {
		fmt.Println("Listening on port", *httpAddress)
		handler := users.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddress, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
