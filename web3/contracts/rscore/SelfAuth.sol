// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title SelfAuth
 * @dev SelfAuth supports real-time monitoring and private dynamic authorization services for on-chain operations
 * @dev The contract activates the private authorization by triggering the authRequest event, and implements the function authResponse. After the 
 * private authorization is successful, the call will be triggered
 * @author refitor
 */
interface SelfAuth {
    /**
     * @dev Trigger private authorization.
     * @param authAddr Account address that requires private authorization.
     * @param params In order to ensure the problem of intermediate state change caused by abnormal conditions, when triggering the event 
     * authRequest, please package the call parameters into bytes data for transmission, and receive and process by the function authResponse
     *
     * Emits a {authRequest} event.
     */
    event authRequest(address authAddr, bytes params);

    /**
     * @dev Receive private authorization response.
     * @param authAddr Account address that requires private authorization.
     * @param params In order to ensure the problem of intermediate state change caused by abnormal conditions, when triggering the event 
     * authRequest, please package the call parameters into bytes data for transmission, and receive and process by the function authResponse
     *
     * @dev encode example for params: abi.encode(1, 1)
     * @dev decode example for params: (uint256 methodID, uint256 value) = abi.decode(params, (uint256, uint256));
     *
     * WARNING: authResponse will only be called if private authorization is successful. It is necessary to perform caller verification in the 
     * implementation of this function. For details, please refer to the demo.
     */
    function authResponse(address authAddr, bytes memory params) external;
}