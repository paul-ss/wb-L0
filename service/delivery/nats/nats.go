package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/paul-ss/wb-L0/service/config"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/cache"
	"github.com/paul-ss/wb-L0/service/repository/postgres"
	"github.com/paul-ss/wb-L0/service/valid"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	cachedDb domain.Cache
	db       domain.Database
}

func NewHandler() *Handler {
	return &Handler{
		cachedDb: cache.NewCache(),
		db:       postgres.NewPgConn(),
	}
}

func (h *Handler) StoreOrder(m *stan.Msg) {
	log.Info("msg: ", m.Sequence)

	if err := h.db.UpdateLastMsgId(config.StanSubject, m.Sequence); err != nil {
		log.Error("update last msg id: ", err.Error())
	}

	var order domain.Order
	if err := json.Unmarshal(m.Data, &order); err != nil {
		log.Error("input message doesn't fit model: " + err.Error())
		return
	}

	if err := valid.V().Struct(&order); err != nil {
		log.Error("invalid struct: " + valid.ErrorMsg(err))
		return
	}

	if err := h.cachedDb.StoreOrder(order.Uid, m.Data); err != nil {
		log.Error("store error: " + err.Error())
		return
	}

	log.Info("stored data1 with id " + order.Uid)
}
