package lokalkash

import (
	"errors"
	"sync"
	"time"

	"githud.com/test-task/insert/kesh"
	"githud.com/test-task/pkg/e"
)

const ErrNotFound string = "not found"
const ErrSignal string = "can't get singl side"
const ErrUpdata string = "can't updata data side"

type LokalKash struct {
	mux  sync.Mutex
	Kash map[string]time.Duration
}

func New(site ...string) *LokalKash {
	var kash LokalKash
	kash.Kash = make(map[string]time.Duration, len(site))
	for _, s := range site {
		kash.Kash[s] = time.Minute
	}
	//fmt.Println(kash.Kash)
	return &kash
}

func (l *LokalKash) Singl(domain string) (w *kesh.Website, err error) {
	defer func() { err = e.IfErr(ErrSignal, err) }()

	l.mux.Lock()

	delay := l.Kash[domain]
	if delay == 0 {
		return nil, errors.New(ErrNotFound)
	}

	l.mux.Unlock()

	return &kesh.Website{Domain: domain, Delay: delay}, nil
}

func (l *LokalKash) Updata(w *kesh.Website) (err error) {
	defer func() { err = e.IfErr(ErrUpdata, err) }()

	l.mux.Lock()

	if l.Kash[w.Domain] != 0 {
		l.Kash[w.Domain] = w.Delay
	} else {
		return errors.New(ErrNotFound)
	}

	l.mux.Unlock()
	return nil
}

func (l *LokalKash) Min() (w *kesh.Website, err error) {
	var time time.Duration = time.Minute
	var d string

	l.mux.Lock()

	for domain, delay := range l.Kash {
		if time > delay {
			time = delay
			d = domain
		}
	}

	l.mux.Unlock()

	return kesh.New(d, time), nil
}

func (l *LokalKash) Max() (w *kesh.Website, err error) {
	var time time.Duration = time.Nanosecond
	var d string

	l.mux.Lock()

	for domain, delay := range l.Kash {
		if time < delay {
			time = delay
			d = domain
		}
	}

	l.mux.Unlock()
	return kesh.New(d, time), nil
}

func (l *LokalKash) List() (list []string, err error) {
	list = make([]string,0, len(l.Kash))

	l.mux.Lock()

	for domain := range l.Kash {
		list = append(list, domain)
	}

	l.mux.Unlock()
	return list, nil
}
