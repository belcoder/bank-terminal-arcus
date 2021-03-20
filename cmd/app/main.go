package main

import (
	"../../internal/delivery/rpc"
	"os"
)

func main() {
	rpc.StartServer(os.Getenv("RPC_SERVER_PORT"))
}