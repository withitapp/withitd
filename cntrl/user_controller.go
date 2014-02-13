package cntrl

import (
	"errors"
	"github.com/withitapp/withitd/dbase"
	"net/url"
)

type UserController struct {
	Conn *dbase.Conn
}

func (c *UserController) Index() interface{} {
	return nil
}

func (c *UserController) Show(id int) (interface{}, error) {
	return c.Conn.UserTable.Select(id)
}

func (c *UserController) Create(values url.Values) (int, error) {
	return 0, errors.New("cntrl.UserController.Create(): not implemented")
}

func (c *UserController) Update(id int, values url.Values) error {
	return errors.New("cntrl.UserController.Update(): not implemented")
}

func (c *UserController) Delete(id int) error {
	return errors.New("cntrl.UserController.Delete(): not implemented")
}
