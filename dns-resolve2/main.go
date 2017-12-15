package main

import (
	"log"
	"net"
	"sync"
)

func main() {
	hosts := []string{
		"tsdb-gateway",
		"series-resampler",
		"metrics-aggregator",
	}
	for {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		for _, host := range hosts {
			go func(host string) {
				defer wg.Done()
				log.Println("looking for", host)
				addrs, err := net.LookupHost(host)
				if err != nil {
					log.Println(err)
					return
				}

				log.Printf("%v", addrs)
			}(host)
		}

		wg.Wait()
	}
}
