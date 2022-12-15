package main

import (
	"../../internal/delivery/rpc"
	"../../internal/service/basic"
	"os"
	"time"
)

func main() {
	//go basic.ScanLog(time.Now())
	go func() {
		time.Sleep(time.Duration(30) * time.Second)

		basic.Run("/o1 /a285800")
	}()
	rpc.StartServer(os.Getenv("RPC_SERVER_PORT"))
}
