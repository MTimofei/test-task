package ping

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"githud.com/test-task/internal/cache"
	"githud.com/test-task/pkg/e"
)

const (
	ErrStart string = "can't start"
	ErrPing  string = "can't ping"
)

type Ping struct {
	client   *http.Client
	chDomain chan (string)
	wg       *sync.WaitGroup
	object   cache.Cache
}

func New(object cache.Cache) *Ping {
	return &Ping{
		client:   &http.Client{Timeout: 1 * time.Second},
		chDomain: make(chan string),
		wg:       &sync.WaitGroup{},
		object:   object,
	}
}

// запускает проверку сайтов
func (p *Ping) Start() (err error) {
	defer func() { err = e.IfErr(ErrStart, err) }()

	log.Println("ping start")

	var n int = runtime.NumCPU()
	var ch = make(chan string)

	p.chDomain = ch

	l, err := p.object.List()
	if err != nil {
		return err
	}

	p.wg.Add(1)
	go p.readList(l...)

	for i := 0; i < n; i++ {
		p.wg.Add(1)
		go p.ping()
	}

	p.wg.Wait()
	log.Println("ping stop")
	return
}

func (p *Ping) readList(domains ...string) {
	defer p.wg.Done()
	for _, domain := range domains {
		p.chDomain <- domain
	}
	close(p.chDomain)
}

// производит проверку сайтав
func (p *Ping) ping() {
	defer p.wg.Done()

	for domain := range p.chDomain {
		req, err := http.NewRequest(http.MethodGet, "https://www."+domain, nil)
		if err != nil {
			log.Println(e.Err(ErrPing, err))
			err = p.object.Update(cache.New(domain, 0))
			if err != nil {
				log.Println(e.Err(ErrPing, err))
				continue
			}
			continue
		}

		startTime := time.Now()

		res, err := p.client.Do(req)
		if err != nil {
			log.Println(e.Err(ErrPing, err))
			continue
		}

		res.Body.Close()

		responseTime := time.Since(startTime)

		err = p.object.Update(cache.New(domain, time.Duration(responseTime.Milliseconds())))
		if err != nil {
			log.Println(e.Err(ErrPing, err))
			continue
		}
	}
}
