package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	app := cli.App("ntpcheck_server", "Listens to ntpcheck request")

	port := app.IntOpt("p port", 8080, "the port to listen on")
	app.Version("v version", "ntpcheck_server 0.0.1")

	app.Action = func() {
		server(*port)
	}

	app.Run(os.Args)
}

func server(port int) {
	log.Printf("listening of %d", port)
	http.HandleFunc("/ntp", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received request %s", r)
		fmt.Fprintf(w, "%d", time.Now().Unix())
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
