/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego/context"

	"quickstart/filters"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterUser)

	// 基本路由 闭包函数构成
	beego.Get("/basic", func(context *context.Context) {
		context.Output.Body([]byte("Basic Get Request"))
	})
	beego.Post("/basic", func(context *context.Context) {
		context.Output.Body([]byte("Basic Post Request"))
	})
	beego.Any("/foo", func(context *context.Context) {
		context.Output.Body([]byte("Match Any Request"))
	})
	// 固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	// 正则路由
	beego.Router("/api/:id", &controllers.RegularController{})
	// 自定义方法及 RESTful 规则
	beego.Router("/api/custom/list", &controllers.RestController{}, "*:ListFood")
	beego.Router("/api/custom/query", &controllers.RestController{}, "Get:Query")
	beego.Router("/api/custom/update", &controllers.RestController{}, "Post:Update")
	beego.Router("/api/custom/updateForm", &controllers.RestController{}, "Post:UpdateFormData")
	//注解路由
	//beego.Include(&controllers.AnnotationController{})

}
