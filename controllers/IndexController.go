/*
 * Copyright (c) 2017. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package controllers

import "fmt"

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.Data["param"] = ""
	this.Layout = "layout/banner.tpl"
	var fun = func() int64 {
		return 1
	}

	fmt.Sprintln("", fun)
	this.TplName = "index.tpl"
}
