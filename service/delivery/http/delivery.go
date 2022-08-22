package http

import (
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/cache"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	mainPageName = "index.html"
	templatesDir = "service/web/public"
)

func NewHandler() *Handler {
	h := &Handler{
		repo: cache.NewCache(),
		tmpl: make(map[string]*template.Template),
	}

	h.tmpl[mainPageName] = template.Must(template.New(mainPageName).
		ParseFiles(filepath.Join(templatesDir, mainPageName)))
	return h
}

type Handler struct {
	repo domain.Cache
	tmpl map[string]*template.Template
}

type MainPageData struct {
	Id    string
	Order string
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	indexData := MainPageData{
		Id: r.URL.Query().Get("id"),
	}

	order, err := h.repo.GetOrderById(indexData.Id)
	if err == nil {
		indexData.Order = string(order)
	}

	if err = h.tmpl[mainPageName].Execute(w, indexData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("execute template" + err.Error())
	}
}

func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Error("recover mw: ", err)
			}
		}()

		next(w, r)
	}
}
