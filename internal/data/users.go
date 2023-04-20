package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type UsersQ interface {
	Get() (*User, error)
	Select() ([]User, error)
	Insert(value User) (int64, error)
	Update(name string) error
	Delete() error

	Sort(sort pgdb.Sorts) UsersQ
	Page(pageParams pgdb.OffsetPageParams) UsersQ
	FilterByName(name ...string) UsersQ
	FilterByAddress(addresses ...string) UsersQ
	SearchByAddress(address string) UsersQ
	FilterByUserID(userIds ...int64) UsersQ
}

type User struct {
	ID        int64     `db:"id" structs:"-"`
	Name      string    `db:"name" structs:"name"`
	Address   string    `db:"address" structs:"address"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
}
