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
	"log"
	"strconv"
	"time"
)

type Netcon struct {
	ID       string    `xorm:"pk 'id'"`
	CreateDT time.Time //生成时间
	NetconID string    //合同编号
	ApplyA   string    //受让方（买方）
	ApplyB   string    //转让方（卖方）
	Addr     string    //房屋地址
	Area     int       //房屋面积
	Balance  int       //转让金额
	Operator string    //操作人员
	IsCCed   int       //是否已上链
}

type EstateTax struct {
	ID       string    `xorm:"pk 'id'"`
	CreateDT time.Time //生成时间
	TaxID    string    //核税编号
	BookID   string    //不动产权证书编号
	Taxer    string    //纳税人
	Area     int       //房屋面积
	Tax      int       //纳税金额
	Operator string    //操作人员
	IsCCed   int       //是否已上链
}

type EstateBook struct {
	ID       string    `xorm:"pk 'id'"`
	CreateDT time.Time //生成时间
	BookID   string    //不动产证书编号
	NetconID string    //网签合同编号
	TaxID    string    //纳税凭证编号
	Owner    string    //户主
	Addr     string    //房屋地址
	Area     int       //房屋面积
	Operator string    //操作人员
	IsCCed   int       //是否已上链
}

func NetconsAll() (netcons []Netcon, err error) {
	err = GetEngin().Find(&netcons)
	if err != nil {
		log.Println(err)
	}
	return
}

func NetconsCreate(netconid, applya, applyb, addr string, area, balance int, operator string) error {
	netcon := Netcon{
		ID:       strconv.FormatInt(time.Now().Unix(), 10),
		CreateDT: time.Now(),
		NetconID: netconid,
		ApplyA:   applya,
		ApplyB:   applyb,
		Addr:     addr,
		Area:     area,
		Balance:  balance,
		Operator: operator,
	}
	_, err := GetEngin().Insert(&netcon)
	return err
}

func NetconsUpdateCC(uuid string, iscced int) (err error) {
	netcon := &Netcon{}
	ok, err := GetEngin().ID(uuid).Get(netcon)
	if ok {
		netcon.IsCCed = iscced
		_, err = GetEngin().ID(uuid).Cols("is_c_ced").Update(netcon)
	}
	return
}

func EstateBookAll() (books []EstateBook, err error) {
	err = GetEngin().Find(&books)
	if err != nil {
		log.Println("error on EstateBookAll:", err.Error())
	}
	return
}

func EstateBookCreate(bookid, netconid, taxid, owner, addr string, area int, operator string) error {
	book := EstateBook{
		ID:       strconv.FormatInt(time.Now().Unix(), 10),
		CreateDT: time.Now(),
		BookID:   bookid,
		NetconID: netconid,
		TaxID:    taxid,
		Owner:    owner,
		Addr:     addr,
		Area:     area,
		Operator: operator,
	}
	_, err := GetEngin().Insert(&book)
	return err
}

func EstateBookUpdateCC(uuid string, iscced int) (err error) {
	book := &EstateBook{}
	ok, err := GetEngin().ID(uuid).Get(book)
	if ok {
		book.IsCCed = iscced
		_, err = GetEngin().ID(book.ID).Cols("is_c_ced").Update(book)
	}
	return
}

func EstateTaxAll() (taxs []EstateTax, err error) {
	err = GetEngin().Find(&taxs)
	if err != nil {
		log.Println(err)
	}
	return
}

func EstateTaxCreate(taxid, taxer string, area, tax int, operator string) error {
	newtax := EstateTax{
		CreateDT: time.Now(),
		ID:       strconv.FormatInt(time.Now().Unix(), 10),
		TaxID:    taxid,
		//BookID:   bookid,
		Taxer:    taxer,
		Area:     area,
		Tax:      tax,
		Operator: operator,
	}
	_, err := GetEngin().Insert(&newtax)
	return err
}

func EstateTaxUpdateCC(uuid string, iscced int) (err error) {
	tax := &EstateTax{}
	ok, err := GetEngin().ID(uuid).Get(tax)
	if ok {
		tax.IsCCed = iscced
		_, err = GetEngin().ID(tax.ID).Cols("is_c_ced").Update(tax)
	}
	return
}
