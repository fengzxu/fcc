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

type CNetcon struct {
	NetconID string `json:"netconid"` //合同编号
	ApplyA   string `json:"applya"`   //受让方（买方）
	ApplyB   string `json:"applyb"`   //转让方（卖方）
	Addr     string `json:"addr"`     //房屋地址
	Area     int    `json:"area"`     //房屋面积
	Balance  int    `json:"balance"`  //转让金额
}

func NetconCreate(uuid, netconid, applya, applyb, addr, area, balance string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	//check if netconid exist
	res = NetconQueryByNetconid(netconid)
	if res.Code >0 {
		return
	}
	ns:=res.Status.([]CNetcon)
	if len(ns)>0 {
		res.Code = 1
		res.Status = netconid+" already exited!"
		return
	}
	//new
	bs, err := CCinvoke(cclient, ccNetcon, "create",
		[]string{uuid, netconid, applya, applyb, addr, area, balance})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		res.Status = string(bs)
	}
	return
}

func NetconQueryByNetconid(netconid string) (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccNetcon, "queryByNetconID", []string{netconid})
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		//res.Status = string(bs)
		str := string(bs)
		var ns []CNetcon
		err = json.Unmarshal([]byte(str), &ns)
		if err != nil {
			res.Code = 1
			res.Status = err.Error()
		} else {
			res.Status = ns
		}
	}
	return
}

func NetconQueryAll() (res comm.ResResult) {
	cclient = GetChannelClient()
	if cclient == nil {
		res.Code = 1
		res.Status = "Chaincode service uninitialed."
		return
	}
	bs, err := CCquery(cclient, ccNetcon, "queryAll", nil)
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		str := string(bs)
		var cs []CNetcon
		err = json.Unmarshal([]byte(str), &cs)
		if err != nil {
			res.Code = 1
			res.Status = err.Error()
		}else {
			res.Status = cs
		}
	}
	return
}

func NetconGetAllCCid() (res comm.ResResult) {
	res = NetconQueryAll()
	if res.Code == 0 {
		ids := []string{}
		for _, con := range res.Status.([]CNetcon) {
			ids = append(ids, con.NetconID)
		}
		res.Status = ids
	}
	return
}