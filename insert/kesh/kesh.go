package kesh

import "time"

// данные о сайте
type WebsiteData struct {
	Url   string
	Delay time.Duration
}

type Kesh interface {
	//обновление данных в кеше
	Updata() error
	//получение даных о сате по названию
	Singl(string) (*WebsiteData, error)
	//получение даных о сате с минимальной задержкой
	Min() (*WebsiteData, error)
	//получение даных о сате с максимальной задержкой
	Max() (*WebsiteData, error)
}
