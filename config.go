package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

type ServerConfig struct {
	Server struct {
		Ports      []int `toml:"ports"`
		ResetPorts []int `toml:"resetPorts"`
	}
}

func LoadServerConfig() *ServerConfig {
	fileBytes, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Printf("서버 설정 파일을 읽는 데 실패했습니다: %v\n", err)
		return nil
	}

	config := ServerConfig{}
	err = toml.Unmarshal(fileBytes, &config)
	if err != nil {
		fmt.Printf("서버 설정 파일을 읽는 데 실패했습니다: %v\n", err)
		return nil
	}
	return &config
}

type ClientConfig struct {
	Client struct {
		Addresses []string `toml:"addresses"`
	}
}

func LoadClientConfig() *ClientConfig {
	fileBytes, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Printf("클라이언트 설정 파일을 읽는 데 실패했습니다: %v\n", err)
		return nil
	}

	config := ClientConfig{}
	err = toml.Unmarshal(fileBytes, &config)
	if err != nil {
		fmt.Printf("클라이언트 설정 파일을 읽는 데 실패했습니다: %v\n", err)
		return nil
	}
	return &config
}
