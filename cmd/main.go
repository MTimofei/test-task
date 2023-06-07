package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	confi "githud.com/test-task/insert"
	"githud.com/test-task/insert/api"
	"githud.com/test-task/insert/kesh/lokalkash"
	"githud.com/test-task/insert/processor/ping"
)

func main() {
	flag.Parse()

	k := lokalkash.New(confi.Domains...)

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

	log.Fatal(http.ListenAndServe(*confi.Host, a.Routr()))
}
