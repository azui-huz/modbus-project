package modbussrv

type Architecture struct {
	HoldingRegisters int   `json:"holding_registers"`
	InputRegisters   int   `json:"input_registers"`
	Coils            int   `json:"coils"`
	DiscreteInputs   int   `json:"discrete_inputs"`
	UnitID           uint8 `json:"unit_id"`
}
