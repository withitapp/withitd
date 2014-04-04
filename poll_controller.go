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
	return 0, errors.New("cntrl.PollController.Create(): not implemented")
}

func (c *PollController) Update(id int, values url.Values) error {
	return errors.New("cntrl.PollController.Update(): not implemented")
}

func (c *PollController) Delete(id int) error {
	return errors.New("cntrl.PollController.Delete(): not implemented")
}