package http

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserLoginRequest struct
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
