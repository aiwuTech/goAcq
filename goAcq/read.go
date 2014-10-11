// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcquisition

import (
	. "goAcquisition/conf" // The import path format is not recommented
	"code.google.com/p/mahonia"
	"fmt"
	"io/ioutil"
	"net/http"
)

func readWebPageSource(url string, encode ETargetEncodeType) (string, error) {
	for i := 1; i <= HttpTryCnt; i++ {

		rsp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer rsp.Body.Close()

		if rsp.StatusCode == http.StatusOK {
			decoder := mahonia.NewDecoder(string(encode))
			data, e := ioutil.ReadAll(decoder.NewReader(rsp.Body))
			if e != nil {
				continue
			}

			return string(data), e
		}
	}

	return "", fmt.Errorf("read web page source timeout.")
}

