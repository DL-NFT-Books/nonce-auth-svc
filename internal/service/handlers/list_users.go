package handlers

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/data"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/requests"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetListUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetListUsersRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch user list request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	users, err := applyQFiltersUsers(helpers.DB(r).Users(), request).Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ape.Render(w, responses.NewListUserResponse(r, users, *request))
}

func applyQFiltersUsers(qUsers data.UsersQ, request *requests.GetListUsersRequest) data.UsersQ {
	if len(request.Id) > 0 {
		qUsers = qUsers.FilterByUserID(request.Id...)
	}
	if len(request.Address) > 0 {
		qUsers = qUsers.FilterByAddress(request.Address...)
	}
	if len(request.Name) > 0 {
		qUsers = qUsers.FilterByAddress(request.Name...)
	}

	qUsers = qUsers.Sort(request.Sorts)
	qUsers = qUsers.Page(request.OffsetPageParams)

	return qUsers
}
