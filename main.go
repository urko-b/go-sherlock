package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"go-sherlock/config"
	"go-sherlock/pkg"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no username input")
	}
	usernames := os.Args[1:]
	var WebS config.WebSites
	err := config.ParseSites(&WebS)
	if err != nil {
		fmt.Println(err.Error())
	}
	var work sync.WaitGroup
	for _, username := range usernames {
		work.Add(len(WebS))
		go pkg.Worker(WebS, username, &work)
	}
	work.Wait()
}
