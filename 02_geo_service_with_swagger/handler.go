package main

import (
	"encoding/json"
	"fmt"
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

// @Summary AddressSearcb
// @Tags search
// @Description get address suggestions from address string
// @ID search
// @Produce json
// @Param input query string true "address string"
// @Success 200 {array} Address
// @Failture 400 {string}
// @Failture 500 {string}
// @Router /api/address/search [post]
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

// @Summary GeoCode
// @Tags geocode
// @Description get address suggestions from location
// @ID geocode
// @Produce json
// @Param lat query string true "latitude"
// @Param lng query string true "longitude"
// @Success 200 {array} Address
// @Failture 400 {string}
// @Failture 500 {string}
// @Router /api/address/geocode [post]
func (h *Handler) GeoCode(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	fmt.Println(lat, lng)
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
