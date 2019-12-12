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

type CEstateTax struct {
	TaxID  string `json:"taxid"`  //核税编号
	BookID string `json:"bookid"` //不动产权证书编号
	Taxer  string `json:"taxer"`  //纳税人
	Area   int    `json:"area"`   //房屋面积
	Tax    int    `json:"tax"`    //纳税金额
}

func EstateTaxCreate(uuid, taxid, bookid, taxer, area, tax string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	//check if taxid exist
	res = EstateTaxQueryByTaxid(taxid)
	if res.Code >0 {
		return
	}
	ns:=res.Status.([]CEstateTax)
	if len(ns)>0 {
		res.Code = 1
		res.Status = taxid+" already exited!"
		return
	}
	bs, err := CCinvoke(cclient, ccEstatetax, "create",
		[]string{uuid, taxid, bookid, taxer, area, tax})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		res.Status = string(bs)
	}
	return
}

func EstateTaxQueryByTaxid(taxid string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccEstatetax, "queryByTaxID", []string{taxid})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		str := string(bs)
		var cs []CEstateTax
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

func EstateTaxQueryAll() (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccEstatetax, "queryAll", nil)
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		str := string(bs)
		var cs []CEstateTax
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

func EstateTaxQueryAllid() (res comm.ResResult) {
	res = EstateTaxQueryAll()
	if res.Code == 0 {
		ids := []string{}
		for _, con := range res.Status.([]CEstateTax) {
			ids = append(ids, con.TaxID)
		}
		res.Status = ids
	}
	return
}
