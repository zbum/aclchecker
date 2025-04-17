package main

import (
	"aclchecker/report"
	"fmt"
	"net"
	"time"
)

const (
	message = `GET / HTTP/1.1
Host: localhost
User-Agent: aclchecker
Accept: */*

`
)

func StartClient(protocol string, clientConfig *ClientConfig) {
	channel := report.BuildChannel()
	switch protocol {
	case "tcp":
		for _, address := range clientConfig.Client.Addresses {
			startTcpClient(address, channel)
		}

		report.PrettyPrintReport()
	case "udp":
		for _, address := range clientConfig.Client.Addresses {
			startUdpClient(address, channel)
		}
	}
}

func startTcpClient(address string, channel chan report.Report) {
	connection, err := net.DialTimeout("tcp", address, 3*time.Second)

	if err != nil {
		fmt.Printf("연결 실패 : %v\n", err)
		channel <- report.Fail(address, message, err.Error())
		return
	} else {
		defer connection.Close()
		fmt.Printf("연결 성공 : %s \n", address)

		_, err := connection.Write([]byte(message))
		if err != nil {
			fmt.Printf("송신 실패 : %v\n", err)
			channel <- report.Fail(address, message, err.Error())
			return
		} else {
			fmt.Printf("송신 성공 : %v\n", message)
		}

		buffer := make([]byte, 50)
		read, err := connection.Read(buffer)
		if err != nil {
			channel <- report.Fail(address, message, err.Error())
			return
		}
		fmt.Printf("수신 성공 : %s\n", string(buffer[:read]))

		channel <- report.Success(address, message, string(buffer[:read]))

	}
}

func startUdpClient(address string, channel chan report.Report) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Printf("클라이언트 시작 실패: %v \n", err)
		return
	}
	connection, err := net.DialUDP("udp", nil, udpAddr)
	defer connection.Close()
	if err != nil {
		fmt.Printf("연결 실패 : %v\n", err)
		return
	} else {
		fmt.Printf("연결 성공 : \n")

		_, err := connection.Write([]byte(message))
		if err != nil {
			fmt.Printf("송신 실패 : %v\n", err)
			return
		} else {
			fmt.Printf("송신 성공 : %v\n", message)
		}

		buffer := make([]byte, 1000)
		read, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			return
		}

		fmt.Printf("수신 성공 : %v\n", addr)
		fmt.Printf("수신 성공 : %s\n", string(buffer[:read]))
	}
}
