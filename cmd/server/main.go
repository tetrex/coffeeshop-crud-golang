package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/tetrex/coffeeshop-crud-golang/internal/db"
	"github.com/tetrex/coffeeshop-crud-golang/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	app := server.FiberApp()

	// db setup
	db.Initilize()
	// defer db.Close(db.MongoClient,context.TODO(),context.canc)
	var eg errgroup.Group
	eg.Go(func() error {
		return app.Listen(":8080")
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}

}
