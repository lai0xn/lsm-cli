package utils

import (
	"log"
	"net"
)
//Retrieve the ip 
//We need this for the qr code

func GetIP() string{
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}
