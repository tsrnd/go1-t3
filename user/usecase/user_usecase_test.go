package usecase_test

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/goweb3/user/repository/mock"
	"github.com/goweb3/user/usecase"
	"github.com/bxcodec/faker"
	"github.com/goweb3/user"
	"github.com/stretchr/testify/assert"
)
func TestGetByIdSuccess(t *testing.T) {
	userFake := user.User{}
	faker.FakeData(&userFake)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepo := mock_repository.NewMockUserRepository(mockCtrl)
	mockUserRepo.EXPECT().GetByID(int64(userFake.ID)).Return(&userFake, nil)
	usecaseTest := usecase.NewUserUsecase(mockUserRepo)
	user, err := usecaseTest.GetByID(int64(userFake.ID))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user, &userFake, "Test success!")
}