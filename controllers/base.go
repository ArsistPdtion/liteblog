package controllers

import (
	"github.com/ArsistPdtion/liteblog/models"
	"github.com/astaxie/beego"
	"log"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	IsLogin bool        //label user is login
	User    models.User //login user
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseControll")
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count, omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOK(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}
