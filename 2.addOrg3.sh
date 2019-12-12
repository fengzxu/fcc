#!/bin/sh
echo 
echo ===========================================
echo ===  add Org3 to network
echo ==========================================
echo

CHANNEL_NAME="mychannel"

echo
echo "========= Creating config transaction to add org3 to network =========== "
echo

echo "Generate the Org3 Crypto Material"
cd ../first-network/org3-artifacts
../../bin/cryptogen generate --config=./org3-crypto.yaml
export FABRIC_CFG_PATH=$PWD && ../../bin/configtxgen -printOrg Org3MSP > ../channel-artifacts/org3.json
cp -r crypto-config/peerOrganizations/org3.example.com/ ../crypto-config/peerOrganizations/
cd ../../fcc

echo "add org3 join channel ...."
export IMAGE_TAG=1.4.3 && export COMPOSE_PROJECT_NAME=net && docker-compose -f ../first-network/docker-compose-org3.yaml up -d 
docker exec cli /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/step1org3.sh
docker exec cli /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/step2org3.sh
