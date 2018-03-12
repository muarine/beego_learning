// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package controllers

// Annotation API
type AnnotationController struct {
	BaseController
}

func (c *AnnotationController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *AnnotationController) StaticBlock() {
	this.TplName = "index.tpl"
}

// @router /allblock/:key [get]
func (this *AnnotationController) AllBlock() {
	this.TplName = "index.tpl"
}
