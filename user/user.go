package user

// User struct
type User struct {
	ID    int64  `json:"id" faker:"unix_time"`
	Email string `json:"email" faker:"email"`
	Name  string `json:"name" faker:"name"`
}

// PrivateUserDetails struct
type PrivateUserDetails struct {
	ID       int64 `faker:"unix_time"`
	Password string `faker:"password"`
	Salt     string `faker:"word"`
}
