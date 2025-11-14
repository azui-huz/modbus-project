package api

import (
	"encoding/json"
	"net/http"

	"modbus-project/internal/config"
	"modbus-project/internal/modbussrv"

	"github.com/go-chi/chi/v5"
)

// ServerAPI expose les endpoints HTTP
type ServerAPI struct {
	srv *modbussrv.ModbusServer
	cfg *config.Config
}

// Constructeur
func NewServerAPI(s *modbussrv.ModbusServer, cfg *config.Config) *ServerAPI {
	return &ServerAPI{
		srv: s,
		cfg: cfg,
	}
}

// Routes API
func (a *ServerAPI) Routes() http.Handler {
	r := chi.NewRouter()

	// GET /api/read/all  → lit toutes les tables Modbus
	r.Get("/api/read/all", a.ReadAll)

	return r
}

// Handler → lit tous les registres du serveur Modbus ou une table spécifique
func (a *ServerAPI) ReadAll(w http.ResponseWriter, r *http.Request) {
	// Paramètre optionnel "type"
	tableType := r.URL.Query().Get("type")

	data := make(map[string]interface{})

	switch tableType {
	case "":
		// Pas de paramètre → tout renvoyer
		data["holding_registers"] = a.srv.GetAllHoldingRegisters()
		data["input_registers"] = a.srv.GetAllInputRegisters()
		data["coils"] = a.srv.GetAllCoils()
		data["discrete_inputs"] = a.srv.GetAllDiscreteInputs()

	case "holding":
		data["holding_registers"] = a.srv.GetAllHoldingRegisters()

	case "input":
		data["input_registers"] = a.srv.GetAllInputRegisters()

	case "coils":
		data["coils"] = a.srv.GetAllCoils()

	case "discrete":
		data["discrete_inputs"] = a.srv.GetAllDiscreteInputs()

	default:
		http.Error(w, "Invalid type parameter. Use: holding, input, coils, discrete", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
