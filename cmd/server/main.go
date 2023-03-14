package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/tetrex/coffeeshop-crud-golang/internal/db"
	"github.com/tetrex/coffeeshop-crud-golang/internal/server"
)

func main() {
	// app setup
	// with routes and middleware
	app := server.FiberApp()

	// db setup
	db.Initilize()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	// ...

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}
	serverShutdown.Wait()

	// dp connection close
	db.Close(db.MongoClient, context.TODO(), *db.MongoCancleFunc)
	// Your cleanup tasks go here

}
