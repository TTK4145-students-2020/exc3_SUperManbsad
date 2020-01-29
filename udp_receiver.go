package main
import (
    ."fmt"
    ."net"
    "time"
    "runtime"
    "reflect"
    "unsafe"
    "bytes"
)

func checkError(err error) {
	if err != nil {
		Println("whoopsi daisy!!! fuk. %v", err)
		return
	}
}

func BytesToString(b []byte) string {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := reflect.StringHeader{bh.Data, bh.Len}
    return *(*string)(unsafe.Pointer(&sh))
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    buffer := make([]byte, 1024)
    ServerAddr,err := ResolveUDPAddr("udp", ":30000")
    checkError(err)
    l, err := ListenUDP("udp", ServerAddr)
    checkError(err)

    defer l.Close() //idunno?

    for {
        time.Sleep(1000*time.Millisecond)
        l.Read(buffer[:])
        checkError(err)
        str1 := bytes.NewBuffer(buffer).String()
        Println("hva er buffærn forno...?", str1)
    }
}
