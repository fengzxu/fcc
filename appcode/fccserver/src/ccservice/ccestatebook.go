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

package ccservice

import (
	"comm"
	"encoding/json"
)

type CEstateBook struct {
	BookID string `json:"bookid"` //不动产证书编号
	Owner  string `json:"owner"`  //户主
	Addr   string `json:"addr"`   //房屋地址
	Area   int    `json:"area"`   //房屋面积
}

func EstateBookCreate(uuid, bookid, owener, addr, area string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	//check if bookid exist
	res = EstateBookQueryByBookid(bookid)
	if res.Code > 0 {
		return
	}
	ns := res.Status.([]CEstateBook)
	if len(ns) > 0 {
		res.Code = 1
		res.Status = bookid + " already exited!"
		return
	}
	//new
	bs, err := CCinvoke(cclient, ccEstateBook, "create",
		[]string{uuid, bookid, owener, addr, area})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		res.Status = string(bs)
	}
	return
}

func EstateBookQueryByBookid(bookid string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccEstateBook, "queryByBookID", []string{bookid})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		str := string(bs)
		var books []CEstateBook
		err = json.Unmarshal([]byte(str), &books)
		if err != nil {
			res.Code = 1
			res.Status = err.Error()
		} else {
			res.Status = books
		}
	}
	return
}

func EstateBookQueryAll() (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccEstateBook, "queryAll", nil)
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		str := string(bs)
		var cs []CEstateBook
		err = json.Unmarshal([]byte(str), &cs)
		if err != nil {
			res.Code = 1
			res.Status = err.Error()
		} else {
			res.Status = cs
		}
	}
	return
}
