package socket

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	StopCharacter = ""
)

func SocketClient(ip string, port int, message string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

}
