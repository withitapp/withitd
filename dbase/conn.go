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
	PollTable       *PollTable
	MembershipTable *MembershipTable
	FriendshipTable *MembershipTable
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
		PollTable:       &PollTable{dbConn},
		MembershipTable: &MembershipTable{dbConn},
		FriendshipTable: &MembershipTable{dbConn},
	}, nil
}

func (c *Conn) Close() {
	c.dbConn.Close()
}
