package gateway

import "github.com.br/silva4dev/golang-event-driven-arch-project/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
