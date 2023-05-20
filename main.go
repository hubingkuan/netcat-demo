package main

import (
	"io"
	"log"
	"net"
	"netcat-demo/client"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: nc host port")
	}
	host := os.Args[1]
	port := os.Args[2]

	//conn, err := net.Dial("tcp", host+":"+port)
	// 连接超时控制
	conn, err := net.DialTimeout("tcp", host+":"+port, time.Duration(timeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 读写超时的连接控制
	timeoutConn := client.NewTimeoutConn(conn, time.Duration(3)*time.Second, time.Duration(3)*time.Second)
	go func() {
		io.Copy(timeoutConn, os.Stdin)
	}()
	io.Copy(os.Stdout, timeoutConn)
}