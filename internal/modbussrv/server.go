package modbussrv

import (
	"fmt"
	"sync"
	// imports pour le serveur modbus choisi
)

type Server struct {
	mu sync.RWMutex
	// underlying server instance (ex: *mbserver.Server)
	// memory for registers/coils
	holding []uint16
	coils   []bool
}

func NewServer(sizeHolding int, sizeCoils int) *Server {
	s := &Server{
		holding: make([]uint16, sizeHolding),
		coils:   make([]bool, sizeCoils),
	}
	return s
}

func (s *Server) Start(host string, port int) error {
	// bind serveur modbus TCP host:port
	// callbacks read/write to s.holding / s.coils
	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) ForceHolding(addr int, val uint16) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr >= 0 && addr < len(s.holding) {
		s.holding[addr] = val
	}
}

func (s *Server) ReleaseHolding(addr int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr >= 0 && addr < len(s.holding) {
		s.holding[addr] = 0
	}
}

func (s *Server) ReadHolding(addr int) (uint16, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if addr < 0 || addr >= len(s.holding) {
		return 0, fmt.Errorf("address %d out of range", addr)
	}

	return s.holding[addr], nil
}

func (s *Server) ReadAllHolding() []uint16 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]uint16, len(s.holding))
	copy(out, s.holding)
	return out
}

func (s *Server) Architecture() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return map[string]interface{}{
		"type":              "modbus-server",
		"holding_registers": len(s.holding),
		"coils":             len(s.coils),
	}
}
