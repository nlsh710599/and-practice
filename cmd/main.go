package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"

	"github.com/nlsh710599/and-practice/internal/config"
	"github.com/nlsh710599/and-practice/internal/database"
	"github.com/nlsh710599/and-practice/internal/method"
)

func main() {

	rds, err := database.New(config.Get().PostgresHost, config.Get().PostgresUser, config.Get().PostgresPassword,
		config.Get().PostgresDatabase, config.Get().PostgresPort)

	if err != nil {
		log.Panicf("Failed to initialize RDS: %v", err)
	}

	if err := rds.CreateTable(); err != nil {
		log.Panicf("Failed to create table: %v", err)
	}

	ctrl := &method.Controller{RDS: rds}
	rpc.Register(ctrl)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(config.Get().Port))
	if err != nil {
		log.Panicf("Failed to resolve TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Panicf("Failed to init tcp listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

}
