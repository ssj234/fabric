export set GOPATH=/home/ssj234/fabricwksp/05-chaincode
export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/02-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.cmbc.com:7051
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/02-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/Admin@org1.cmbc.com/msp

# 初始化chaincode 代码

peer chaincode instantiate -o orderer.cmbc.com:7050 -C cmbcchannel666 -n statechaincode -v 1.3 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
