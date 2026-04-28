package main

import (
    "io"
    "net"
    "os"
)

func handleClient(clientConn net.Conn, targetAddr string) {
    defer clientConn.Close()

    remoteConn, _ := net.Dial("tcp", targetAddr)
    defer remoteConn.Close()

    go func() {
        io.Copy(remoteConn, clientConn)
    }()

    io.Copy(clientConn, remoteConn)
}

func main() {
    listenAddr := ":" + os.Getenv("8080")
    targetAddr := os.Getenv("45.61.163.95") + ":443"
    listener, _ := net.Listen("tcp", listenAddr)
    defer listener.Close()

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(clientConn, targetAddr)
    }
}
