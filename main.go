package main

import (
	"fmt"
	"google.golang.org/grpc"
	"time"
)

// TODO: 开源以后完善
func main() {
	conn, err := grpc.Dial("127.0.0.1:34904", grpc.WithInsecure(), grpc.WithBlock())
	fmt.Println("Hello github", conn, err)
}
