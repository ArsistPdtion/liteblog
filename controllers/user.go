package controllers

import (
	"github.com/ArsistPdtion/liteblog/models"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {
	//email can't null
	email := c.GetString("email", "email can't as null")
	//password can't null
	pwd := c.GetString("password", "password can't as null")
	var (
		user models.User
		err  error
	)
	if user, err = models.QueryUserByEmailAndPassword(email, pwd); err != nil {
		c.Abort("email or pwd error")
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOK("login successful", "/")
}

// @router /reg [post]
func (c *UserController) Reg() {
	name := c.GetString("name", "name can't null")
	email := c.GetString("email", "email can't null")
	pwd1 := c.GetString("password", "password can't null")
	pwd2 := c.GetString("password2", "confirm password can't null")
	if strings.Compare(pwd1, pwd2) != 0 {
		c.Abort("password must same")
	}
	if u, err := models.QueryUserByName(name); err == nil && u.ID != 0 {
		c.Abort("user name existed.")
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID != 0 {
		c.Abort("email existed.")
	}
	//start save user
	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    pwd1,
		Avatar: "",
		Role:   1,
	}); err != nil {
		c.Abort("user register error")
	}
	c.JSONOK("register success", "/user")
}
