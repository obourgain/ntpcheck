package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	"time"
)

func main() {
	app := cli.App("ntpcheck", "Check timestamp of several severs to check the behavior of ntpd")

	servers := app.StringsOpt("s server", []string{}, "a server to contact, as a <host:port> pair")
	app.Version("v version", "ntpcheck 0.0.1")

	app.Action = func() {
		check(servers)
	}

	app.Run(os.Args)
}

func check(servers *[]string) {
	for _, server := range *servers {
		go doCheck(server)
	}
	time.Sleep(time.Duration(4) * time.Second)
}

func doCheck(server string) {
	url := fmt.Sprintf("http://%s/ntp", server)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err %s", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s : %s", server, body)
}
