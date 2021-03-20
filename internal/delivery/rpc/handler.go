package rpc

import (
	"../../../pkg/logger"
	"../../service/basic"
	"../../service/pb"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	pb.UnimplementedBankTerminalArcusServer
}

func (s *Server) Run(ctx context.Context, in *pb.RunRequest) (*pb.RunResponse, error) {
	//
	result, _ := json.Marshal(basic.Run(in.Command))
	return &pb.RunResponse{Body: string(result)}, nil
}

//
func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.New("Error start RPC server")
		logger.New(fmt.Sprintf("port: %s", port))
		logger.New(fmt.Sprintf("error: %s", err.Error()))
		return
	}
	rpcServ := grpc.NewServer()

	pb.RegisterBankTerminalArcusServer(rpcServ, &Server{})

	reflection.Register(rpcServ)

	logger.New("RPC server is running on port", port)
	err = rpcServ.Serve(listener)
	if err != nil {
		logger.New("failed to serve")
		logger.New(err)
		return
	}
}
