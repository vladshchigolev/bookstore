package main

import (
	"bookstore/pkg/api"
	"bookstore/server"
	"flag"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/grpcserver.toml", "Путь к конфигурационному файлу")
}
func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	srv := server.New(config) // создаём инстанс нашего grpcserver'а, реализующего интерфейс BookStorageServer
	if err := srv.ConfigureLogger(); err != nil {
		log.Fatal(err)
	}
	if err := srv.ConfigureDatabase(); err != nil {
		log.Fatal(err)
	}
	api.RegisterBookStorageServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
