package pg

import (
	"artarn/gentree/domain/user"
	"github.com/jmoiron/sqlx"
)

type pgUserRepository struct {
	DB *sqlx.DB
}

func NewPGUserRepository(connection *sqlx.DB) user.Repository {
	return &pgUserRepository{
		DB: connection,
	}
}

func (p pgUserRepository) Store(user user.User) error {
	panic("implement me")
}

func (p pgUserRepository) FindById(id int) (user.User, error) {
	var u user.User
	err := p.DB.Get(&u, "select id, name from general.users where id  = $1", id)

	return u, err
}
