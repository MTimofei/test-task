package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	confi "githud.com/test-task/insert"
	"githud.com/test-task/insert/api"
	"githud.com/test-task/insert/kesh/lokalkash"
)

func TestRoutr(t *testing.T) {
	// Создаем экземпляр объекта application
	app := api.New(lokalkash.New(confi.Domains...))

	// Создаем запросы для каждого пути
	reqSite := httptest.NewRequest("GET", "/site", nil)
	reqSiteMin := httptest.NewRequest("GET", "/site/min", nil)
	reqSiteMax := httptest.NewRequest("GET", "/site/max", nil)
	reqEndpoint := httptest.NewRequest("GET", "/endpoint", nil)

	// Создаем рекордеры ответов HTTP
	recSite := httptest.NewRecorder()
	recSiteMin := httptest.NewRecorder()
	recSiteMax := httptest.NewRecorder()
	recEndpoint := httptest.NewRecorder()

	// Вызываем функцию Routr() для получения маршрутизатора
	mux := app.Routr()

	// Выполняем запросы к каждому пути
	mux.ServeHTTP(recSite, reqSite)
	mux.ServeHTTP(recSiteMin, reqSiteMin)
	mux.ServeHTTP(recSiteMax, reqSiteMax)
	mux.ServeHTTP(recEndpoint, reqEndpoint)

	// Проверяем коды статуса ответов
	if recSite.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, получен %d", http.StatusOK, recSite.Code)
	}
	if recSiteMin.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, получен %d", http.StatusOK, recSiteMin.Code)
	}
	if recSiteMax.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, получен %d", http.StatusOK, recSiteMax.Code)
	}
	if recEndpoint.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, получен %d", http.StatusOK, recEndpoint.Code)
	}

	// Добавьте дополнительные проверки, связанные с ожидаемым поведением каждого пути

	// ...
}

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
