package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/paul-ss/wb-L0/service/config"
	http2 "github.com/paul-ss/wb-L0/service/delivery/http"
	"net/http"
	"os"
)

type NatsService struct {
	conn stan.Conn
	subs map[string]*stan.Subscription
}

func NewNatsService() *NatsService {
	conn, err := stan.Connect(config.StanClusterId, config.StanClientId)
	if err != nil {
		panic(err)
	}

	return &NatsService{
		conn: conn,
		subs: make(map[string]*stan.Subscription),
	}
}

func (ns *NatsService) Subscribe(subName string, f func(msg *stan.Msg)) {
	sub, err := ns.conn.Subscribe(subName, f, stan.DeliverAllAvailable())

	if err != nil {
		panic(err)
	}

	ns.subs[config.StanSubject] = &sub
}

func (ns *NatsService) Close() {
	for _, s := range ns.subs {
		err := (*s).Unsubscribe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	if err := ns.conn.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func NewHttpServer() *http.Server {
	mux := http.NewServeMux()

	handler := http2.NewHandler()
	mux.HandleFunc("/", handler.MainPage)

	return &http.Server{Addr: config.ServerAddress, Handler: mux}
}
