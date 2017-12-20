package article

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchArticle(t *testing.T) {
	var mockArticle Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	mockUCase := new(ArticleUsecaseMock)
	mockListArticle := make([]*Article, 0)
	mockListArticle = append(mockListArticle, &mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := HttpArticleHandler{
		AUsecase: mockUCase,
	}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(ArticleUsecaseMock)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", INTERNAL_SERVER_ERROR)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := HttpArticleHandler{
		AUsecase: mockUCase,
	}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByIDInHandler(t *testing.T) {
	var mockArticle Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(ArticleUsecaseMock)

	num := int(mockArticle.ID)

	mockUCase.On("GetByID", int64(num)).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := HttpArticleHandler{
		AUsecase: mockUCase,
	}
	handler.GetByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStoreInHandler(t *testing.T) {
	mockArticle := Article{
		Title:     "Title",
		Content:   "Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(ArticleUsecaseMock)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.AnythingOfType("*article.Article")).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/article", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/article")

	handler := HttpArticleHandler{
		AUsecase: mockUCase,
	}
	handler.Store(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteInHandler(t *testing.T) {
	var mockArticle Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(ArticleUsecaseMock)

	num := int(mockArticle.ID)

	mockUCase.On("Delete", int64(num)).Return(true, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := HttpArticleHandler{
		AUsecase: mockUCase,
	}
	handler.Delete(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
