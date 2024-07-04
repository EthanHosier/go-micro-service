package main

// For microservices, only need to use a different package for when types need to be shared between client and the running microservice
// Rest in root

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ethanhosier/go-micro-service/client"
)

func main() {
	client := client.New("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "RTCCC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", price)
	return
	listenAddr := flag.String("listenaddr", ":3000", "HTTP server listen address")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

}
