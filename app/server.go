package main

import (
	"fmt"
	"net"
	"os"
)

const PORT string = ":6379"

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
		return
	}
	_, err = l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	defer l.Close()
}
