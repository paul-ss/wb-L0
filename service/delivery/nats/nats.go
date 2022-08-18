package nats

import (
	"github.com/nats-io/stan.go"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/usecase"
)

type Handler struct {
	uc domain.Usecase
}

func NewHandler() *Handler {
	return &Handler{
		uc: usecase.NewUsecase(),
	}
}

func (h *Handler) StoreOrder(m *stan.Msg) {

}
