package echosrv

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/LexusEgorov/api-calculator/internal/calculator"
	"github.com/LexusEgorov/api-calculator/internal/storage/cache"
	"github.com/LexusEgorov/api-calculator/internal/storage/requests"
	"github.com/sirupsen/logrus"
)

func TestServer_Run_Stop(t *testing.T) {
	testPort := 8080
	testLogger := logrus.New()
	testHandler := calculator.New(testLogger, cache.New(), requests.New())
	testServer := New(testHandler, testLogger, testPort)

	go testServer.Run()
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	defer testServer.Stop(ctx)

	time.Sleep(time.Second * 1)

	t.Run("check running", func(t *testing.T) {
		if _, err := http.Get(fmt.Sprintf("http://localhost:%d", testPort)); err != nil {
			t.Fatalf("Server.Run() error = %v,", err)
		}
	})

	t.Run("check stopping", func(t *testing.T) {
		if err := testServer.Stop(ctx); err != nil {
			t.Fatalf("Server.Stop() error = %v,", err)
		}
	})
}
