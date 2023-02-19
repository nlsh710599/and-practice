package main

import (
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/nlsh710599/and-practice/internal/config"
	"github.com/nlsh710599/and-practice/internal/database"
	"github.com/nlsh710599/and-practice/internal/service"
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

	bncs := &service.BigNumberComputationService{RDS: rds}

	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterService(bncs, "")
	http.Handle("/rpc", server)

	if err := http.ListenAndServe("localhost:"+strconv.Itoa(config.Get().Port), nil); err != nil {
		log.Panicf("Failed to listen and serve: %v", err)
	}
}
