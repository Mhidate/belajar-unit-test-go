package service

import (
	"errors"
	"testing"

	"github.com/Mhidate/belajar-unit-test-go/my_website/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock struct dari UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(id int) (*model.User, error) {
	args := m.Called(id)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserService_GetUserName(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := UserService{Repo: mockRepo}

	// Data dummy
	expectedUser := &model.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Test sukses
	mockRepo.On("GetByID", 1).Return(expectedUser, nil)

	name, err := service.GetUserName(1)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", name)

	mockRepo.AssertCalled(t, "GetByID", 1)

	// Test error (user not found)
	mockRepo.On("GetByID", 2).Return(nil, errors.New("user not found"))

	name, err = service.GetUserName(2)
	assert.Error(t, err)
	assert.Equal(t, "", name)
	assert.Equal(t, "user not found", err.Error())

	mockRepo.AssertCalled(t, "GetByID", 2)
	mockRepo.AssertExpectations(t)
}
