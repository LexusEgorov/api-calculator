package echomiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func testHandler(c echo.Context) error {
	return nil
}

func Test_calcMiddleware_WithAuth(t *testing.T) {
	e := echo.New()
	logger := logrus.New()
	mw := calcMiddleware{logger: logger}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "123")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := mw.WithAuth(testHandler)

	err := handler(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func Test_calcMiddleware_WithAuth_Unauthorized(t *testing.T) {
	e := echo.New()
	logger := logrus.New()
	mw := calcMiddleware{logger: logger}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := mw.WithAuth(testHandler)

	handler(ctx)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func panicHandler(c echo.Context) error {
	panic("test panic")
}

func Test_calcMiddleware_WithRecover(t *testing.T) {
	e := echo.New()
	logger := logrus.New()
	mw := calcMiddleware{logger: logger}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := mw.WithRecover(panicHandler)

	err := handler(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
