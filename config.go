package main

import "os"

type Config struct {
	Mode         string
	GatewayPort  string
	ReaderUrl    string
	WriterUrl    string
	ReaderPort   string
	WriterPort   string
	DBConnection string
}

func NewConfig() Config {
	return Config{
		Mode:         os.Getenv("MODE"),
		GatewayPort:  os.Getenv("GATEWAY_PORT"),
		ReaderUrl:    os.Getenv("READER_URL"),
		WriterUrl:    os.Getenv("WRITER_URL"),
		ReaderPort:   os.Getenv("READER_PORT"),
		WriterPort:   os.Getenv("WRITER_PORT"),
		DBConnection: os.Getenv("DB_CONNECTION"),
	}
}

func (c Config) GetPort() string {
	if c.Mode == "reader" {
		return c.ReaderPort
	}
	if c.Mode == "writer" {
		return c.WriterPort
	}
	return c.GatewayPort
}
