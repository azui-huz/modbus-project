package modbussrv

import (
	"fmt"
	"sync"
)

type Server struct {
	mu sync.RWMutex

	// Modbus memory maps
	holding        []uint16 // 4xxxx
	inputRegisters []uint16 // 3xxxx
	coils          []bool   // 0xxxx
	inputs         []bool   // 1xxxx

	unitID uint8
}

// NewServer creates a new Modbus server with specified sizes for each memory area
func NewServer(sizeHolding, sizeCoils, sizeInputs, sizeInputRegs int) *Server {
	s := &Server{
		holding:        make([]uint16, sizeHolding),
		inputRegisters: make([]uint16, sizeInputRegs),
		coils:          make([]bool, sizeCoils),
		inputs:         make([]bool, sizeInputs),
	}
	return s
}

// Start starts the Modbus server (to be implemented with a Modbus library)
func (s *Server) Start(host string, port int) error {
	// TODO: integrate with actual Modbus library
	// Example: github.com/simonvetter/modbus or github.com/goburrow/modbus
	fmt.Printf("Starting Modbus server on %s:%d (Unit ID %d)\n", host, port, s.unitID)
	return nil
}

// Stop stops the Modbus server
func (s *Server) Stop() error {
	fmt.Println("Stopping Modbus server...")
	return nil
}

// ======================
// HOLDING REGISTERS (4x)
// ======================

func (s *Server) ForceHolding(addr int, val uint16) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr < 0 || addr >= len(s.holding) {
		return fmt.Errorf("holding register address %d out of range", addr)
	}
	s.holding[addr] = val
	return nil
}

func (s *Server) ReleaseHolding(addr int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr < 0 || addr >= len(s.holding) {
		return fmt.Errorf("holding register address %d out of range", addr)
	}
	s.holding[addr] = 0
	return nil
}

func (s *Server) ReadHolding(addr int) (uint16, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if addr < 0 || addr >= len(s.holding) {
		return 0, fmt.Errorf("holding register address %d out of range", addr)
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

// ======================
// INPUT REGISTERS (3x)
// ======================

func (s *Server) ReadInputRegister(addr int) (uint16, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if addr < 0 || addr >= len(s.inputRegisters) {
		return 0, fmt.Errorf("input register address %d out of range", addr)
	}
	return s.inputRegisters[addr], nil
}

func (s *Server) WriteInputRegister(addr int, val uint16) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr < 0 || addr >= len(s.inputRegisters) {
		return fmt.Errorf("input register address %d out of range", addr)
	}
	s.inputRegisters[addr] = val
	return nil
}

func (s *Server) ReadAllInputRegisters() []uint16 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]uint16, len(s.inputRegisters))
	copy(out, s.inputRegisters)
	return out
}

// ======================
// COILS (0x)
// ======================

func (s *Server) ReadCoil(addr int) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if addr < 0 || addr >= len(s.coils) {
		return false, fmt.Errorf("coil address %d out of range", addr)
	}
	return s.coils[addr], nil
}

func (s *Server) WriteCoil(addr int, val bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr < 0 || addr >= len(s.coils) {
		return fmt.Errorf("coil address %d out of range", addr)
	}
	s.coils[addr] = val
	return nil
}

func (s *Server) ReadAllCoils() []bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]bool, len(s.coils))
	copy(out, s.coils)
	return out
}

// ======================
// DISCRETE INPUTS (1x)
// ======================

func (s *Server) ReadInput(addr int) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if addr < 0 || addr >= len(s.inputs) {
		return false, fmt.Errorf("input address %d out of range", addr)
	}
	return s.inputs[addr], nil
}

func (s *Server) WriteInput(addr int, val bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if addr < 0 || addr >= len(s.inputs) {
		return fmt.Errorf("input address %d out of range", addr)
	}
	s.inputs[addr] = val
	return nil
}

func (s *Server) ReadAllInputs() []bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]bool, len(s.inputs))
	copy(out, s.inputs)
	return out
}

// ======================
// ARCHITECTURE
// ======================

func (s *Server) Architecture() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return map[string]interface{}{
		"unit_id":           s.unitID,
		"holding_registers": len(s.holding),
		"input_registers":   len(s.inputRegisters),
		"coils":             len(s.coils),
		"discrete_inputs":   len(s.inputs),
	}
}
