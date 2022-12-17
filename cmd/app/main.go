package main

import (
	"../../internal/delivery/rpc"
	"../../internal/service/rabbit"
	"context"
	"log"
	"os"
	"time"
)

func main() {
	rabbit.Context, _ = context.WithTimeout(context.Background(), 5*time.Second)

	// соединяемся с сервером AMPQ
	if rabbit.Connect() {
		// открытие канала AMPQ
		if rabbit.OpenChanel() {
			// регистрация очереди
			rabbit.QueueDeclare("TERMINAL_EVENTS")
		}
	}

	//go basic.ScanLog(time.Now())
	go func() {
		//if false {
		//	time.Sleep(time.Duration(30) * time.Second)
		//
		//	basic.Run("/o1 /a285800")
		//}

		time.Sleep(time.Duration(30) * time.Second)
		rabbit.PublishWithContext("TERMINAL_EVENTS", "TEST")
	}()
	rpc.StartServer(os.Getenv("RPC_SERVER_PORT"))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
