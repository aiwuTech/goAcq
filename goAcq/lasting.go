// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package goAcquisition

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	. "goAcquisition/conf"
	"log"
	"path/filepath"
	"time"
	"os"
)

var (
	Engine *xorm.Engine
)

// lasting for AcqNodes
type TblAcqNode struct {
	Id             int64
	NodeName       string `xorm:"unique"`
	NodeEncode     string
	MatchMode      int32
//	targetUrlMatch `xorm:"extends"`
	AcqCnt         int32
	Created        time.Time `xorm:"created"`
	LastAcq        time.Time `xomr:"updated"`
}

// lasting for lists
type TblTargetList struct {
	Id      int64
	NodeId  int64
	ListUrl string
	AcqCnt  int32
	Created time.Time `xorm:"created"`
	LastAcq time.Time `xorm:"updated"`
}

// lasting for targets
type TblTarget struct {
	Id        int64
	NodeId    int64
	ListId    int64
	TargetUrl string
	Keyword   string
	Summary   string
	Title     string
	From      string
	PostTime  string
	Content   string
	SavePath  string
	AcqCnt    int32
	Created   time.Time `xorm:"created"`
	LastAcq   time.Time `xorm:"updated"`
}

func init() {

	var err error
	Engine, err = setDbEngine()
	if err != nil {
		log.Fatalf("create database connection return error: %v", err)
	}

	// sync table
	if err := Engine.Sync(new(TblAcqNode), new(TblTargetList), new(TblTarget)); err != nil {
		log.Fatalf("sync database's table return error: %v", err)
	}
}

func setDbEngine() (*xorm.Engine, error) {
	engine, err := xConnDb()
	if engine == nil || err != nil {
		return nil, err
	}

	engine.SetMapper(&core.SameMapper{})
	engine.ShowSQL = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	return engine, nil
}

func xConnDb() (*xorm.Engine, error) {

	switch DbType {
	case "sqlite":
		return xorm.NewEngine("sqlite3", filepath.Join(os.Getenv("GOPATH"), "src", "goAcquisition", "data", fmt.Sprintf("%s.db", DbName)))
	case "mysql":
		return xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%v)/%s?charset=%v", DbUser, DbPwd, DbHost, DbName, DbCharSet))
	case "pgsql":
		return xorm.NewEngine("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPwd, DbName))
	}

	return nil, fmt.Errorf("UnSopported database type: %v", DbType)
}
