package calculator

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/LexusEgorov/api-calculator/internal/storage/cache"
	"github.com/LexusEgorov/api-calculator/internal/storage/requests"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func TestCalcHandler_sendGoodResponse(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req  *http.Request
		rec  *httptest.ResponseRecorder
		body any
	}
	type response struct {
		code int
		body string
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "empty body",
			e:    *testHandler,
			args: args{
				req:  httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  httptest.NewRecorder(),
				body: nil,
			},
			want: response{
				code: http.StatusOK,
				body: "null",
			},
			wantErr: false,
		},
		{
			name: "empty array",
			e:    *testHandler,
			args: args{
				req:  httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  httptest.NewRecorder(),
				body: []float64{},
			},
			want: response{
				code: http.StatusOK,
				body: "[]",
			},
			wantErr: false,
		},
		{
			name: "JSON body",
			e:    *testHandler,
			args: args{
				req:  httptest.NewRequest(http.MethodGet, "/", nil),
				rec:  httptest.NewRecorder(),
				body: models.CalcAction{},
			},
			want: response{
				code: http.StatusOK,
				body: `{"input":"","action":"","result":0}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)

			if err := tt.e.sendGoodResponse(ctx, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.sendGoodResponse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.sendGoodResponse() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.sendGoodResponse() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}
		})
	}
}

func TestCalcHandler_sendBadResponse(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
		err error
	}
	type response struct {
		code int
		body string
	}
	tests := []struct {
		name string
		e    CalcHandler
		args args
		want response
	}{
		{
			name: "error respose",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodGet, "/", nil),
				rec: httptest.NewRecorder(),
				err: errors.New("test err"),
			},
			want: response{
				code: http.StatusBadRequest,
				body: `{"error":"test err"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)
			tt.e.sendBadResponse(ctx, tt.args.err)

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.sendGoodResponse() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.sendGoodResponse() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}
		})
	}
}

func TestCalcHandler_getBody(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name    string
		c       CalcHandler
		args    args
		want    *models.Input
		wantErr bool
	}{
		{
			name: "no body",
			c:    *testHandler,
			args: args{
				body: httptest.NewRequest(http.MethodGet, "/", nil).Body,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid input",
			c:    *testHandler,
			args: args{
				body: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2"}`)).Body,
			},
			want: &models.Input{
				Input: "1, 2",
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			c:    *testHandler,
			args: args{
				body: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2}`)).Body,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.getBody(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.getBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcHandler.getBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcHandler_HandleSum(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
		uID string
	}
	type response struct {
		code        int
		body        string
		contentType string
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2"}`)),
				rec: httptest.NewRecorder(),
				uID: "123",
			},
			want: response{
				code:        http.StatusOK,
				body:        `{"input":"1, 2","action":"SUM","result":3}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
		{
			name: "bad request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2}`)),
				rec: httptest.NewRecorder(),
			},
			want: response{
				code:        http.StatusBadRequest,
				body:        `{"error":"unexpected end of JSON input"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.Header.Set("Authorization", tt.args.uID)
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)

			if err := tt.e.HandleSum(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleSum() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.HandleSum() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.HandleSum() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcHandler.HandleSum() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}

func TestCalcHandler_HandleMult(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
		uID string
	}
	type response struct {
		code int
		body string

		contentType string
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2"}`)),
				rec: httptest.NewRecorder(),
				uID: "123",
			},
			want: response{
				code:        http.StatusOK,
				body:        `{"input":"1, 2","action":"MULT","result":2}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
		{
			name: "bad request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "1, 2}`)),
				rec: httptest.NewRecorder(),
			},
			want: response{
				code:        http.StatusBadRequest,
				body:        `{"error":"unexpected end of JSON input"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.Header.Set("Authorization", tt.args.uID)
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)

			if err := tt.e.HandleMult(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleMult() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.HandleMult() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.HandleMult() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcHandler.HandleMult() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}

func TestCalcHandler_HandleCalculate(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
		uID string
	}
	type response struct {
		code        int
		body        string
		contentType string
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "2*8"}`)),
				rec: httptest.NewRecorder(),
				uID: "123",
			},
			want: response{
				code:        http.StatusOK,
				body:        `{"input":"2*8","action":"CALC","result":16}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
		{
			name: "bad request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"input": "2*8}`)),
				rec: httptest.NewRecorder(),
			},
			want: response{
				code:        http.StatusBadRequest,
				body:        `{"error":"unexpected end of JSON input"}`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.Header.Set("Authorization", tt.args.uID)
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)

			if err := tt.e.HandleCalculate(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleCalculate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.HandleCalculate() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.HandleCalculate() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcHandler.HandleCalculate() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}

func TestCalcHandler_HandleHistory(t *testing.T) {
	testHandler := New(logrus.New(), cache.New(), requests.New())

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
		uID string
	}
	type response struct {
		code        int
		body        string
		contentType string
	}
	tests := []struct {
		name    string
		e       CalcHandler
		args    args
		want    response
		wantErr bool
	}{
		{
			name: "success request",
			e:    *testHandler,
			args: args{
				req: httptest.NewRequest(http.MethodGet, "/", nil),
				rec: httptest.NewRecorder(),
				uID: "123",
			},
			want: response{
				code:        http.StatusOK,
				body:        `[]`,
				contentType: models.HeaderApplicationJSON,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.Header.Set("Authorization", tt.args.uID)
			ctx := echo.New().NewContext(tt.args.req, tt.args.rec)

			if err := tt.e.HandleHistory(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CalcHandler.HandleHistory() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != tt.want.code {
				t.Errorf("CalcHandler.HandleHistory() code = %v, want %v", tt.args.rec.Code, tt.want.code)
			}

			if tt.args.rec.Body.String() != tt.want.body+"\n" {
				t.Errorf("CalcHandler.HandleHistory() body = %v, want %v", tt.args.rec.Body.String(), tt.want.body)
			}

			if tt.args.rec.Header().Get(models.HeaderContentTypeKey) != tt.want.contentType {
				t.Errorf("CalcHandler.HandleHistory() header %v = %v, want %v", models.HeaderContentTypeKey, tt.args.rec.Header().Get(models.HeaderContentTypeKey), tt.want.contentType)
			}
		})
	}
}
