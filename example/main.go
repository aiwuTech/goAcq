// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package main

import (
	"goAcquisition"
)

func main() {
	acqNode := goAcquisition.NewDefaultAcqNode("test")
	acqNode.AddListUrlsByTag("http://cd.qq.com/l/auto/focusauto/list20131121151120_(*).htm", 2, 5, 1)
	acqNode.SetMatchMode(goAcquisition.Mode_String)
	acqNode.SetTargetEncode(goAcquisition.EncodeType_GB2312)
	acqNode.SetTargetUrlBeginHtml("<div class=\"box_hr16\"></div>")
	acqNode.SetTargetUrlEndHtml("<div class=\"box_hr16\"></div>")

	acqNode.Exec()
}

