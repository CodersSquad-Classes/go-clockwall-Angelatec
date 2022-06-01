// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
func handleConn(c net.Conn, name string) {
	defer c.Close()
	for {
		t, erro := TimeIn(time.Now(), name)
		if erro != nil {
			return // e.g., client disconnected
		}
		_, err := io.WriteString(c, name+" :"+t.Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing Parameters")
		os.Exit(-1)
	}
	name := os.Getenv("TZ")
	fmt.Println("Aux", name)
	port := "0.0.0.0:" + os.Args[2]
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, name) // handle connections concurrently
	}
}
