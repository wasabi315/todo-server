package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/wasabi315/todo-server/repository"
	"github.com/wasabi315/todo-server/router"
)

func main() {
	repo, err := repository.NewMapRepository()
	if err != nil {
		return
	}

	e := router.Setup(&router.Config{
		Repo: repo,
	})

	go func() {
		if err := e.Start(":3000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
