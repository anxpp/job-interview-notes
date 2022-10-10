package main

import (
	"GoNotes/go/gokit/pkg/endpoint"
	"GoNotes/go/gokit/pkg/server"
	"GoNotes/go/gokit/pkg/service"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our napodate service
	srv := service.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射端点
	endpoints := endpoint.Endpoints{
		GetEndpoint:      endpoint.MakeGetEndpoint(srv),
		StatusEndpoint:   endpoint.MakeStatusEndpoint(srv),
		ValidateEndpoint: endpoint.MakeValidateEndpoint(srv),
	}

	// HTTP 传输
	go func() {
		log.Println("napodate is listening on port:", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
