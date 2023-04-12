package models

import "github.com/dl-nft-books/nonce-auth-svc/resources"

func NewNonceModel(message string) resources.AuthNonceResponse {
	response := resources.AuthNonceResponse{
		Data: resources.AuthNonce{
			Key: resources.Key{Type: resources.AUTH_NONCE_MESSAGE},
			Attributes: resources.AuthNonceAttributes{
				Message: message,
			},
		},
	}
	return response
}
