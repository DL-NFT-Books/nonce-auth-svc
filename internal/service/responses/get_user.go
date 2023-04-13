package responses

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/data"
	"github.com/dl-nft-books/nonce-auth-svc/resources"
)

func NewGetUserResponse(user data.User) resources.UserResponse {
	return resources.UserResponse{
		Data: resources.User{
			Attributes: resources.UserAttributes{
				Address:   user.Address,
				CreatedAt: user.CreatedAt,
				Name:      user.Name,
			},
		},
	}
}
