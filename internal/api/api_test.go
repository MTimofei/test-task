package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	config "githud.com/test-task/internal"
	"githud.com/test-task/internal/api"
	loсalcache "githud.com/test-task/internal/cache/localcache"
	"githud.com/test-task/internal/processor/ping"
)

func TestAPI(t *testing.T) {

	reqValid := httptest.NewRequest(http.MethodGet, "/site", strings.NewReader(``))
	q := reqValid.URL.Query()
	q.Add("domain", "yandex.ru")
	reqValid.URL.RawQuery = q.Encode()

	reqValid1 := httptest.NewRequest(http.MethodGet, "/endpoint", strings.NewReader(``))
	q = reqValid1.URL.Query()
	q.Add("key", "admin")
	reqValid1.URL.RawQuery = q.Encode()

	testCases := []struct {
		title          string
		req            *http.Request
		rec            *httptest.ResponseRecorder
		statusExpected int
	}{
		{
			"test-valid-min",
			httptest.NewRequest(http.MethodGet, "/site/min", nil),
			httptest.NewRecorder(),
			200,
		},
		{
			"test-valid-max",
			httptest.NewRequest(http.MethodGet, "/site/max", nil),
			httptest.NewRecorder(),
			200,
		},
		{
			"test-valid-site",
			reqValid,
			httptest.NewRecorder(),
			200,
		},
		{
			"test-invalid-site",
			httptest.NewRequest(http.MethodGet, "/site", nil),
			httptest.NewRecorder(),
			400,
		},
		{
			"test-valid-endpoint",
			reqValid1,
			httptest.NewRecorder(),
			200,
		},
		{
			"test-invalid-endpoint",
			httptest.NewRequest(http.MethodGet, "/endpoint", nil),
			httptest.NewRecorder(),
			400,
		},
		{
			"test1-invalid-middleware-url",
			httptest.NewRequest(http.MethodGet, "/site/minw", nil),
			httptest.NewRecorder(),
			404,
		},
		{
			"test2-invalid-middleware-method",
			httptest.NewRequest(http.MethodPatch, "/site/min", nil),
			httptest.NewRecorder(),
			405,
		},
		{
			"test3-invalid-middleware-method",
			httptest.NewRequest(http.MethodPut, "/site/min", nil),
			httptest.NewRecorder(),
			405,
		},
	}
	k := loсalcache.New(config.Domains...)
	p := ping.New(k)
	err := p.Start()
	if err != nil {
		t.Error(err)
	}

	a := api.New(k)
	mux := a.Router()
	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			mux.ServeHTTP(tC.rec, tC.req)
			if tC.rec.Code != tC.statusExpected {
				t.Errorf("Ожидался код статуса %d, получен %d", tC.statusExpected, tC.rec.Code)
			}
		})
	}
}
