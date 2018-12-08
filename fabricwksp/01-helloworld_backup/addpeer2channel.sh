export set FABRIC_CFG_PATH=/home/ssj234/fabricwksp/01-helloworld/peer
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_MSPCONFIGPATH=/home/ssj234/fabricwksp/01-helloworld/fabricconfig/crypto-config/peerOrganizations/org1.cmbc.com/users/Admin@org1.cmbc.com/msp

peer channel join -b /home/ssj234/fabricwksp/01-helloworld/cmbcchannel.block
