package dbase

type Conn struct {
	UserTable *UserTable
	PollTable *PollTable
	MembershipTable *MembershipTable
	FriendshipTable *MembershipTable
}

func NewConn() (*Conn, error) {
	return &Conn{
		UserTable: &UserTable{},
		PollTable: &PollTable{},
		MembershipTable: &MembershipTable{},
		FriendshipTable: &MembershipTable{},
	}, nil
}
