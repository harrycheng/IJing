package controllers

import (
	ijing "IJing/service"
	"bytes"
	"encoding/base64"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
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

		if len(newDivinatoryDetail) == 0 {
			newDivinatory, newDivinatoryDetail = ijing.IjDivinatory()
			IjMap[newDivinatory] = newDivinatoryDetail
		}
	}

	t := time.Now()
	year, month, day := t.Date()
	nextoDay := time.Date(year, month, day+1, 0, 0, 0, 0, t.Location())

	duration := (nextoDay.Unix() - t.Unix())
	c.Ctx.SetCookie("divinatory", base64.StdEncoding.EncodeToString([]byte(newDivinatory)), duration)

	detailHtml := strings.ReplaceAll(newDivinatoryDetail, "\r", "</br>")
	detailHtml = strings.ReplaceAll(newDivinatoryDetail, "\r", "</br>")
	c.Data["Website"] = "IJing"
	c.Data["Email"] = "hc.harrycheng@gmail.com"
	c.Data["divinatory"] = newDivinatory
	c.Data["detail"] = detailHtml
	c.Data["divinatoryTime"] = GetLunarStr()
	c.TplName = "index.tpl"
}

func GetLunarStr() string {
	t := time.Now()
	c := calendar.BySolar(int64(t.Year()), int64(t.Month()), int64(t.Day()), int64(t.Hour()), int64(t.Minute()), int64(t.Second()))
	var buffer bytes.Buffer
	buffer.WriteString("农历：" + c.Ganzhi.YearGanzhiAlias() + "年")
	buffer.WriteString(c.Ganzhi.MonthGanzhiAlias() + "月")
	buffer.WriteString(c.Ganzhi.DayGanzhiAlias() + "日")

	buffer.WriteString("公历：" + t.Format("2006-01-02 03:04:05"))

	return buffer.String()
}

func (c *MainController) Post() {
	c.TplName = "index.tpl"
}
