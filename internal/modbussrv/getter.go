package modbussrv

import (
	"errors"
)

func (m *ModbusServer) GetCoil(address uint16) (bool, error) {
	if int(address) >= len(m.Server.Coils) {
		return false, errors.New("coil address out of range")
	}
	return m.Server.Coils[address] == 1, nil
}

func (m *ModbusServer) GetAllCoils() []bool {
	res := make([]bool, len(m.Server.Coils))
	for i, v := range m.Server.Coils {
		res[i] = v == 1
	}
	return res
}

func (m *ModbusServer) GetDiscreteInput(address uint16) (bool, error) {
	if int(address) >= len(m.Server.DiscreteInputs) {
		return false, errors.New("discrete input address out of range")
	}
	return m.Server.DiscreteInputs[address] == 1, nil
}

func (m *ModbusServer) GetAllDiscreteInputs() []bool {
	res := make([]bool, len(m.Server.DiscreteInputs))
	for i, v := range m.Server.DiscreteInputs {
		res[i] = v == 1
	}
	return res
}

func (m *ModbusServer) GetInputRegister(address uint16) (uint16, error) {
	if int(address) >= len(m.Server.InputRegisters) {
		return 0, errors.New("input register address out of range")
	}
	return m.Server.InputRegisters[address], nil
}

func (m *ModbusServer) GetAllInputRegisters() []uint16 {
	res := make([]uint16, len(m.Server.InputRegisters))
	copy(res, m.Server.InputRegisters)
	return res
}

func (m *ModbusServer) GetHoldingRegister(address uint16) (uint16, error) {
	if int(address) >= len(m.Server.HoldingRegisters) {
		return 0, errors.New("holding register address out of range")
	}
	return m.Server.HoldingRegisters[address], nil
}

func (m *ModbusServer) GetAllHoldingRegisters() []uint16 {
	res := make([]uint16, len(m.Server.HoldingRegisters))
	copy(res, m.Server.HoldingRegisters)
	return res
}
