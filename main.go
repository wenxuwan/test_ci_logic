package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:34904", grpc.WithInsecure(), grpc.WithBlock())
	fmt.Printf("hello: %+v, err:%+v", conn, err)
}
