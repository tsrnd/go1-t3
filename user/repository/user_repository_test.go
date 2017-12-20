package repository_test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goweb3/user/repository"
)


func TestGetByIdSuccess(t *testing.T) {
	db, sqlmock, err := sqlmock.New()
    if err != nil {
        t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	sqlmock.ExpectBegin()
	userRepositoryTest = repository.NewUserRepository(db)
}