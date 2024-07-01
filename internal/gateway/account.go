package gateway

import "github.com.br/silva4dev/golang-event-driven-arch-project/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
