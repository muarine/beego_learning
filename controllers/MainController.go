/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package controllers

import "html/template"

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "muarine.com"
	this.Data["Email"] = "maoyun0903@gmail.com"
	this.Data["xsrfdata"] = template.HTML(this.XSRFToken())
	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	this.Ctx.WriteString("Hello World")
}
