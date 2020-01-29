package main
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return //os.exit(1)
	}

}

/*func read() {
	buffer := make([]byte, 1024)
	tcpAddr, err := ResolveTCPAddr("tcp", "10.100.23.255:20013")
	listener, err := DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//conn, err := listener.Accept()
	Println("Hei!")
	for {
		_,err := listener.Read(buffer)
		checkError(err)
		//Printf("Rcv %d bytes: %s\n",n, buffer)
		Printf(string(buffer))
	}
}*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// recvSock =
	//buffer := make([]byte, 1024)
	tcpAddr, err := ResolveTCPAddr("tcp", ":30000")
	checkError(err)
	conn, err := DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for {
		time.Sleep(1000*time.Millisecond)
		_, err := conn.Write([]byte("( ͡° ͜ʖ ͡°)\n\x00")) // \x00
		checkError(err)


	}





}
