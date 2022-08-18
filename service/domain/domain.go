package domain

type Repository interface {
	StoreOrder(id string, data []byte) error
	GetOrderById(id string) ([]byte, error)
}

type Usecase interface {
	StoreOrder(data []byte) error
	GetOrderById(id string) ([]byte, error)
}
