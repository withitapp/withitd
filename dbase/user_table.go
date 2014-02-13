package dbase

import (
	"database/sql"
	"errors"
	"github.com/withitapp/withitd/model"
)

const (
	UserTableName    string = "users"
	UserTableColumns        = "ID, created_at, updated_at, username, email, first_name, last_name, facebook_id"
)

type UserTable struct {
	db *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
	return &UserTable{db}
}

func (t *UserTable) Select(id int) (*model.User, error) {
	row := t.db.QueryRow("SELECT "+UserTableColumns+" FROM "+UserTableName+" WHERE id=?", id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.FacebookID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// BUG Implement
func (t *UserTable) SelectAll(id []int) (*[]model.User, error) {
	return nil, errors.New("dbase.UserTable.SelectAll(): not implemented")
}

// BUG Implement
func (t *UserTable) Insert(user *model.User) (int, error) {
	return 0, errors.New("dbase.UserTable.Insert(): not implemented")
}

// BUG Implement
func (t *UserTable) Update(user *model.User) error {
	return errors.New("dbase.UserTable.Update(): not implemented")
}

// BUG Implement
func (t *UserTable) Remove(user *model.User) error {
	return errors.New("dbase.UserTable.Remove(): not implemented")
}
