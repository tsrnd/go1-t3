package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/goweb3/services/crypto"
	"github.com/goweb3/user/usecase"
	"github.com/goweb3/services/cache"
)

// UserController type
type UserController struct {
  Usecase usecase.UserUsecase
  Cache cache.Cache
}

// NewUserController func
func NewUserController(r chi.Router, uc usecase.UserUsecase, c cache.Cache) *UserController {
	handler := &UserController{
		Usecase: uc,
		Cache:   c,
	}
	
	r.Post("/users", handler.UserRegister)
	r.Post("/auth", handler.UserLogin)
	r.Get("/users", handler.Index)
	return handler
}

// UserRegister func
func (ctrl *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var rr UserRegisterRequest
	err := decoder.Decode(&rr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := ctrl.Usecase.Create(rr.Email, rr.Name, rr.Password)
	if err != nil {
		log.Fatalf("Add user to database error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Generate token Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = ctrl.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(int(id)), oneMonth)
	if err != nil {
		log.Fatalf("Add token to redis Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// UserLogin func
func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var lr UserLoginRequest
	err := decoder.Decode(&lr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := ctrl.Usecase.GetPrivateUserDetailsByEmail(lr.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	password := crypto.HashPassword(lr.Password, user.Salt)
	if user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}
	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = ctrl.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(int(user.ID)), oneMonth)
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (ctrl *UserController) Index(w http.ResponseWriter, r *http.Request) {
	user,_ := ctrl.Usecase.GetByID(1)
	p := map[string]interface{}{
		"token": user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
