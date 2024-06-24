package create_account

import (
	"testing"

	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/entity"
	"github.com.br/silva4dev/golang-event-driven-arch-project/internal/gateway/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	clientGatewayMock := mocks.NewClientGatewayMock()
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountGatewayMock := mocks.NewAccountGatewayMock()
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
