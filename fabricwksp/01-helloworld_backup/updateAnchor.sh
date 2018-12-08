export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/01-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/01-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/Admin@org1.cmbc.com/msp

peer channel update -o orderer.cmbc.com:7050 -c cmbcchannel -f /home/ssj234/fabricwksp/01-helloworld/orderer/Org1MSPAnchors.tx
