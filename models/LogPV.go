// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package models

import (
	"path/filepath"
	"strings"
)

var (
	NotPV = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
)

const big = 0xFFFFFF

func LogPV(urls string) bool {
	ext := filepath.Ext(urls)
	if ext == "" {
		return true
	}
	for _, v := range NotPV {
		if v == strings.ToLower(ext) {
			return false
		}
	}
	return true
}
