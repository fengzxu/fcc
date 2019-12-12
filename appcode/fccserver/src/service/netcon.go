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

package service

import (
	"ccservice"
	"comm"
	"db"
)

func NetconGetAll() (res comm.ResResult) {
	netcons, err := db.NetconsAll()
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
		res.Status = netcons
	}
	return
}

func NetconCreate(netconid, applya, applyb, addr string, area, balance int) (res comm.ResResult) {
	err := db.NetconsCreate(netconid, applya, applyb, addr, area, balance, GetOperator())
	if err != nil {
		res.Code = 1
		res.Status = err.Error()
	} else {
		res.Code = 0
	}
	return
}

func NetconToCC(uuid, netconid, applya, applyb, addr, area, balance string) (res comm.ResResult) {
	res = ccservice.NetconCreate(uuid, netconid, applya, applyb, addr, area, balance)
	if res.Code == 0 {
		err := db.NetconsUpdateCC(uuid, 1)
		if err != nil {
			res.Code = 1
			res.Status = "上链成功，但更新业务数据库中的上链标志失败：" + err.Error()
		}
	}
	return
}
