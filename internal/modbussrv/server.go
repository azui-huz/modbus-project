package modbussrv

import (
	"fmt"
	"log"

	"github.com/tbrandon/mbserver"

	"modbus-project/internal/config"
)

type ModbusServer struct {
	Server *mbserver.Server
	Config *config.Config
}

// New : instancie un serveur Modbus + charge config + initialise registres
// func NewModbusServer(configPath string) (*ModbusServer, error) {
func NewModbusServer(cfg *config.Config) (*ModbusServer, error) {
	srv := mbserver.NewServer()

	if err := config.ApplyToServer(cfg, srv); err != nil {
		return nil, err
	}

	// Charger YAML et appliquer la config au serveur
	//cfg, err := config.LoadAndApply(configPath, srv)
	//if err != nil {
	//	return nil, fmt.Errorf("failed loading config: %w", err)
	//}

	return &ModbusServer{
		Server: srv,
		Config: cfg,
	}, nil
}

// Start : démarre le serveur TCP (host + port depuis config.yaml)
func (m *ModbusServer) Start() error {
	address := fmt.Sprintf("%s:%d", m.Config.Server.Host, m.Config.Server.Port)

	log.Printf("Starting Modbus TCP on %s (UnitID=%d)",
		address,
		m.Config.Server.UnitID,
	)

	if err := m.Server.ListenTCP(address); err != nil {
		return fmt.Errorf("cannot start Modbus TCP server: %w", err)
	}

	return nil
}

// Close : arrêt propre du serveur
func (m *ModbusServer) Close() {
	if m.Server != nil {
		m.Server.Close()
	}
}
