package main

import (
	"authorization/app/istio"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	httpServer := istio.NewHttpAuthorizer(8080)
	grpcServer := istio.NewGrpcAuthorizer(9090)
	go httpServer.Start(&wg)
	go grpcServer.Start(&wg)
	defer httpServer.Stop()
	defer grpcServer.Stop()

	// Wait for the process to be shutdown.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
