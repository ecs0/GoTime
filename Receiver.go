package main
import (
    "fmt" 
    "net"  
)


func main() {

    in := make([]byte, 2048)
	
	addr := net.UDPAddr{
        Port: 1337,
        IP: net.ParseIP("127.0.0.1"),
    }
	
    lstn, err := net.ListenUDP("udp", &addr)
	
    if err != nil {
        fmt.Printf("Error %v\n", err)
    }
	
	//read and print incoming messages
	//till the end of time
	//or ctrl-c
    for {
        _, err := lstn.Read(in)
        fmt.Printf("Packet sent at: %s", in)
        if err !=  nil {
            fmt.Printf("Error  %v", err)
        }
    }
}

