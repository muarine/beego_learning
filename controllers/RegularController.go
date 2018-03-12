/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package controllers

import "github.com/astaxie/beego"

type RegularController struct {
	BaseController
}

func (this *RegularController) Get() {
	id := this.Ctx.Input.Param(":id")
	beego.Informational("received url param :", id)
	this.Data["param"] = "Regular Controller"
	this.Data["Website"] = "website"
	this.Data["Email"] = "mail"
	this.TplName = "index.tpl"
}
