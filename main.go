package main

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	_ = timestamppb.New(time.Now())
	_ = grpc.NewServer()
}
