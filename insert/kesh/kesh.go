package kesh

import "time"

// данные о сайте
type Website struct {
	Url   string
	Delay time.Duration
}

type Kesh interface {
	//обновление данных в кеше
	Updata() error
	//получение даных о сате по названию
	Singl(string) (*Website, error)
	//получение даных о сате с минимальной задержкой
	Min() (*Website, error)
	//получение даных о сате с максимальной задержкой
	Max() (*Website, error)
}
