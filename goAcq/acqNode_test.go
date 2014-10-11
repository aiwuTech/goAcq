// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcquisition

import "testing"

func TestAddListUrls(t *testing.T) {
	acq := NewDefaultAcqNode("test")

	in := acq.AddListUrls("http:www.baidu.com/", "http://192.168.1.1", "http://www.baidu.com/s?wd=a&rsv_spt=1&issp=1&rsv_bp=0&ie=utf-8&tn=baiduhome_pg&inputT=1236")
	if acq.Len() != 2 {
		t.Errorf("Test failed. in: %v", in)
	}
}

func TestAddListUrlsByTag(t *testing.T) {
	acq := NewDefaultAcqNode("test")

	in := acq.AddListUrlsByTag("http://www.aiwutech.com/test/list_(*).html", 1, 10, 1)
	if acq.Len() != 10 {
		t.Errorf("Test failed. in: %v", in)
	}
}
