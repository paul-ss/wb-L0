package srv

import (
	"github.com/nats-io/stan.go"
	"github.com/paul-ss/wb-L0/service/config"
	httpDelivery "github.com/paul-ss/wb-L0/service/delivery/http"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/postgres"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NatsService struct {
	conn stan.Conn
	subs map[string]*stan.Subscription
	repo domain.Database
}

func NewNatsService() *NatsService {
	conn, err := stan.Connect(config.StanClusterId, config.StanClientId)
	if err != nil {
		panic(err)
	}
	log.Infof("connected to stan cluster '%s'", config.StanClusterId)

	return &NatsService{
		conn: conn,
		subs: make(map[string]*stan.Subscription),
		repo: postgres.NewPgConn(),
	}
}

func (ns *NatsService) Subscribe(subName string, f func(msg *stan.Msg)) {
	var (
		sub stan.Subscription
		err error
	)

	id, ok := ns.repo.GetLastMsgId(subName)
	if !ok {
		sub, err = ns.conn.Subscribe(subName, f, stan.DeliverAllAvailable())
	} else {
		sub, err = ns.conn.Subscribe(subName, f, stan.StartAtSequence(id+1))
	}

	if err != nil {
		panic(err)
	}

	ns.subs[config.StanSubject] = &sub
}

func (ns *NatsService) Close() {
	for _, s := range ns.subs {
		err := (*s).Unsubscribe()
		if err != nil {
			log.Error("nats service close: ", err.Error())
		}
	}

	if err := ns.conn.Close(); err != nil {
		log.Error("nats service close: ", err.Error())
	}
}

func NewHttpServer() *http.Server {
	mux := http.NewServeMux()

	handler := httpDelivery.NewHandler()
	mux.HandleFunc("/", httpDelivery.RecoverMiddleware(handler.MainPage))

	return &http.Server{Addr: config.ServerAddress, Handler: mux}
}
