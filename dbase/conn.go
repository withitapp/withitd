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
	dbConn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	return &Conn{
		dbConn:          dbConn,
		UserTable:       &UserTable{},
		PollTable:       &PollTable{},
		MembershipTable: &MembershipTable{},
		FriendshipTable: &MembershipTable{},
	}, nil
}
