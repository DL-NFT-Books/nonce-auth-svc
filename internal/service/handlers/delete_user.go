package handlers

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

func DeleteUserByAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUserByAddressRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch user by id request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	user, err := helpers.DB(r).Users().FilterByAddress(request.Address).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if user == nil {
		helpers.Log(r).WithFields(logan.F{"address": request.Address}).Error("user with such address not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if err = helpers.DB(r).Users().FilterByAddress(request.Address).Delete(); err != nil {
		helpers.Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
