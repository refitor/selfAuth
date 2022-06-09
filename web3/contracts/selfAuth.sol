//SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

interface ISelfAuth {

    // Trigger private authorization
    event authRequest(address authAddr, bytes params);

    // Receive private authorization response
    function authResponse(address authAddr, bytes memory params) external;
}

/*
contract SelfAuthDemo is ISelfAuth {
    address public SELFAUTH_OFFICIAL_ADDRESS=0x61e1CEB3E8bd42D2be9AD0d4fD58fc0e10C7bb5e;

    // authAddr: Ethereum account address for private authorization requests
    // params: In order to ensure the problem of intermediate state change caused by abnormal conditions, when triggering the event authRequest, 
    // please package the call parameters into bytes data for transmission, and receive and process by the function authResponse
    //
    // Notice: authResponse will only be called if private authorization is successful
    function authResponse(address authAddr, bytes memory params) external {
        require(msg.sender == SELFAUTH_OFFICIAL_ADDRESS, "permission denied");
    }
}
*/