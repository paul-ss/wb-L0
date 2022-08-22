package main

import (
	"context"
	"github.com/paul-ss/wb-L0/service/config"
	"github.com/paul-ss/wb-L0/service/delivery/nats"
	"github.com/paul-ss/wb-L0/service/repository/postgres"
	service "github.com/paul-ss/wb-L0/service/srv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	defer postgres.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Error("main recover: ", err)
		}
	}()

	// Nats
	ns := service.NewNatsService()
	nh := nats.NewHandler()
	ns.Subscribe(config.StanSubject, nh.StoreOrder)
	defer ns.Close()

	// Http
	srv := service.NewHttpServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c

		log.Info("shutdown server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Error("http server shutdown: ", err.Error())
		}
	}()

	log.Info("http server listens at ", srv.Addr)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Error("http listen and serve: ", err.Error())
	}
}
