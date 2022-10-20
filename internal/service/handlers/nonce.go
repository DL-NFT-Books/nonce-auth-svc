package handlers

import (
	"net/http"
	"time"

	nonces "github.com/LarryBattle/nonce-golang"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/data"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/models"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/util"
)

func GetNonce(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewNonceRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, apierrors.BadRequest(apierrors.CodeBadRequestData, err))
		return
	}
	address := request.Data.Attributes.Address
	termsHash := request.Data.Attributes.TermsHash
	db := helpers.DB(r)

	expireTime := time.Now().UTC().Add(helpers.ServiceConfig(r).NonceExpireTime)
	nonceToken := nonces.NewToken()

	var message string = util.NonceToMessage(nonceToken)

	if termsHash != nil {
		message = util.NonceToTermsMessage(nonceToken, *termsHash)
	}

	nonce := data.Nonce{
		Message: nonceToken,
		Expires: expireTime.Unix(),
		Address: address,
	}

	// Required to make sure we have a clean `insert`-able state, as we're racing with nonce cleaner here
	err = db.Nonce().FilterByAddress(address).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, apierrors.InternalError(err))
		return
	}
	_, err = db.Nonce().Insert(nonce)
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, apierrors.InternalError(err))
		return
	}

	ape.Render(w, models.NewNonceModel(message))
}
