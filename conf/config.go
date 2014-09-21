// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
	"path/filepath"
)

const (
	DEFAULT_HTTP_TRY_CNT = 5
	DEFAULT_DB_NAME      = "goAcq"
	DEFAULT_DB_USER      = "root"
	DEFAULT_DB_HOST      = "localhost"
	DEFAULT_DB_PWD       = "admin"
	DEFAULT_CHARSET      = "utf8"
	DEFAULT_DB_TYPE      = "sqlite"
)

var (
	confFilePath = filepath.Join("conf", "conf.ini")
	confFile     *goconfig.ConfigFile

	// default section
	HttpTryCnt int

	// db section
	DbName    string
	DbUser    string
	DbHost    string
	DbPwd     string
	DbCharSet string
	DbType    string
)

func init() {
	var err error
	confFile, err = goconfig.LoadConfigFile(confFilePath)
	if err != nil {
		log.Fatalf("Load configure file failed, return error: %v", err)
	}

	// load config
	HttpTryCnt = confFile.MustInt("", "http_try_cnt", DEFAULT_HTTP_TRY_CNT)

	// db section
	DbName = confFile.MustValue("db", "db_name", DEFAULT_DB_NAME)
	DbUser = confFile.MustValue("db", "db_user", DEFAULT_DB_USER)
	DbHost = confFile.MustValue("db", "db_host", DEFAULT_DB_HOST)
	DbPwd = confFile.MustValue("db", "db_pwd", DEFAULT_DB_PWD)
	DbCharSet = confFile.MustValue("db", "db_charSet", DEFAULT_CHARSET)
	DbType = confFile.MustValue("db", "db_type", DEFAULT_DB_TYPE)
}
