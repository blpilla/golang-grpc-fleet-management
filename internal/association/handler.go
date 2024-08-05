package association

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) AssociateDriverToVehicle(w http.ResponseWriter, r *http.Request) {
	var assoc Association
	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.AssociateDriverToVehicle(assoc.DriverID, assoc.VehicleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAllAssociations(w http.ResponseWriter, r *http.Request) {
	associations, err := h.Service.GetAllAssociations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(associations)
}

func (h *Handler) GetDriversByVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID, err := strconv.Atoi(params["vehicle_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	driverIDs, err := h.Service.GetDriversByVehicle(vehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(driverIDs)
}

func (h *Handler) GetVehiclesByDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	driverID, err := strconv.Atoi(params["driver_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vehicleIDs, err := h.Service.GetVehiclesByDriver(driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehicleIDs)
}

func (h *Handler) DissociateDriverFromVehicle(w http.ResponseWriter, r *http.Request) {
	var assoc Association
	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.DissociateDriverFromVehicle(assoc.DriverID, assoc.VehicleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
