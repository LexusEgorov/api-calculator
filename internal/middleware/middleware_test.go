package echomiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func testHandler(c echo.Context) error {
	type ok struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, ok{
		Status: "ok",
	})
}

func panicHandler(c echo.Context) error {
	panic("test")
}

func Test_calcMiddleware_WithAuth(t *testing.T) {
	testMiddleware := New(logrus.New())

	type response struct {
		code        int
		body        string
		contentType string
	}
	type args struct {
		next echo.HandlerFunc
		req  http.Request
		rec  httptest.ResponseRecorder
		uID  string
	}
	tests := []struct {
		name    string
		c       calcMiddleware
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success auth",
			c:    *testMiddleware,
			args: args{
				next: testHandler,
				req:  *httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  *httptest.NewRecorder(),
				uID:  "123",
			},
			want: response{
				code:        http.StatusOK,
				body:        `{"status":"ok"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
		{
			name: "bad auth",
			c:    *testMiddleware,
			args: args{
				next: testHandler,
				req:  *httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  *httptest.NewRecorder(),
				uID:  "",
			},
			want: response{
				code:        http.StatusUnauthorized,
				body:        `{"error":"user not found"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.Header.Set("Authorization", tt.args.uID)
			ctx := echo.New().NewContext(&tt.args.req, &tt.args.rec)
			handler := tt.c.WithAuth(tt.args.next)

			if err := handler(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcMiddleware.WithAuth() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcMiddleware.WithAuth() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcMiddleware.WithAuth() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcMiddleware.WithAuth() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}

func Test_calcMiddleware_WithRecover(t *testing.T) {
	testMiddleware := New(logrus.New())

	type response struct {
		code        int
		body        string
		contentType string
	}
	type args struct {
		next echo.HandlerFunc
		req  http.Request
		rec  httptest.ResponseRecorder
	}
	tests := []struct {
		name    string
		c       calcMiddleware
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success recover",
			c:    *testMiddleware,
			args: args{
				next: panicHandler,
				req:  *httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  *httptest.NewRecorder(),
			},
			want: response{
				code:        http.StatusInternalServerError,
				body:        `{"error":"internal server error"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
		{
			name: "without recover",
			c:    *testMiddleware,
			args: args{
				next: testHandler,
				req:  *httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  *httptest.NewRecorder(),
			},
			want: response{
				code:        http.StatusOK,
				body:        `{"status":"ok"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := echo.New().NewContext(&tt.args.req, &tt.args.rec)
			handler := tt.c.WithRecover(tt.args.next)

			if err := handler(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcMiddleware.WithRecover() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcMiddleware.WithRecover() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcMiddleware.WithRecover() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcMiddleware.WithRecover() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}
