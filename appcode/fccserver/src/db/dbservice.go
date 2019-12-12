/*
Copyright xujf000@gmail.com .2020. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var dbfile = "./server.db"
var engins map[string]*xorm.Engine = make(map[string]*xorm.Engine)

func SetEngin(orm *xorm.Engine, keys ...string) {
	if len(keys) == 0 {
		engins["default"] = orm
	} else {
		engins[keys[0]] = orm
	}
}

func GetEngin(keys ...string) (e *xorm.Engine) {
	if len(keys) == 0 {
		return engins["default"]
	} else {
		return engins[keys[0]]
	}
}

func InitDB() error {
	//init dabase
	orm, err := xorm.NewEngine("sqlite3", dbfile)
	if err != nil {
		log.Println("error on init db file:", err.Error())
		return err
	}
	orm.DatabaseTZ = time.Local //
	orm.TZLocation = time.Local //
	orm.SetMaxIdleConns(10)
	orm.SetMaxOpenConns(30)
	orm.ShowSQL(false)
	SetEngin(orm)

	//init tables
	err = initTables(orm)
	if err != nil {
		log.Println("system db error:", err.Error())
	} else {
		log.Println("system db initiated successfully.")
	}
	return err
}

func initTables(orm *xorm.Engine) (err error) {

	//create tables if not exits
	err = orm.CreateTables(
		&Netcon{},
		&EstateBook{},
		&EstateTax{},
	)
	return
}
