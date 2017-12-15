package config

import (
	"database/sql"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	productRepo "github.com/goweb3/product/repository/psql"
	productCase "github.com/goweb3/product/usecase"
	productDeliver "github.com/goweb3/product/delivery/http"
	"github.com/goweb3/services/cache"
	userRepo "github.com/goweb3/user/repository"
	userCase "github.com/goweb3/user/usecase"
	userDeliver "github.com/goweb3/user/delivery/http"
)

// Router func
func Router(db *sql.DB, c cache.Cache) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addUserRoutes(r, db, c)
	addProductRoutes(r, db, c)
	return r
}

func addUserRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

func addProductRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := productRepo.NewProductRepository(db)
	uc := productCase.NewProductUsecase(repo)
	productDeliver.NewProductController(r, uc, c)
}
