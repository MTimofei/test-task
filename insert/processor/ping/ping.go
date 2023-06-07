package ping

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"githud.com/test-task/insert/kesh"
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
	object   kesh.Kesh
}

func New(object kesh.Kesh) *Ping {
	return &Ping{
		client:   &http.Client{},
		chDomain: make(chan string),
		wg:       &sync.WaitGroup{},
		object:   object,
	}
}

func (p *Ping) Start() (err error) {
	defer func() { err = e.IfErr(ErrStart, err) }()

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
	return
}

func (p *Ping) readList(domens ...string) {
	defer p.wg.Done()
	for _, domen := range domens {
		p.chDomain <- domen
	}
	close(p.chDomain)
	p.wg.Done()
}

func (p *Ping) ping() {
	defer p.wg.Done()

	for domain := range p.chDomain {
		req, err := http.NewRequest(http.MethodGet, "https://www."+domain, nil)
		if err != nil {
			log.Println(e.Err(ErrPing, err))
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

		err = p.object.Updata(kesh.New(domain, responseTime))
		if err != nil {
			log.Println(e.Err(ErrPing, err))
			continue
		}
	}
}
