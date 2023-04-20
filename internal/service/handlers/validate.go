package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/dl-nft-books/nonce-auth-svc/internal/service/helpers"
)

func Validate(w http.ResponseWriter, r *http.Request) {
	doorman := helpers.DoormanConnector(r)

	token, err := doorman.GetAuthToken(r)
	if err != nil {
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "invalid authorization header",
			Meta: &map[string]interface{}{
				"Authorization": "invalid authorization header",
			},
		})
		return
	}

	_, err = doorman.ValidateJwt(token)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
