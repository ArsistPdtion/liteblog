package routers

import (
	"github.com/arsistPdtion/liteblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Include(&controllers.IndexController{})
}
