package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"
)

func fibonacci(i int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	for k := 0; k < i; k++ {
		a, b = b, a.Add(a, b)
	}
	return a
}

func startFib(mes int) string {
	var value string
	cache := make(map[int]*big.Int)

	if val, log := cache[mes]; log == false {
		t := time.Now()
		res := fibonacci(mes)
		tNew := time.Since(t)
		cache[mes] = res
		value = tNew.String() + " " + res.String()
	} else {
		t := time.Now()
		res := val
		tNew := time.Since(t)
		value = tNew.String() + " " + res.String()
	}
	return value
}

func main() {
	var listenAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}
	var numb int

	fmt.Println("Server is runnig...")
	ln, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		fmt.Println("Failed to initialize server: ", err)
		return
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Println("Cannot get connection: ", err)
		}
		message := json.NewDecoder(conn)
		message.Decode(&numb)
		if numb == -404 {
			conn.Close()
			return
		}
		str := startFib(numb)
		encode := json.NewEncoder(conn)
		erro := encode.Encode(str)
		if erro != nil {
			fmt.Println("Error: ", err)
		}
	}
}
