package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	server, err := modbussrv.NewModbusServer(cfg)
	if err != nil {
		log.Fatal("Failed to create Modbus server:", err)
	}

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start Modbus server:", err)
	}
	log.Printf("Modbus TCP server running on %s:%d", server.Config.Server.Host, server.Config.Server.Port)

	// 3️⃣ Créer le handler API avec serveur + config
	apiSrv := api.NewServerAPI(server, cfg)

	// 4️⃣ Lier les routes à HTTP
	http.Handle("/", apiSrv.Routes())

	apiAddr := fmt.Sprintf(":%d", cfg.API.Port)
	log.Printf("HTTP API listening on %s", apiAddr)

	// 5️⃣ Démarrer le serveur HTTP dans un goroutine
	go func() {
		if err := http.ListenAndServe(apiAddr, nil); err != nil {
			log.Fatal(err)
		}
	}()

	// 6️⃣ Attendre CTRL+C pour fermer proprement
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Shutting down...")
	server.Close()
}
