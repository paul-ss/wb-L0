package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/cache"
	"github.com/paul-ss/wb-L0/service/valid"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	repo domain.Repository
}

func NewHandler() *Handler {
	return &Handler{
		repo: cache.NewCache(),
	}
}

func (h *Handler) StoreOrder(m *stan.Msg) {
	var order domain.Order

	if err := json.Unmarshal(m.Data, &order); err != nil {
		log.Error("input message doesn't fit model: " + err.Error())
		return
	}

	if err := valid.V().Struct(&order); err != nil {
		log.Error("invalid struct: " + valid.ErrorMsg(err))
		return
	}

	if err := h.repo.StoreOrder(order.Uid, m.Data); err != nil {
		log.Error("store error: " + err.Error())
		return
	}

	log.Info("stored data with id" + order.Uid)
}
