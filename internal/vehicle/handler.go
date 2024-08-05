package vehicle

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fleet_management/proto"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle proto.CreateVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Service.Create(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetVehicleByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid vehicle ID", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.GetByID(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp == nil {
		http.Error(w, "vehicle not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle proto.UpdateVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Service.Update(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid vehicle ID", http.StatusBadRequest)
		return
	}

	err = h.Service.Delete(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
