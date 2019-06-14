package controllers

import "github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"

type NoteController struct {
	BaseController
}

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.Data["key"] = uuid.NewUUID().String()
	ctx.TplName = "note_new.html"
}

func (ctx *NoteController) NestPrepare() {
	ctx.MustLogin() //user must login
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("You don't have ri"))
	}

}
