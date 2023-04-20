package handlers

import (
	"github.com/dl-nft-books/nonce-auth-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/nonce-auth-svc/solidity/generated/rolemanager"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"

	"github.com/dl-nft-books/nonce-auth-svc/internal/service/errors/apierrors"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
	"github.com/dl-nft-books/nonce-auth-svc/internal/service/requests"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewAdminLoginRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, apierrors.BadRequest(apierrors.CodeBadRequestData, err))
		return
	}
	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage
	apiErr, err := helpers.ValidateNonce(ethAddress, signature, r)
	if err != nil {
		logger.WithError(err).Debug("failed to validate nonce")
		ape.RenderErr(w, apiErr)
		return
	}
	networker := helpers.NetworkConnector(r)
	networks, err := networker.GetNetworksDetailed()
	if err != nil {
		logger.WithError(err).Debug("failed to get networks")
		ape.RenderErr(w, apierrors.InternalError())
		return
	}
	isAdmin := false
	for _, network := range networks.Data {
		contractsRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
		if err != nil {
			logger.WithError(err).Debug("failed to create contract registry")
			ape.RenderErr(w, apierrors.InternalError())
			return
		}
		roleManagerContract, err := contractsRegistry.GetRoleManagerContract(nil)
		roleManager, err := rolemanager.NewRolemanager(roleManagerContract, network.RpcUrl)
		if err != nil {
			logger.WithError(err).Debug("failed to create role manager")
			ape.RenderErr(w, apierrors.InternalError())
			return
		}
		isAdmin, err = roleManager.RolemanagerCaller.HasAnyRole(nil, common.HexToAddress(ethAddress))
		if err != nil {
			logger.WithError(err).Debug("failed to check is admin")
			ape.RenderErr(w, apierrors.InternalError())
			return
		}
		// check if user is admin at least in one network
		if isAdmin {
			break
		}
	}

	if !isAdmin {
		logger.Debug("not admin's address")
		ape.RenderErr(w, apierrors.Forbidden(apierrors.CodeAdminNotFound))
		return
	}
	// success logic
	doorman := helpers.DoormanConnector(r)
	pair, err := doorman.GenerateJwtPair(ethAddress, "session")
	if err != nil {
		logger.WithError(err).Error("failed to generate jwt")
		ape.RenderErr(w, apierrors.InternalError(err))
		return
	}

	ape.Render(w, pair)
}
