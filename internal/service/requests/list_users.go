package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
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

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal list users request")
	}

	return &request, nil
}
