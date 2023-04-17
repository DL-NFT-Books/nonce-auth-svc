package requests

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/util"
	"github.com/dl-nft-books/nonce-auth-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type CreateUserRequest struct {
	resources.CreateUserRequest
}

func NewCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	var request CreateUserRequest

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal create user request")
	}
	return &request, request.validate()
}

func (r CreateUserRequest) validate() error {
	return validation.Errors{
		"/data/type": validation.Validate(&r.Data.Type, validation.Required, validation.In(resources.CREATE_USER)),
		"/data/attributes/address": validation.Validate(&r.Data.Attributes.Address,
			validation.Required,
			validation.Match(util.AddressRegexp)),
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.Name,
			validation.Required),
	}.Filter()
}
