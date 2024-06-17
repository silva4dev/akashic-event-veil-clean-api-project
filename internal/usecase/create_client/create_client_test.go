package createclient

import (
	"github/silva4dev/golang-event-driven-arch-project/internal/gateway/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := mocks.NewClientGatewayMock()
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	input := &CreateClientInputDTO{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "john@doe.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
