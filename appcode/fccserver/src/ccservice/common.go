package ccservice

import (
	"errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	ccFile       = "./config_test.yaml"
	userName     = "User1"
	orgName      = "Org1"
	channelName  = "mychannel"
	ccNetcon     = "netcon"
	ccEstateBook = "estatebook"
	ccEstatetax  = "estatetax"
)

var sdk *fabsdk.FabricSDK
var cclient *channel.Client

func GetSDK() (*fabsdk.FabricSDK, error) {
	if sdk != nil {
		return sdk, nil
	}
	sdk, err := fabsdk.New(config.FromFile(ccFile))
	if err != nil {
		return nil, err
	}
	return sdk, nil
}

func GetChannelClient() (*channel.Client, error) {
	if cclient != nil {
		return cclient, nil
	}
	sdk, err := GetSDK()
	if err != nil {
		return nil, err
	}
	clientContext := sdk.ChannelContext(channelName, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	if clientContext == nil {
		return nil, errors.New("get clientContext failed!")
	}
	cclient, err = channel.New(clientContext)
	return cclient, err
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
