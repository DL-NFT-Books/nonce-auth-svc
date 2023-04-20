package responses

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/data"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/requests"
	"github.com/dl-nft-books/nonce-auth-svc/resources"
	"net/http"
)

func NewListUserResponse(r *http.Request, users []data.User, request requests.GetListUsersRequest) resources.UserListResponse {
	response := resources.UserListResponse{}

	if len(users) == 0 {
		return resources.UserListResponse{
			Data: make([]resources.User, 0),
		}
	}

	for _, user := range users {
		response.Data = append(response.Data, resources.User{
			Attributes: resources.UserAttributes{
				Address:   user.Address,
				CreatedAt: user.CreatedAt,
				Name:      user.Name,
			},
		})
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return response

}
