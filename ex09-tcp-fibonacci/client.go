package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	var connectAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081}
	var res string

	conn, err := net.DialTCP("tcp", nil, connectAddr)
	if err != nil {
		fmt.Println("Cannot conect: ", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Fibonacci sequence: ")
	fibon, _, _ := reader.ReadLine()
	value, err := strconv.Atoi(string(fibon))
	if value < 0 {
		fmt.Println("Please, don't use negative values")
		return
	}
	if string(fibon) == "exit" {
		value = -404
	} else if err != nil {
		fmt.Println(err)
		return
	}
	encode := json.NewEncoder(conn)
	encode.Encode(value)
	decode := json.NewDecoder(conn)
	err1 := decode.Decode(&res)
	if err1 != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(res)
}
