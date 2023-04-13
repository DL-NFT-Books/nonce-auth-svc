package handlers

import (
	"github.com/dl-nft-books/nonce-auth-svc/internal/data"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/requests"
	"github.com/dl-nft-books/nonce-auth-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateUserRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create user request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	user, err := helpers.DB(r).Users().FilterByAddress(request.Data.Attributes.Address).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if user != nil {
		helpers.Log(r).WithFields(logan.F{"address": user.Address}).Error("user with such address is already exists")
		ape.RenderErr(w, problems.Conflict())
		return
	}
	id, err := helpers.DB(r).Users().Insert(data.User{
		Name:      request.Data.Attributes.Name,
		Address:   request.Data.Attributes.Address,
		CreatedAt: time.Now(),
	})
	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(id, resources.CREATE_USER),
	})
}
