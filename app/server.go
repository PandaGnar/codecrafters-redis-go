package main

import (
  "fmt"
  "net"
  "os"
)

const PORT string = ":6379"
const BUFFER_SIZE int = 512

func handleRequest(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, BUFFER_SIZE)
    for {
      n, err := conn.Read(buffer)

      if n == 0 {
        break
      }
      if err != nil {
        fmt.Println("Error reading from connection: ", err.Error())
      }

      conn.Write([]byte("$4\r\nPONG\r\n"))
    }
}

func main() {
  ln, err := net.Listen("tcp", PORT)
  if err != nil {
    fmt.Println("Error binding port: ", err.Error())
    os.Exit(1)
    return
  }
  defer ln.Close()

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error accepting connection: ", err.Error())
      continue
    }
    go handleRequest(conn)
  }
}
