package main

import (
	"encoding/binary"
	. "fmt"
	"net"
	t "time"
)

var counter uint64
var buf = make([]byte, 16)

func main() {
	/* server */
	addr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	pr := false
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		Println("Error:", err)
	}

	/* Backup */
	Println("Backup!")
	for !pr {
		conn.SetReadDeadline(t.Now().Add(2 * t.Second))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			pr = true
		} else {
			counter = binary.BigEndian.Uint64(buf[:n])
		}
	}
	conn.Close()

	/* primary */

	Println("Primary!")
	bcastConn, _ := net.DialUDP("udp", nil, addr)
	for {

		Println(counter)

		counter++
		binary.BigEndian.PutUint64(buf, counter)

		_, _ = bcastConn.Write(buf)
		t.Sleep(t.Second)
	}

}
