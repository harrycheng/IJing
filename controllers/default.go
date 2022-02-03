package controllers

import (
	ijing "IJing/service"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	divinatory, divinatoryDetail := ijing.IjDivinatory()

	c.Data["Website"] = "IJing"
	c.Data["Email"] = "hc.harrycheng@gmail.com"
	c.Data["divinatory"] = divinatory
	c.Data["divinatoryDetail"] = divinatoryDetail
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	c.TplName = "index.tpl"
}
