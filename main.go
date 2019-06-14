package main

import (
	"encoding/gob"
	"github.com/ArsistPdtion/liteblog/models"
	_ "github.com/ArsistPdtion/liteblog/routers"
	"github.com/astaxie/beego"
	"strings"
	_ "github.com/ArsistPdtion/liteblog/models"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}

func initTemplate(){
	beego.AddFuncMap("equrl", func(x,y string)bool {
		s1 := strings.Trim(x,"/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1,s2) == 0
	})
}

func initSession(){
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

