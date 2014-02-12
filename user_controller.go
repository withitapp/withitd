package main

import (
	"net/url"
)

type UserController struct {
	Table *UserTable
}

func (c *UserController) Index() interface{} {
	return nil
}

func (c *UserController) Show(id int) (interface{}, error) {
	return &User{}, nil
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
