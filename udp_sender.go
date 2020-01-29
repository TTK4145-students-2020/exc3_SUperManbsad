package main
import (
  ."fmt"
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("whoopsi daisy!!! %v", err)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	udpAddr, err := ResolveUDPAddr("udp4", "10.100.23.255:20013")
	checkError(err)
	conn, err := DialUDP("udp4", nil, udpAddr)
	checkError(err)

  //defer conn.Close()
	for {
		time.Sleep(1000*time.Millisecond)

		// WRITE
		_, err := conn.Write([]byte("armon er best:))\n"))
		checkError(err)
	 }
}
//( ͡° ͜ʖ ͡°)
