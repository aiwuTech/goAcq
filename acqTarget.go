// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcq

type AcqTarget struct {
	TargetUrl     string
	KeyWrodFilter string
	SummaryFilter string
	TitleRule     *OriginRule
	AuthorRule    *OriginRule
	FromRule      *OriginRule
	PostTimeRule  *OriginRule
	ContentRule   *OriginRule
}

type OriginRule struct {
	Match  string
	Filter string
}
