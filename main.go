package main

import (
	_ "IJing/routers"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
)

func main() {
	exec.Command(`open`, `http://localhost:8080/`).Start()
	beego.Run()
}
