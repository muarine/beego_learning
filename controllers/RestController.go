// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"

	"quickstart/utils"

	"github.com/astaxie/beego"
)

type User struct {
	Id    int         `form:"id"`
	Name  interface{} `form:"name"`
	Age   int         `form:"age"`
	Email string      `form:"email"`
}

type RestController struct {
	BaseController
}

func (this *RestController) ListFood() {
	this.Data["json"] = utils.Build()
	this.ServeJSON()
}

func (this *RestController) Query() {
	urlExt := this.Ctx.Input.Param(":ext")
	// req url ext : json
	beego.Informational("req url ext :", urlExt)
	this.TplName = "index.tpl"
}

// 模拟post json请求，直接解析到struct里
func (this *RestController) Update() {
	u := User{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	beego.Debug("解析入参：", u)
	if result, err := json.Marshal(u); err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = string(result)
	}
	this.ServeJSON()
}

func (this *RestController) UpdateFormData() {
	u := User{}
	this.ParseForm(&u)
	beego.Debug("解析表单提交参数:", u)
	this.TplName = "index.tpl"
}
