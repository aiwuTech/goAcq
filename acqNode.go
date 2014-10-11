// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcquisition

import (
	"regexp"
	"strconv"
	"strings"
)

type ETargetEncodeType string

const (
	EncodeType_GB2312 ETargetEncodeType = "gb2312"
	EncodeType_UTF8   ETargetEncodeType = "utf8"
	EncodeType_BIG5   ETargetEncodeType = "big5"
)

type EMatchMode int32

const (
	Mode_Regex EMatchMode = iota
	Mode_String
)

var UrlPattern = "((http|ftp|https)://)(([a-zA-Z0-9\\._-]+\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?"


type targetUrlMatch struct {
	beginHtml     string
	endHtml       string
	imgHandleFlag bool
	mustRegex     string
	noRegex       string
}

type AcqNode struct {
	nodeName       string
	targetEncode   ETargetEncodeType
	matchMode      EMatchMode
	targetListUrls []string
	targetUrlRule  *targetUrlMatch
	targetRule     []*AcqTarget
}

func NewDefaultAcqNode(nodeName string) *AcqNode {
	return &AcqNode{
		nodeName:       nodeName,
		targetEncode:   EncodeType_GB2312,
		matchMode:      Mode_String,
		targetListUrls: make([]string, 0),
		targetUrlRule: &targetUrlMatch{
			beginHtml:     "",
			endHtml:       "",
			imgHandleFlag: true,
			mustRegex:     "",
			noRegex:       "",
		},
	}
}

func (a *AcqNode) SetNodeName(nodeName string) {
	a.nodeName = nodeName
}

func (a *AcqNode) GetNodeName() string {
	return a.nodeName
}

func (a *AcqNode) SetTargetEncode(encode ETargetEncodeType) {
	a.targetEncode = encode
}

func (a *AcqNode) GetTargetEncode() ETargetEncodeType {
	return a.targetEncode
}

func (a *AcqNode) SetMatchMode(mode EMatchMode) {
	a.matchMode = mode
}

func (a *AcqNode) GetMatchMode() EMatchMode {
	return a.matchMode
}

func (a *AcqNode) SetTargetUrlBeginHtml(html string) {
	a.targetUrlRule.beginHtml = html
}

func (a *AcqNode) GetTargetUrlBeginHtml() string {
	return a.targetUrlRule.beginHtml
}

func (a *AcqNode) SetTargetUrlEndHtml(html string) {
	a.targetUrlRule.endHtml = html
}

func (a *AcqNode) GetTargetUrlEndHtml() string {
	return a.targetUrlRule.endHtml
}

func (a *AcqNode) addListUrl(url string) {
	a.targetListUrls = append(a.targetListUrls, url)
}

// 添加采集列表
func (a *AcqNode) AddListUrls(urls ...string) (in []string) {

	in = make([]string, 0)

	for _, url := range urls {
		if matched, err := regexp.MatchString(UrlPattern, url); matched && err == nil {
			a.addListUrl(url)

			in = append(in, url)
		}
	}

	return
}

// http://www.aiwutech.com/test/list_(*).html
func (a *AcqNode) AddListUrlsByTag(urlMatch string, min, max, gap uint32) (in []string) {

	pattern := "(*)"
	if !strings.Contains(urlMatch, pattern) || (min > max) {
		return
	}

	in = make([]string, 0)
	for min <= max {
		tmpNum := strconv.Itoa(int(min))
		url := strings.Replace(urlMatch, pattern, tmpNum, -1)

		in = append(in, url)
		min += gap
	}

	return a.AddListUrls(in...)
}

// 采集目标列表个数
func (a *AcqNode) Len() int {
	return len(a.targetListUrls)
}

// 采集列表url
func (a *AcqNode) GetTargetListUrls() []string {
	return a.targetListUrls
}

func (a *AcqNode) Exec() {
	for _, listUrl := range a.GetTargetListUrls() {
		pageSource, err := readWebPageSource(listUrl, a.GetTargetEncode())
		if err != nil {
			continue
		}

		switch a.GetMatchMode() {
		case Mode_Regex:
		case Mode_String:
			pageSource = strings.SplitAfterN(pageSource, a.GetTargetUrlBeginHtml(), 1)[0]
			pageSource = strings.Split(pageSource, a.GetTargetUrlEndHtml())[0]

			println(pageSource)
		}
	}
}
