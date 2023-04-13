package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type GetListUsersRequest struct {
	pgdb.OffsetPageParams
	Sorts pgdb.Sorts `url:"sort" default:"id"`

	Id      []int64  `filter:"id"`
	Address []string `filter:"address"`
	Name    []string `filter:"name"`
}

func NewGetListUsersRequest(r *http.Request) (*GetListUsersRequest, error) {
	var request GetListUsersRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, nil
}
