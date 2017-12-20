package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/monstar-lab/fr-circle-api/article"
	"github.com/monstar-lab/fr-circle-api/gist"
	"github.com/monstar-lab/fr-circle-api/infrastructure"
	"github.com/monstar-lab/fr-circle-api/shared/handler"
	mMiddleware "github.com/monstar-lab/fr-circle-api/shared/middleware"
)

// Router is application struct hold mux and db connection
type Router struct {
	mux                *chi.Mux
	sqlHandler         *infrastructure.SQLHandler
	cacheHandler       *infrastructure.CacheHandler
	loggerHandler      *infrastructure.LoggerHandler
	translationHandler *infrastructure.TranslationHandler
}

// InitializeRouter initializes mux and middleware
func (r *Router) InitializeRouter() {
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	r.mux.Use(r.translationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.mux.Use(mMiddleware.Logger(r.loggerHandler))
	// Custom middleware(Header)
	r.mux.Use(mMiddleware.Header(r.loggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.loggerHandler.Log)
	r.mux.NotFound(eh.StatusNotFound)
	r.mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	// base set.
	bh := handler.NewHTTPBaseHandler(r.loggerHandler.Log)

	// article set.
	ar := article.NewArticleRepository(r.sqlHandler.Database, r.cacheHandler.Conn)
	au := article.NewUsecase(r.sqlHandler.Database, ar)
	ah := article.NewArticleHTTPHandler(bh, au)

	r.mux.Get("/article", ah.Get)
	r.mux.Get("/article/{id}", ah.GetID)
	r.mux.Post("/article/add", ah.PostAdd)
	r.mux.Delete("/article/{id}", ah.DeleteID)
	r.mux.Get("/article/count", ah.GetCount)
	r.mux.Post("/article/count", ah.PostCount)
	r.mux.Post("/article/visenze/discoversearch", ah.PostVisenzeDiscoversearch)

	// gist
	gr := gist.NewGistRepository()
	gu := gist.NewGistUsecase(gr)
	gh := gist.NewGistHTTPHandler(bh, gu)
	r.mux.Get("/gist", gh.ListGists)
	// job
	// product
}
