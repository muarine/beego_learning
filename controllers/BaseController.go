/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type NestPreparer interface {
	NestPreparer()
}

type BaseController struct {
	beego.Controller
	i18n.Locale
	isLogin bool
}

func (this *BaseController) Prepare() { // page start time
	this.Data["PageStartTime"] = time.Now()
	// Setting properties.
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}
	this.EnableXSRF = false
}

func (base *BaseController) render() {

}
