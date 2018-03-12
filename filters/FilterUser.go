// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context) {
	uri := ctx.Request.RequestURI
	beego.Debug("过滤器已拦截 uri:", uri)
	var result = `{"code":201,"message":"token is not null"}`
	// 校验token
	token := ctx.Input.Header("token")
	if token == "" {
		token = ctx.Input.Query("token")
	}
	beego.Debug("token=", token)
	if token == "" {
		ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Output.Body([]byte(result))
		ctx.ResponseWriter.Flush()
	}
}
