// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcquisition

import (
	"regexp"
	"strings"
	"strconv"
)

type ETargetEncodeType int32

const (
	EncodeType_GB2312 ETargetEncodeType = iota
	EncodeType_UTF8
	EncodeType_BIG5
)

type EMatchMode int32

const (
	Mode_Regex EMatchMode = iota
	Mode_String
)

var UrlPattern = "((http|ftp|https)://)(([a-zA-Z0-9\\._-]+\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?"

type AcqNode struct {
	NodeName       string
	TargetEncode   ETargetEncodeType
	MatchMode      EMatchMode
	TargetListUrls []string
	TargetUrlRule  *TargetUrlMatch
}

func NewDefaultAcqNode(nodeName string) *AcqNode {
	return &AcqNode{
		NodeName: nodeName,
		TargetEncode: EncodeType_GB2312,
		MatchMode: Mode_String,
		TargetListUrls: make([]string, 0),
		TargetUrlRule: &TargetUrlMatch{
			BeginHtml: "",
			EndHtml: "",
			ImgHandleFlag: true,
			MustRegex: "",
			NoRegex: "",
		},
	}
}

func (a *AcqNode) SetNodeName(nodeName string) {
	a.NodeName = nodeName
}

func (a *AcqNode) GetNodeName() string {
	return a.NodeName
}

func (a *AcqNode) SetTargetEncode(encode ETargetEncodeType) {
	a.TargetEncode = encode
}

func (a *AcqNode) GetTargetEncode() ETargetEncodeType {
	return a.TargetEncode
}

func (a *AcqNode) SetMatchMode(mode EMatchMode) {
	a.MatchMode = mode
}

func (a *AcqNode) GetMatchMode() EMatchMode {
	return a.MatchMode
}

func (a *AcqNode) addListUrl(url string) {
	a.TargetListUrls = append(a.TargetListUrls, url)
}

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

func (a *AcqNode) Len() int {
	return len(a.TargetListUrls)
}

func (a *AcqNode) GetTargetListUrls() []string {
	return a.TargetListUrls
}

type TargetUrlMatch struct {
	BeginHtml     string
	EndHtml       string
	ImgHandleFlag bool
	MustRegex     string
	NoRegex       string
}

func (t *TargetUrlMatch) SetBeginHtml(html string) {
	t.BeginHtml = html
}

func (t *TargetUrlMatch) GetBeginHtml() string {
	return t.BeginHtml
}

func (t *TargetUrlMatch) SetEndHtml(html string) {
	t.EndHtml = html
}

func (t *TargetUrlMatch) GetEndHtml() string {
	return t.EndHtml
}

func (t *TargetUrlMatch) SetImgHandleFlag(flag bool) {
	t.ImgHandleFlag = flag
}

func (t *TargetUrlMatch) GetImgHandleFlag() bool {
	return t.ImgHandleFlag
}

func (t *TargetUrlMatch) SetMustRegex(regex string) {
	t.MustRegex = regex
}

func (t *TargetUrlMatch) GetMustRegex() string {
	return t.MustRegex
}

func (t *TargetUrlMatch) SetNoRegex(regex string) {
	t.NoRegex = regex
}

func (t *TargetUrlMatch) GetNoRegex() string {
	return t.NoRegex
}
