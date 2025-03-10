package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", ":33546")
  for {
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\000")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}
