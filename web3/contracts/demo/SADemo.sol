//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './node_modules/@openzeppelin/contracts/access/Ownable.sol';
import './node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol';
import 'https://github.com/refitor/selfAuth/blob/v1.0.0/web3/contracts/rscore/SelfAuth.sol';

// @title SelfAuth demo inherited from SelfAuth
// @author refitor
contract SADemo is Ownable, SelfAuth {
    using SafeMath for uint256;
    mapping (address => uint256) internal _verifyMap;

    /**
     * @dev Trigger private authorization.
     * @param authAddr The address of the verified.
     */
    function trigger(address authAddr) external onlyOwner {
        emit authRequest(authAddr, abi.encode(_verifyMap[authAddr]));
    }

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
    function authResponse(address authAddr, bytes memory params) external {
        require(msg.sender == super.owner(), "permission denied");
        uint256 verifiedCount = abi.decode(params, (uint256));
        _updateVerifyMap(authAddr, verifiedCount);
    }

    /**
     * @dev Verification count for each private authorized address.
     * @param authAddr The address of the verified.
     * @param authAddr The count verified.
     */
    function _updateVerifyMap(address authAddr, uint256 verifiedCount) internal {
        require(verifiedCount == _verifyMap[authAddr], "invalid operation");
        _verifyMap[authAddr] = verifiedCount + 1;
    }

    /**
     * @dev Get the number of private authorizations for the specified address.
     * @param authAddr The address of the verified.
     */
    function verified(address authAddr) public view onlyOwner returns (uint256) {
        return _verifyMap[authAddr];
    }
}
