package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"modbus-project/internal/modbussrv"

	"github.com/go-chi/chi/v5"
)

type ServerAPI struct {
	srv *modbussrv.Server
}

func NewServerAPI(s *modbussrv.Server) *ServerAPI {
	return &ServerAPI{srv: s}
}

func (a *ServerAPI) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/api/force", a.Force)
	r.Post("/api/release", a.Release)
	r.Get("/api/read/holding/{addr}", a.ReadHolding)
	r.Get("/api/read/all", a.ReadAll)
	r.Get("/api/architecture", a.Architecture)
	return r
}

func (a *ServerAPI) Force(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Type  string `json:"type"`
		Addr  int    `json:"addr"`
		Value uint16 `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Type != "holding" {
		http.Error(w, "unsupported type", http.StatusBadRequest)
		return
	}

	a.srv.ForceHolding(body.Addr, body.Value)
	w.WriteHeader(http.StatusNoContent)
}

func (a *ServerAPI) Release(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Type string `json:"type"`
		Addr int    `json:"addr"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Type != "holding" {
		http.Error(w, "unsupported type", http.StatusBadRequest)
		return
	}

	a.srv.ReleaseHolding(body.Addr)
	w.WriteHeader(http.StatusNoContent)
}

func (a *ServerAPI) ReadHolding(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "addr")
	addr, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid addr", http.StatusBadRequest)
		return
	}

	v, err := a.srv.ReadHolding(addr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"addr":  addr,
		"value": v,
	})
}

func (a *ServerAPI) ReadAll(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"holding_registers": a.srv.ReadAllHolding(),
		"input_registers":   a.srv.ReadAllInputRegisters(),
		"coils":             a.srv.ReadAllCoils(),
		"discrete_inputs":   a.srv.ReadAllDiscreteInputs(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (a *ServerAPI) Architecture(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.srv.Architecture())
}
