#export set GOPATH=/home/ssj234/fabricwksp/05-chaincode
#export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/02-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.cmbc.com:7051
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/02-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/User1@org1.cmbc.com/msp
# 调用chaincode 代码

peer chaincode invoke -o orderer.cmbc.com:7050 -C cmbcchannel666 -n twoorgchaincode -v 1.0 -c '{"Args":["invoke","a","b","10"]}'
