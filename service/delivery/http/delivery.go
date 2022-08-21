package http

import (
	"fmt"
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/usecase"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func NewHandler() *Handler {
	h := &Handler{
		uc:   usecase.NewUsecase(),
		tmpl: make(map[string]*template.Template),
	}

	h.tmpl["index"] = template.Must(template.New("index.html").
		ParseFiles("service/web/public/index.html"))
	return h
}

type Handler struct {
	uc   domain.Usecase
	tmpl map[string]*template.Template
}

type IndexData struct {
	Id    string
	Order string
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(":%s:\n", r.URL.Path)

	indexData := IndexData{
		Id: r.URL.Query().Get("id"),
	}

	order, err := h.uc.GetOrderById(indexData.Id)
	if err == nil {
		indexData.Order = string(order)
	}

	if err = h.tmpl["index"].Execute(w, indexData); err != nil {
		log.Error("execute template" + err.Error())
	}
}
