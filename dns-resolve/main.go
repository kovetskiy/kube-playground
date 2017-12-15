package main

import (
	"log"
	"net"
)

func main() {
	host := "tsdb-gateway"
	for {
		log.Println("looking for", host)
		addrs, err := net.LookupHost(host)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("%v", addrs)
	}
}
