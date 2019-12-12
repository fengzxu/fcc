package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Define the Smart Contract structure
type NetconContract struct {
}

type Netcon struct {
	NetconID string `json:"netconid"` //合同编号
	ApplyA   string `json:"applya"`   //受让方（买方）
	ApplyB   string `json:"applyb"`   //转让方（卖方）
	Addr     string `json:"addr"`     //房屋地址
	Area     int    `json:"area"`     //房屋面积
	Balance  int    `json:"balance"`  //转让金额
}

type RecordsInfo struct {
	Size  uint64
	Start string
	End   string
}

var key_recordinfo = "recordeinfo"
var evn_netcon = "evn_netcon"

func (s *NetconContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *NetconContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "create":
		return s.create(APIstub, args)
	case "queryByNetconID":
		return s.queryByNetconID(APIstub, args)
	case "queryByPara":
		return s.queryByPara(APIstub, args)
	case "queryAll":
		return s.queryAll(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}
}

func (s *NetconContract) create(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7. ")
	}
	area, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("area value wrong.")
	}
	balance, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("balance value wrong.")
	}
	netCon := &Netcon{
		NetconID: args[1],
		ApplyA:   args[2],
		ApplyB:   args[3],
		Addr:     args[4],
		Area:     area,
		Balance:  balance,
	}
	jsBytes, err := json.Marshal(netCon)
	if err != nil {
		return shim.Error("marshal json error:" + err.Error())
	}
	err = APIstub.PutState(args[0], jsBytes)
	if err != nil {
		return shim.Error("error on putstate:" + err.Error())
	}
	//update recodeinfo
	recordInfo := &RecordsInfo{}
	rebs, err := APIstub.GetState(key_recordinfo)
	if len(rebs) == 0 {
		recordInfo.Size = 1
		recordInfo.Start = args[0]
		recordInfo.End = args[0]
	} else {
		err = json.Unmarshal(rebs, &recordInfo)
		if err != nil {
			return shim.Error("error on unmarsh recorderinfo:" + err.Error())
		}
		recordInfo.Size = recordInfo.Size + 1
		recordInfo.End = args[0]
	}
	rebs, err = json.Marshal(recordInfo)
	if err != nil {
		return shim.Error("error on marsh new recorderinfo:" + err.Error())
	}
	err = APIstub.PutState(key_recordinfo, rebs)
	if err != nil {
		return shim.Error("error on put new recorderinfo:" + err.Error())
	}
	//broadcast event
	err = APIstub.SetEvent(evn_netcon, []byte("new netcon created with key:"+args[0]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("new netcon created with key:" + args[0]))
}

func (s *NetconContract) queryByNetconID(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 ")
	}
	queryString := fmt.Sprintf("{\"selector\":{\"netconid\":\"%s\"}}", args[0])
	qis, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("queryByNetconID error:" + err.Error())
	}
	defer qis.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for qis.HasNext() {
		queryResponse, err := qis.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *NetconContract) queryByPara(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 ")
	}
	queryString := fmt.Sprintf("{\"selector\":{\""+args[0]+"\":\"%s\"}}", args[1])
	qis, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("queryByNetconID error:" + err.Error())
	}
	defer qis.Close()
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *NetconContract) queryAll(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	rebs, err := APIstub.GetState(key_recordinfo)
	if err != nil {
		return shim.Error("error on get recorderinfo:" + err.Error())
	}
	recordInfo := &RecordsInfo{}
	if len(rebs) == 0 {
		return shim.Success([]byte{})
	}
	err = json.Unmarshal(rebs, &recordInfo)
	if err != nil {
		return shim.Error("error on unmarsh recorderinfo:" + err.Error())
	}
	resultsIterator, err := APIstub.GetStateByRange(recordInfo.Start, recordInfo.End+"1")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	println("<recordeinfo>size:" + strconv.FormatUint(recordInfo.Size, 10) + "  start:" + recordInfo.Start + " end:" + recordInfo.End)
	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(NetconContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
