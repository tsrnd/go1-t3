package article

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/monstar-lab/fr-circle-api/shared/handler"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// HTTPArticleHandler struct.
type HTTPArticleHandler struct {
	handler.HTTPBaseHandler
	aUsecase ArticleUsecase
}

// Get get all from article.
func (h *HTTPArticleHandler) Get(w http.ResponseWriter, r *http.Request) {

	article, err := h.aUsecase.Get("")
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("h.aUsecase.Get errors")

		h.StatusServerError(w)
		return
	}

	// translation
	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("success")}
	jsonData := ArticleGetResponse{CommonResponse: common, ResponseArticle: article}
	h.ResponseJSON(w, jsonData)
}

// GetID search id from article.
func (h *HTTPArticleHandler) GetID(w http.ResponseWriter, r *http.Request) {
	// get url params
	id := chi.URLParam(r, "id")

	// validate get data.
	v := validator.New()
	err := v.Struct(ArticleGetIDRequest{ID: id})

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// change string to int
	iid, _ := strconv.Atoi(id)

	article, err := h.aUsecase.GetID(iid)
	if err != nil {
		T := h.GetTranslaterFunc(r)
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal(T("error", map[string]interface{}{
			"Function": "h.aUsecase.GetID",
		}))
		h.StatusServerError(w)
		return
	}

	// translation
	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("success")}
	jsonData := ArticleGetTitleResponse{CommonResponse: common, ResponseArticle: article}
	h.ResponseJSON(w, jsonData)
}

// PostAdd method.
func (h *HTTPArticleHandler) PostAdd(w http.ResponseWriter, r *http.Request) {
	var err error
	// mapping post to struct.
	requestArticlePostAdd := ArticlePostAddRequest{}
	h.Parse(r, &requestArticlePostAdd)

	// log output.
	h.Logger.WithFields(logrus.Fields{
		"r.PostForm":      r.PostForm,
		"r.MultipartForm": r.MultipartForm,
	}).Debug("post data.")

	// log output.
	h.Logger.WithFields(logrus.Fields{
		"Title":   r.PostFormValue("title"),
		"Content": r.PostFormValue("content"),
	}).Debug("post data.")

	h.Logger.WithFields(logrus.Fields{
		"struct": requestArticlePostAdd,
	}).Debug("struct data.")

	// validate post data.
	v := validator.New()
	err = v.Struct(requestArticlePostAdd)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}
	rowAffected, err := h.aUsecase.Add(requestArticlePostAdd.Title, requestArticlePostAdd.Content)
	if err != nil {
		h.StatusServerError(w)
		return
	}
	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("add", map[string]interface{}{"Title": "article", "RowAffected": rowAffected})}
	h.ResponseJSON(w, common)
}

// DeleteID delete from article.
func (h *HTTPArticleHandler) DeleteID(w http.ResponseWriter, r *http.Request) {
	var err error
	// get url params
	id := chi.URLParam(r, "id")

	// validate get data.
	v := validator.New()
	err = v.Struct(ArticleDeleteIDRequest{ID: id})

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// change string to int
	iid, _ := strconv.Atoi(id)

	rowAffected, err := h.aUsecase.Delete(iid)
	if err != nil {
		h.StatusServerError(w)
		return
	}

	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("delete", map[string]interface{}{"Title": "article", "RowAffected": rowAffected})}
	h.ResponseJSON(w, common)
}

// GetCount get count from redis.
func (h *HTTPArticleHandler) GetCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.aUsecase.GetCount()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("ArticleUsecase.GetCount() error.")
		h.StatusServerError(w)
		return
	}
	common := CommonResponse{Result: 0, Message: ""}
	jsonData := ArticleGetCountResponse{CommonResponse: common, Count: count}
	h.ResponseJSON(w, jsonData)
}

// PostCount get count from redis.
func (h *HTTPArticleHandler) PostCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.aUsecase.AddCount()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("ArticleUsecase.GetCount() error.")
		h.StatusServerError(w)
		return
	}
	common := CommonResponse{Result: 0, Message: ""}
	jsonData := ArticlePostCountResponse{CommonResponse: common, Count: count}
	h.ResponseJSON(w, jsonData)
}

// PostVisenzeDiscoversearch requests visenze/discoversearch API
func (h *HTTPArticleHandler) PostVisenzeDiscoversearch(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := ArticlePostVisenzeDiscoverSearchRequest{}
	h.ParseMultipart(r, &request)
	h.Logger.Debug(request)

	// file.
	f, fh, _ := r.FormFile("file")
	request.File = f
	// file close.
	if f != nil {
		defer f.Close()
	}

	// validate get data.
	v := validator.New()
	err := v.Struct(request)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// save file.
	saveFilename := h.GetRandomTempFileName("upload_", fh.Filename)
	h.SaveToFile(r, "file", saveFilename)

	// request visenze/dicoversearch API.
	discoversearch, err := h.aUsecase.GetDiscoversearch(request.Page, request.ResultLimit, saveFilename)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("ArticleUsecase.GetDiscoversearch() error.")
		h.StatusServerError(w)
		return
	}

	common := CommonResponse{Result: 0, Message: ""}
	jsonData := ArticlePostVisenzeDiscoversearchResponse{CommonResponse: common, VisenzeDiscoversearchResponse: *discoversearch}
	h.ResponseJSON(w, jsonData)
}

// NewArticleHTTPHandler responses new HTTPArticleHandler instance.
func NewArticleHTTPHandler(bh *handler.HTTPBaseHandler, us ArticleUsecase) *HTTPArticleHandler {
	return &HTTPArticleHandler{HTTPBaseHandler: *bh, aUsecase: us}
}
