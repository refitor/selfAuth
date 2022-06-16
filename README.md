# selfAuth
SelfAuth supports real-time monitoring and private dynamic authorization services for on-chain operations

## Warning
The current project is only used for personal test environment verification and has not been verified by the Ethereum main network. At the same time, main.go only implements the basic authorization and verification process, and does not provide security protection for data storage and transmission. For the use of the Ethereum main network environment, please visit the official website. Website: https://selfauth.refitor.com

![/selfauth.png](/selfauth.png)

## Supported Features: Self-Host
1. go generate: Automatically bind the solidity contract to the go source code
2. contract deploy: Supports the deployment of the demo contract SADemo directly through the attached parameter --deploy=true
3. event monitor: Supports directional monitoring of EVM events, triggering authorization link sending and google authenticator dynamic verification
4. trigger authRequest: Support running demo example to trigger authRequest event, Each Trigger function call will fire the authRequest event
5. call authResponse: Each successful verification of the private authorization request will trigger the authResponse function call

## Setup
1. install go
2. install abigen
3. npm install openzeppelin

## Getting Start
1. git clone https://github.com/refitor/selfAuth.git
2. cd ./selfAuth && cp ./web3/contracts/rscore/SelfAuth.sol ./web3/contracts/demo/
3. go mod init selfAuth && go mod tidy
4. go generate
5. go build -i
6. ./selfAuth --deploy=true
7. ./selfAuth 
===> 1. Trigger authRequest, Each Trigger function call will fire the authRequest event
===> 2. Access the private authorization link received by the mailbox and perform dynamic verification through google authenticator
===> 3. Check verified count, Each successful verification of the private authorization request will trigger the authResponse function call, and increment the verified count in the contract SADemo
