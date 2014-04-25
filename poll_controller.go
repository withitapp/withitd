package main

import (
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/url"
	"strconv"
)

type PollController struct {
	Conn *dbase.Conn
}

func (c *PollController) Index(values url.Values) (interface{}, error) {
	polls, err := c.Conn.PollTable.SelectAllBy("", "")
	return polls.([]*model.Poll), err
}

func (c *PollController) Show(id int) (interface{}, error) {
	return c.Conn.PollTable.Select(id)
}

func (c *PollController) Create(values url.Values) (int, error) {
	poll, err := model.NewPollFromValues(values)
	if err != nil {
		return 0, err
	}

	pollID, err := c.Conn.PollTable.Insert(poll)
	if err != nil {
		return 0, err
	}

	return pollID, nil
}

func (c *PollController) Update(id int, values url.Values) error {
	poll, err := c.Conn.PollTable.Select(id)
	if err != nil {
		return err
	}

	err = poll.(*model.Poll).UpdateFromValues(values)
	if err != nil {
		return err
	}

	return c.Conn.PollTable.Update(poll.(*model.Poll))
}

func (c *PollController) Delete(id int) error {
	//Convert ID numbers to string
	pollID := strconv.Itoa(id)

	//select all memberships based on pollID
	memberships, err := c.Conn.MembershipTable.SelectAllBy("poll_id", pollID)
	if err != nil {
		panic(err)
		return err
	}

	//iterate through all memberships to delete
	for _, membership := range memberships.([]*model.Membership) {
		err := c.Conn.MembershipTable.Delete(membership.ID)
		if err != nil {
			return err
		}
	}

	return c.Conn.PollTable.Delete(id)
}
