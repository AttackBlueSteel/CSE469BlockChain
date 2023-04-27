Group: #77
Names: Yue Fang, Ryan Hughes, Michael Gilson, Alexandr Zincenko
Program Description: In order to run the program, it is vital to follow the steps outlined in the following website https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html (all 3 sections of Getting Started -Install). Afterwards, please unzip the code file and place it in your go/src/github.com folder. After placing the folder, please follow the demo video and copy & paste the included commands filling in the appropriate information indicated by "<>" into terminal.

** Please Note this was tested on Ubuntu 22.04 and should work on previous version of Ubuntu (18.04) if all the installed libraries are up to date to their newest version. **


********************************************* COMMANDS **********************************************

sudo su;

source /etc/profile;cd /home/<your_username>/go/src/github.com/hyperledger/fabric/fabric-samples/test-network;./network.sh down;./network.sh up createChannel;cd /home/<your_username>/go/src/github.com/hyperledger/fabric/fabric-samples/chaincode/case/;GO111MODULE=on go mod vendor;cd /home/<your_username>/go/src/github.com/hyperledger/fabric/fabric-samples/test-network;export PATH=${PWD}/../bin:${PWD}:$PATH;export FABRIC_CFG_PATH=$PWD/../config/;export CORE_PEER_TLS_ENABLED=true;export CORE_PEER_LOCALMSPID="Org1MSP";export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt;export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp;export CORE_PEER_ADDRESS=localhost:7051; source /etc/profile;peer lifecycle chaincode package case.tar.gz --path ../chaincode/case/ --lang golang --label case_1.0;export CORE_PEER_TLS_ENABLED=true;export CORE_PEER_LOCALMSPID="Org1MSP";export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt;export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp;export CORE_PEER_ADDRESS=localhost:7051;source /etc/profile;peer lifecycle chaincode install case.tar.gz;export CORE_PEER_LOCALMSPID="Org2MSP";export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt;export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt;export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp;export CORE_PEER_ADDRESS=localhost:9051;source /etc/profile;peer lifecycle chaincode install case.tar.gz;peer lifecycle chaincode queryinstalled;cd /home/<your_username>/go/src/github.com/hyperledger/fabric/fabric-samples/test-network

export CC_PACKAGE_ID=<your_package_ID>

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name case  --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem;export CORE_PEER_TLS_ENABLED=true;export CORE_PEER_LOCALMSPID="Org1MSP";export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt;export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp;export CORE_PEER_ADDRESS=localhost:7051;source /etc/profile;peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name case --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem;peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name case --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json;peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name case --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt;peer lifecycle chaincode querycommitted --channelID mychannel --name case --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem;

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"initLedger","Args":[]}'

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"QueryAllCases","Args":[]}';


peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"CreateCase","Args":["evidence7","7-8a4a6d4ba723a6675c5c906c48aa532a","case7","evidence7","testpp","","","CHECKEDIN","7674526440","CgMBUWA="]}'


peer chaincode query -C mychannel -n case -c '{"function":"QueryCase","Args":["evidence7"]}'


###Checkin

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"Checkin","Args":["evidence7","7-88884ba723a6675c5c906c48aa532a","7888826440"]}'

###checkout
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"Checkout","Args":["evidence7","7-99994ba723a6675c5c906c48aa532a","7888826440"]}'

###Remove
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n case --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"Remove","Args":["evidence7","7-100000ba7675c5c906c48aa532a","100006440"]}'



docker logs <your_cli>
## Example: 
## docker logs orderer.example.com
## docker logs peer0.org1.example.com


peer channel getinfo -c mychannel

peer channel fetch newest mychannel.block -c mychannel;configtxgen --inspectBlock mychannel.block > mychannel.json

docker ps -a

docker exec -it <container_ID> /bin/sh

cd /var/hyperledger/production/ledgersData/chains
