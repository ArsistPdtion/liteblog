package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController)Prepare(){
	log.Println("BaseControll")
	if app,ok := ctx.AppController.(NestPreparer);ok{
		app.NestPrepare()
	}
}
