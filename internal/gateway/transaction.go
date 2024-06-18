package gateway

import "github/silva4dev/golang-event-driven-arch-project/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
