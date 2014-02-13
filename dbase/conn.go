package dbase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	dbConn *sql.DB

	UserTable       *UserTable
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
		UserTable:       &UserTable{dbConn},
		PollTable:       &PollTable{dbConn},
		MembershipTable: &MembershipTable{dbConn},
		FriendshipTable: &MembershipTable{dbConn},
	}, nil
}

func (c *Conn) Close() {
	c.dbConn.Close()
}
