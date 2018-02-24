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
	//log.Printf("%s\t", req.Id)
	l.Counter.Inc()
	return &proto.LogRes{}, nil
}

func main() {
	loggerServer := &LoggerServer{Counter: &tools.Counter{}}
	go func() {
		for {
			<-time.After(time.Second * 1)
			fmt.Printf("RPC: %v sec\n", loggerServer.Counter.Count())
			loggerServer.Counter.Clear()
		}
	}()

	serverAddress := ":50000"
	log.Printf("Server adsress [%s]", serverAddress)
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterLoggerServer(grpcServer, loggerServer)
	grpcServer.Serve(lis)
}
