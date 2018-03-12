/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"fmt"
	_ "quickstart/routers"
	"time"

	"os"

	"github.com/astaxie/beego"
)

type Person struct {
	UserName string
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"./logs/quickstart.log"}`)
	}

	beego.AppConfig.Set("up_time", fmt.Sprintf("%d", time.Now().Unix()))
	// beego默认注册了static为静态目录
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/down1", "static/download1")
	beego.AddTemplateExt(".tpl")
	beego.SetLogFuncCall(true)
	beego.Informational("运行模式:", beego.BConfig.RunMode)
	beego.Informational("环境变量: a=", os.Getenv("a"), ",b=", os.Getenv("b"))
	beego.Run()
}
