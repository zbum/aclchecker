package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func StartServer(protocol string, serverConfig *ServerConfig) {
	if serverConfig == nil {
		fmt.Printf("서버 설정을 불러오는 데 실패했습니다.\n")
		return
	}
	switch protocol {
	case "tcp":
		for _, port := range serverConfig.Server.Ports {
			go startTcpServer("", port)
		}
		for _, port := range serverConfig.Server.ResetPorts {
			go startTcpResetServer("", port)
		}
	case "udp":
		for _, port := range serverConfig.Server.Ports {
			go startUdpServer("", port)
		}
	}

}

func startTcpServer(ip string, port int) {
	listener, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("서버 시작 실패: %v \n", err)
		return
	}
	fmt.Printf("서버 시작: %v \n", listener.Addr())

	for {
		connection, err := listener.Accept()
		if err != nil {
			return
		}

		fmt.Printf("연결 성공: Remote: %v, Local: %v\n", connection.RemoteAddr(), connection.LocalAddr())

		go func() {
			defer connection.Close()

			for {

				buffer := make([]byte, 1000)

				read, err := connection.Read(buffer)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("연결 종료 : %v\n", err)
					} else {
						fmt.Printf("수신 실패 : %v\n", err)
					}
					return
				}

				fmt.Printf("수신 성공[%s]: %s\n", net.JoinHostPort(ip, strconv.Itoa(port)), string(buffer[:read]))
				_, err = connection.Write(buffer[:read])
			}
		}()
	}
}

func startTcpResetServer(ip string, port int) {
	listener, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("리셋 서버 시작 실패: %v \n", err)
		return
	}
	fmt.Printf("리셋 서버 시작: %v \n", listener.Addr())

	for {
		connection, err := listener.Accept()
		if err != nil {
			return
		}

		fmt.Printf("리셋 서버 연결 성공: Remote: %v, Local: %v\n", connection.RemoteAddr(), connection.LocalAddr())

		go func() {

			for {

				buffer := make([]byte, 1000)

				read, err := connection.Read(buffer)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("리셋 서버 연결 종료 : %v\n", err)
					} else {
						fmt.Printf("리셋 서버 수신 실패 : %v\n", err)
					}
					return
				}

				fmt.Printf("리셋 서버 수신 성공[%s]: %s\n", net.JoinHostPort(ip, strconv.Itoa(port)), string(buffer[:read]))
				connection.Close()
			}
		}()
	}
}

func startUdpServer(ip string, port int) {
	udpAddr, err := net.ResolveUDPAddr("udp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("서버 시작 실패: %v \n", err)
		return
	}
	connection, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		fmt.Printf("서버 시작 실패: %v \n", err)
		return
	}
	defer connection.Close()
	fmt.Printf("서버 시작: %v \n", udpAddr)

	for {
		buffer := make([]byte, 1000)

		read, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("수신 실패 : %v\n", err)
		}

		fmt.Printf("수신 성공 : %s\n", string(buffer[:read]))
		_, err = connection.WriteToUDP(buffer[:read], addr)
	}

}
