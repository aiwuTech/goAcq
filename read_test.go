// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this soReurce code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcq

import (
	"testing"
)

func TestReadWebPageSource(t *testing.T) {
	url := "http://www.baidu.com/"

	html, err := readWebPageSource(url, EncodeType_UTF8)
	if err != nil {
		t.Errorf("test failed. data: %v, err: %v", html, err)
	}
}
