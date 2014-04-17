package dbase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oguzbilgic/ninow"
	"github.com/withitapp/withitd/model"
)

type Conn struct {
	dbConn *sql.DB

	UserTable       *ninow.Table
	PollTable       *ninow.Table
	MembershipTable *ninow.Table
	FriendshipTable *ninow.Table
}

func NewConn(url string) (*Conn, error) {
	dbConn, err := sql.Open("mysql", url+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return &Conn{
		dbConn:          dbConn,
		UserTable:       ninow.TableFor(model.User{}, dbConn),
		PollTable:       ninow.TableFor(model.Poll{}, dbConn),
		MembershipTable: ninow.TableFor(model.Membership{}, dbConn),
		FriendshipTable: ninow.TableFor(model.Friendship{}, dbConn),
	}, nil
}

func (c *Conn) Close() {
	c.dbConn.Close()
}

func (c *Conn) SelectAllMembersOfPoll(pollID string) ([]*model.User, error) {
	query := "select users.* "
	query += "from users, memberships "
	query += "where " + pollID + "=memberships.poll_id AND memberships.user_id = users.id"

	members, err := c.UserTable.Query(query)
	if err != nil {
		return nil, err
	}

	return members.([]*model.User), nil
}
