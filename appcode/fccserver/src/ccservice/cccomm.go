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
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

var (
	ccFile       = "./config.yaml"
	userName     = "User1"
	orgName      = "Org1"
	channelName  = "mychannel"
	ccNetcon     = "netcon"
	ccEstateBook = "estatebook"
	ccEstatetax  = "estatetax"
)

var sdk *fabsdk.FabricSDK
var cclient *channel.Client

func InitCCOnStart() error {
	sdk, err := fabsdk.New(config.FromFile(ccFile))
	if err != nil {
		log.Println("WARN: init Chaincode SDK error:", err.Error())
		return err
	}
	clientContext := sdk.ChannelContext(channelName, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	if clientContext == nil {
		log.Println("WARN: init Chaincode clientContext error:", err.Error())
		return err
	} else {
		cclient, err = channel.New(clientContext)
		if err != nil {
			log.Println("WARN: init Chaincode cclient error:", err.Error())
		}
	}
	log.Println("Chaincode client initialed successfully.")
	return nil
}

func GetChannelClient() *channel.Client {
	return cclient
}

func CCinvoke(channelClient *channel.Client, ccname, fcn string, args []string) ([]byte, error) {
	var tempArgs [][]byte
	for i := 0; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}
	qrequest := channel.Request{
		ChaincodeID:     ccname,
		Fcn:             fcn,
		Args:            tempArgs,
		TransientMap:    nil,
		InvocationChain: nil,
	}
	//log.Println("cc exec request:",qrequest.ChaincodeID,"\t",qrequest.Fcn,"\t",qrequest.Args)
	response, err := channelClient.Execute(qrequest)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func CCquery(channelClient *channel.Client, ccname, fcn string, args []string) ([]byte, error) {
	var tempArgs [][]byte
	if args == nil {
		tempArgs = nil
	} else {
		for i := 0; i < len(args); i++ {
			tempArgs = append(tempArgs, []byte(args[i]))
		}
	}
	qrequest := channel.Request{
		ChaincodeID:     ccname,
		Fcn:             fcn,
		Args:            tempArgs,
		TransientMap:    nil,
		InvocationChain: nil,
	}
	response, err := channelClient.Query(qrequest)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}
