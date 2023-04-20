package requests

import (
	"encoding/json"
	"github.com/dl-nft-books/nonce-auth-svc/resources"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
)

type UpdateUserRequest struct {
	Address string
	resources.UpdateUserRequest
}

func NewUpdateUserRequest(r *http.Request) (*UpdateUserRequest, error) {
	var request UpdateUserRequest
	request.Address = chi.URLParam(r, "address")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal update user request")
	}
	return &request, nil
}
