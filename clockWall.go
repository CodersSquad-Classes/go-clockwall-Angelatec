package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*func outClock(conn net.Conn) {
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}*/

func main() {
	ch := make(chan int, 3)

	for i := 1; i <= len(os.Args)-1; i++ {
		input := os.Args[i]
		for j := 0; j <= len(input)-1; j++ {
			if input[j:j+1] == "=" {
				go func() {
					fmt.Printf("jasda", input[j+1:len(input)])
					conn, err := net.Dial("tcp", input[j+1:len(input)])
					if err != nil {
						fmt.Println("f")
						log.Fatal(err)
					}
					defer conn.Close()
					mustCopy(os.Stdout, conn)
					ch <- 2
				}()
				break
			}
		}

	}
	<-ch
}
