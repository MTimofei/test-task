package storage

// содержит даные о запросах по эндпоитам
type Endpoints struct {
}

type Sotorage interface {
	//обновление данах
	Updata(*Endpoints) error
	//получение всех даных
	All() (*Endpoints, error)
}
