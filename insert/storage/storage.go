package storage

// содержит даные о запросах по эндпоитам
type Data struct {
}

type Sotorage interface {
	//обновление данах
	Updata(*Data) error
	//получение всех даных
	All() (*Data, error)
}
