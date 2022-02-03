package controllers

import (
	ijing "IJing/service"
	"encoding/base64"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type MainController struct {
	beego.Controller
}

var IjMap = make(map[string]string)

func (c *MainController) Get() {
	preDivinatory := c.Ctx.GetCookie("divinatory")

	newDivinatory := ""
	newDivinatoryDetail := ""
	if len(preDivinatory) == 0 {
		newDivinatory, newDivinatoryDetail = ijing.IjDivinatory()

		IjMap[newDivinatory] = newDivinatoryDetail
	} else {
		preDiv, _ := base64.StdEncoding.DecodeString(preDivinatory)
		newDivinatory = string(preDiv)
		newDivinatoryDetail = IjMap[newDivinatory]
	}

	t := time.Now()
	year, month, day := t.Date()
	nextoDay := time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())

	duration := (nextoDay.Unix() - t.Unix())
	c.Ctx.SetCookie("divinatory", base64.StdEncoding.EncodeToString([]byte(newDivinatory)), duration)

	c.Data["Website"] = "IJing"
	c.Data["Email"] = "hc.harrycheng@gmail.com"
	c.Data["divinatory"] = newDivinatory
	c.Data["detail"] = newDivinatoryDetail
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.TplName = "index.tpl"
}
