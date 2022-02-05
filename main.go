package main

import (
	_ "IJing/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//exec.Command(`open`, `http://localhost:20080/`).Start()
	beego.Run()
}
