package main

import (
	"mbooke/grpc-server/grpcmain"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		grpcmain.Grpcmain()
		wg.Done()
	}()
	go func() {
		grpcmain.Streammain()
		wg.Done()
	}()

	wg.Wait()
}
