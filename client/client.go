package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-bench/proto"
	"grpc-bench/tools"
	"log"
	"time"
	"runtime"
)

func main() {
	counter := &tools.Counter{}
	go func() {
		for {
			<-time.After(time.Second * 1)
			fmt.Printf("\rRPC: %v sec", counter.Count())
			counter.Clear()
		}
	}()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:50000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewLoggerClient(conn)

	for {
		_, err = client.SaveLog(context.Background(), &proto.LogData{Id: "ID", Payload: "Payload"})
		if err != nil {
			log.Printf("%v.GetFeatures(_) = _, %v: \n", client, err)
		}
		counter.Inc()
		//<-time.After(time.Microsecond * 1)
		runtime.Gosched()
	}
}
