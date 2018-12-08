configtxgen -profile TestTwoOrgsChannel -outputAnchorPeersUpdate ./Org1MSPAnchors.tx -channelID cmbcchannel -asOrg Org1MSP
configtxgen -profile TestTwoOrgsChannel -outputAnchorPeersUpdate ./Org2MSPAnchors.tx -channelID cmbcchannel -asOrg Org2MSP
