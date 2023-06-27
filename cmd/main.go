package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	config "githud.com/test-task/internal"
	"githud.com/test-task/internal/api"
	localcache "githud.com/test-task/internal/cache/localcache"
	"githud.com/test-task/internal/processor/ping"
)

func main() {
	flag.Parse()

	k := localcache.New(config.Domains...)

	go func() {
		for {
			process := ping.New(k)
			err := process.Start()
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Minute)
		}
	}()

	a := api.New(k)

	log.Fatal(http.ListenAndServe(*config.Host, a.Router()))
}
