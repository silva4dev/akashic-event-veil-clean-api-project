package gateway

import "github/silva4dev/golang-event-driven-arch-project/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
