package main

import (
  "bufio"
  "fmt"
  "net"
)

func handleInput(c net.Conn) {
  b := bufio.NewReader(c)
  fmt.Println("Echoing..")
  for {
    line, err := b.ReadBytes('\n')
    if nil != err {
      break
    }

    c.Write(line)
  } 
  fmt.Println("Terminating connection..")
  c.Close()
}

func spawnServer() {
  l, err := net.Listen("tcp", ":8080")

  if nil != err {
    return  
  }

  for {
    c, err := l.Accept()

    if nil != err {
      fmt.Println("Error accepting connection..")
    } else {
      fmt.Println("Accepted connection")
      go handleInput(c)
    }
  }
}

func main() {
  fmt.Println("Spawning server..")
  spawnServer()
}
