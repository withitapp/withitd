package cntrl

import (
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/url"
)

type UserController struct {
	Table *dbase.UserTable
}

func (c *UserController) Index() interface{} {
	return nil
}

func (c *UserController) Show(id int) (interface{}, error) {
	return &model.User{}, nil
}

func (c *UserController) Create(values url.Values) (int, error) {
	return 0, nil
}

func (c *UserController) Update(id int, values url.Values) error {
	return nil
}

func (c *UserController) Delete(id int) error {
	return nil
}