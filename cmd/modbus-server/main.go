package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"modbus-project/internal/api"
	"modbus-project/internal/config"
	"modbus-project/internal/modbussrv"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	srv := modbussrv.NewServer(
		cfg.Server.HoldingRegisters.Size,
		cfg.Server.Coils.Size,
		cfg.Server.DiscreteInputs.Size,
		cfg.Server.InputRegisters.Size,
	)

	if err := srv.Start(cfg.Server.Host, cfg.Server.Port); err != nil {
		log.Fatal("failed to start modbus server:", err)
	}
	defer srv.Stop()

	apiSrv := api.NewServerAPI(srv)
	http.Handle("/", apiSrv.Routes())

	apiAddr := fmt.Sprintf(":%d", cfg.API.Port)
	log.Println("API listening on", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr, nil))
}
