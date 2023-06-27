package loсalcache

import (
	"errors"
	"sync"
	"time"

	config "githud.com/test-task/internal"
	"githud.com/test-task/internal/cache"
	"githud.com/test-task/pkg/e"
)

const ErrNotFound string = "not found"
const ErrSignal string = "can't get single side"
const ErrUpdate string = "can't update data side"
const ErrSiteUnavailable string = "site is unavailable"

type LocalCache struct {
	mux   sync.Mutex
	Cache map[string]time.Duration
}

func New(site ...string) *LocalCache {
	var cache LocalCache
	cache.Cache = make(map[string]time.Duration, len(site))
	for _, s := range site {
		cache.Cache[s] = time.Hour
	}
	return &cache
}

func (l *LocalCache) Single(domain string) (w *cache.Website, err error) {
	defer func() { err = e.IfErr(ErrSignal, err) }()

	l.mux.Lock()
	defer l.mux.Unlock()

	delay := l.Cache[domain]

	switch {
	case delay == 0:
		return nil, errors.New(ErrNotFound)
	case delay > time.Second*time.Duration(config.Timeout):
		return nil, errors.New(ErrNotFound)
	default:

	}

	return &cache.Website{Domain: domain, Delay: delay}, nil
}

func (l *LocalCache) Update(w *cache.Website) (err error) {
	defer func() { err = e.IfErr(ErrUpdate, err) }()

	l.mux.Lock()
	defer l.mux.Unlock()

	if l.Cache[w.Domain] != 0 {
		l.Cache[w.Domain] = w.Delay
	} else {
		return errors.New(ErrNotFound)
	}

	return nil
}

func (l *LocalCache) Min() (w *cache.Website, err error) {
	var time time.Duration = time.Minute
	var d string

	l.mux.Lock()
	defer l.mux.Unlock()
	for domain, delay := range l.Cache {
		if time > delay {
			time = delay
			d = domain
		}
	}

	return cache.New(d, time), nil
}

func (l *LocalCache) Max() (w *cache.Website, err error) {
	var hai time.Duration = time.Second * time.Duration(config.Timeout)
	var time time.Duration = time.Nanosecond
	var d string

	l.mux.Lock()
	defer l.mux.Unlock()

	for domain, delay := range l.Cache {
		if time < delay && delay < hai {
			time = delay
			d = domain
		}
	}

	return cache.New(d, time), nil
}

func (l *LocalCache) List() (list []string, err error) {
	list = make([]string, 0, len(l.Cache))

	l.mux.Lock()
	defer l.mux.Unlock()

	for domain := range l.Cache {
		list = append(list, domain)
	}

	return list, nil
}
