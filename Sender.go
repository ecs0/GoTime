package main
import (
    "fmt"
    "net"
	"time"
)

func main() {

	//Make the socket
    conn, err := net.Dial("udp", "127.0.0.1:1337")
	
    if err != nil {
        fmt.Printf("Error %v", err)
    }
	
	time := time.Now()
	
	//send time string through socket
    fmt.Fprintf(conn, time.String())
	
    conn.Close()
}