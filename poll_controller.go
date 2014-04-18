package main

import (
	"errors"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/url"
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
	//TODO remove all membership rows for this poll ID
	err := c.Conn.DeleteAllMembersOfPoll(id)
	if err != nil {
		return err
	}

	//Remove poll from database
	return c.Conn.PollTable.Delete(id)
}
