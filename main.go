package main

import (
    "io"
    "log"
    "net"
    "os"
    "sync"
)

func handleClient(clientConn net.Conn, targetAddr string) {
    defer clientConn.Close()

    remoteConn, err := net.Dial("tcp", targetAddr)
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer remoteConn.Close()

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        io.Copy(remoteConn, clientConn)
    }()

    go func() {
        defer wg.Done()
        io.Copy(clientConn, remoteConn)
    }()

    wg.Wait()
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    targetHost := os.Getenv("TARGET_HOST")
    if targetHost == "" {
        targetHost = "45.61.163.95"
    }

    listenAddr := ":" + port
    targetAddr := targetHost + ":80"

    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal("listen error:", err)
    }
    defer listener.Close()

    log.Println("Listening on", listenAddr, "→", targetAddr)

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(clientConn, targetAddr)
    }
}
