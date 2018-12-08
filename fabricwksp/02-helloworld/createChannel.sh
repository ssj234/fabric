export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/02-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/Admin@org1.cmbc.com/msp

peer channel create -t 50 -o orderer.cmbc.com:7050 -c cmbcchannel888 -f /home/ssj234/fabricwksp/02-helloworld/orderer/cmbcchannel888.tx
