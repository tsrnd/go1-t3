package psql

import (
	"database/sql"
  "fmt"
	"github.com/goweb3/services/crypto"
  model "github.com/goweb3/user"
  repos "github.com/goweb3/user/repository"
)

type userPsqlRepository struct {
	DB *sql.DB
}

func (m *userPsqlRepository) GetByID(id int64) (*model.User, error) {
  fmt.Println("Postgress")
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      id = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userPsqlRepository) GetByEmail(email string) (*model.User, error) {
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      email = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userPsqlRepository) GetPrivateUserDetailsByEmail(email string) (*model.PrivateUserDetails, error) {
	const query = `
    select
      id,
      password,
      salt
    from
      users
    where
      email = $1
  `
	var u model.PrivateUserDetails
	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Password, &u.Salt)
	return &u, err
}

func (m *userPsqlRepository) Create(email, name, password string) (int64, error) {
	const query = `
    insert into users (
      email,
      name,
      password,
      salt
    ) values (
      $1,
      $2,
      $3,
      $4
    )
    returning id
  `
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int64
	err := m.DB.QueryRow(query, email, name, hashedPassword, salt).Scan(&id)
	return id, err
}

func NewUserPsqlRepository(db *sql.DB) repos.UserRepository {
  return &userPsqlRepository{db}
}