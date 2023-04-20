package pg

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/nonce-auth-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	usersTableName = "users"
	usersId        = "id"
	usersAddress   = "address"
	usersName      = "name"
	usersCreatedAt = "created_ad"
)

func newUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type usersQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
	sort       *pgdb.Sorts
}

func (q *usersQ) Get() (*data.User, error) {
	var result data.User
	stmt := q.sql.Select("*").From(usersTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, usersAddress)
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get address from db")
	}
	return &result, nil
}

func (q *usersQ) Select() ([]data.User, error) {
	var result []data.User
	stmt := q.sql.Select("*").From(usersTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, usersAddress)
	}
	if q.sort != nil {
		stmt = q.sort.ApplyTo(stmt, map[string]string{
			"id": "id",
		})
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select addresses from db")
	}
	return result, nil
}

func (q *usersQ) Insert(value data.User) (id int64, err error) {
	clauses := structs.Map(value)
	stmt := sq.Insert(usersTableName).SetMap(clauses).Suffix("returning id")
	err = q.db.Get(&id, stmt)
	return
}

func (q *usersQ) Update(name string) error {
	return q.db.Exec(q.sql.Update(usersTableName).Set(usersName, name))
}

func (q *usersQ) Delete() error {
	err := q.db.Exec(q.sql.Delete(usersTableName))
	if err != nil {
		return errors.Wrap(err, "failed to delete address from db")
	}
	return nil
}

func (q *usersQ) Sort(sort pgdb.Sorts) data.UsersQ {
	q.sort = &sort

	return q
}
func (q *usersQ) Page(pageParams pgdb.OffsetPageParams) data.UsersQ {
	q.pageParams = &pageParams
	return q
}

func (q *usersQ) FilterByName(name ...string) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{usersName: name})
	return q
}

func (q *usersQ) FilterByAddress(addresses ...string) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{usersAddress: addresses})
	return q
}

func (q *usersQ) SearchByAddress(address string) data.UsersQ {
	q.sql = q.sql.Where(sq.Like{usersAddress: fmt.Sprint(address, "%")})
	return q
}

func (q *usersQ) FilterByUserID(userIds ...int64) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{usersId: userIds})
	return q
}
