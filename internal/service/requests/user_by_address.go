package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type UserByAddressRequest struct {
	Address string
}

func NewUserByAddressRequest(r *http.Request) (*UserByAddressRequest, error) {
	address := chi.URLParam(r, "address")

	return &UserByAddressRequest{
		Address: address,
	}, nil
}
