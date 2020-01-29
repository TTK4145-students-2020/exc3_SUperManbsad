package main
import (
  ."fmt"
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	buffer := make([]byte, 1024)
	tcpAddr, err := ResolveTCPAddr("tcp", ":30000")
	checkError(err)

	listener, err := ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
    Printf("sjekk om den går videre lizm\n")
		conn, err := listener.AcceptTCP()
		checkError(err)

		time.Sleep(1000*time.Millisecond)

		n,err := conn.Read(buffer)
		checkError(err)

		Printf("Received %d bytes: %s\n",n, buffer)

	}


}
