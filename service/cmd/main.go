package main

import (
	"context"
	"fmt"
	"github.com/paul-ss/wb-L0/service/config"
	"github.com/paul-ss/wb-L0/service/delivery/nats"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	defer func() {
		err := recover()
		fmt.Fprintln(os.Stderr, err)
	}()

	// Nats
	ns := NewNatsService()
	nh := nats.NewHandler()
	ns.Subscribe(config.StanSubject, nh.StoreOrder)
	defer ns.Close()

	// Http
	srv := NewHttpServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c

		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}()

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
