package main

import (
	"errors"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/url"
	"strconv"
)

type PollController struct {
	Conn *dbase.Conn
}

func (c *PollController) Index() (interface{}, error) {
	polls, err := c.Conn.PollTable.SelectAllBy("", "")
	return polls.([]*model.Poll), err
}

func (c *PollController) Show(id int) (interface{}, error) {
	return c.Conn.PollTable.Select(id)
}

func (c *PollController) Create(values url.Values) (int, error) {
	poll, err := model.NewPollFromValues(values)
	if err != nil {
		panic(err)
	}

	pollID, err := c.Conn.PollTable.Insert(poll)
	if err != nil {
		panic(err)
	}

	return pollID, nil
}

func (c *PollController) Update(id int, values url.Values) error {
	return errors.New("cntrl.PollController.Update(): not implemented")
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
