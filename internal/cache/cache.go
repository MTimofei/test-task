package cache

import "time"

// данные о сайте
type Website struct {
	Domain string
	Delay  time.Duration
}

type Cache interface {
	//изменение даных
	Update(*Website) error
	//получение даных о сате по названию
	Single(string) (*Website, error)
	//получение даных о сате с минимальной задержкой
	Min() (*Website, error)
	//получение даных о сате с максимальной задержкой
	Max() (*Website, error)
	//получение всех домнеов
	List() ([]string, error)
}

func New(domain string, delay time.Duration) *Website {
	return &Website{Domain: domain, Delay: delay}
}
