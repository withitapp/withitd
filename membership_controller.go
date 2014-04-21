package main

import (
	"errors"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/url"
)

type MembershipController struct {
	Conn *dbase.Conn
}

func (c *MembershipController) Index(values url.Values) (interface{}, error) {
	return c.Conn.MembershipTable.SelectAllBy("poll_id", values.Get("poll_id"))
}

func (c *MembershipController) Show(id int) (interface{}, error) {
	return nil, errors.New("cntrl.MemberhsipController.Show(): not implemented")
}

func (c *MembershipController) Create(values url.Values) (int, error) {
	membership, err := model.NewMembershipFromValues(values)
	if err != nil {
		panic(err)
	}

	membershipID, err := c.Conn.MembershipTable.Insert(membership)
	if err != nil {
		panic(err)
	}

	return membershipID, nil
}

func (c *MembershipController) Update(id int, values url.Values) error {
	return errors.New("cntrl.MembershipController.Update(): not implemented")
}

func (c *MembershipController) Delete(id int) error {
	return c.Conn.MembershipTable.Delete(id)
}
