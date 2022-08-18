package usecase

import (
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/cache"
)

type Usecase struct {
	repo domain.Repository
}

func NewUsecase() *Usecase {
	return &Usecase{
		repo: cache.NewCache(),
	}
}

func (uc *Usecase) StoreOrder(data []byte) error {
	return uc.repo.StoreOrder("1", data)
}

func (uc *Usecase) GetOrderById(id string) ([]byte, error) {
	return uc.repo.GetOrderById(id)
}
