export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/02-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org2MSP
export set CORE_PEER_ADDRESS=peer0.org2.cmbc.com:17051
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/02-helloworld/fabricconfig/crypto-config/peerOrganizations/org2.cmbc.com/users/Admin@org2.cmbc.com/msp

peer channel join -b /home/ssj234/fabricwksp/02-helloworld/org2peer/cmbcchannel666.block
