package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-bench/proto"
	"grpc-bench/tools"
	"log"
	"net"
	"time"
)

type LoggerServer struct {
	Counter *tools.Counter
}

func (l *LoggerServer) SaveLog(ctx context.Context, req *proto.LogData) (*proto.LogRes, error) {
	//log.Printf("ID:%s\tPayload:%s\n", req.Id, req.Payload)
	l.Counter.Inc()
	return &proto.LogRes{}, nil
}

func main() {
	loggerServer := &LoggerServer{Counter: &tools.Counter{}}
	go func() {
		for {
			<-time.After(time.Second * 1)
			fmt.Printf("\rRPC: %v sec", loggerServer.Counter.Count())
			loggerServer.Counter.Clear()
		}
	}()

	lis, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterLoggerServer(grpcServer, loggerServer)
	grpcServer.Serve(lis)
}
