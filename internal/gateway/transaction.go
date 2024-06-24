package gateway

import "github.com.br/silva4dev/golang-event-driven-arch-project/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
