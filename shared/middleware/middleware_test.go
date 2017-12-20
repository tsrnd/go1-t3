package middleware_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/monstar-lab/fr-circle-api/middleware"

	test "net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	res := test.NewRecorder()
	c := e.NewContext(req, res)
	m := middleware.InitMiddleware()

	h := m.CORS(echo.HandlerFunc(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}))
	h(c)

	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
