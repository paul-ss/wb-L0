package http

import (
	"fmt"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewHandler() *Handler {
	return &Handler{
		uc: usecase.NewUsecase(),
	}
}

type Handler struct {
	uc domain.Usecase
}

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	order, err := h.uc.GetOrderById(id)

	if err != nil {
		log.Error("http delivery: " + err.Error())
	}

	fmt.Fprintln(w, order)
}
