package main

import (
	"../../internal/delivery/rpc"
	"os"
)

func main() {
	//go basic.ScanLog(time.Now())
	rpc.StartServer(os.Getenv("RPC_SERVER_PORT"))
}
