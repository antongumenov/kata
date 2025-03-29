package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *GeoService
}

func NewHandler(service *GeoService) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) AddressSearch(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	if input == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Reguest\n"))
		return
	}
	addresses, err := h.service.AddressSearch(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error\n"))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addresses)

}

func (h *Handler) GeoCode(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	if lat == "" || lng == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Reguest\n"))
		return
	}

	addresses, err := h.service.GeoCode(lat, lng)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error\n"))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addresses)
}
