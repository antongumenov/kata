package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func main() {
	r := chi.NewRouter()
	s := NewGeoService("778f4730c1775ea32e9db73813d4cf5344254fa5", "af44d9bd2bd6b05e3fa8e1794f6d1b31f6e29626")
	h := NewHandler(s)

	r.Post("/api/address/search", h.AddressSearch)
	r.Post("/api/address/geocode", h.GeoCode)

	http.ListenAndServe(":8080", r)
}
