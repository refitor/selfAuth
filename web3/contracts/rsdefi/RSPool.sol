// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './node_modules/@openzeppelin/contracts/utils/escrow/Escrow.sol';
import './node_modules/@openzeppelin/contracts/access/Ownable.sol';
import './node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol';
import './SelfAuth.sol';

// @title rs payment gateway inherited from OpenZeppelin Escrow
// @author refitor
contract RSPool is Ownable, SelfAuth {
    Escrow escrow;
    using SafeMath for uint256;
    mapping (address => uint256) internal _lockedMap;

    constructor() {
        escrow = new Escrow();
    }

    /**
     * @dev Deposit funds for every address paid.
     */
    function deposit() external payable {
        require(msg.value > 0, "illegal operation");
        escrow.deposit{value: msg.value / 2}(msg.sender);
        _lockedMap[msg.sender] = block.timestamp + 365 days;
    }

    /**
     * @dev Withdraw funds for every address paid.
     */
    function withdraw() external payable {
        require(escrow.depositsOf(msg.sender) > 0, "permission denied");
        emit authRequest(msg.sender, abi.encode(1, escrow.depositsOf(msg.sender)));
    }

    /**
     * @dev Withdraw the specified amount to the wallet.
     * @param wallet The destination address of the funds.
     * @param amount The specified amount of the funds.
     *
     * Emits a {authRequest} event.
     */
    function withdrawOf(address payable wallet, uint256 amount) external onlyOwner {
        emit authRequest(wallet, abi.encode(2, amount));
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
        (uint256 mid, uint256 getValue) = abi.decode(params, (uint256, uint256));
        require(getValue > 0, "insufficient balance");
        address payable payaddr = payable(authAddr);
        if (mid == 1) {
            if(_lockedMap[authAddr] >= block.timestamp) {
                escrow.deposit{value: escrow.depositsOf(authAddr)}(payaddr);
            }
            escrow.withdraw(payaddr);
        } else if (mid == 2) {
            (bool success,) = payaddr.call{value: getValue}("");
            require(success, "Failed to send money");
            escrow.deposit{value: getValue}(payaddr);
        }
    }

    /**
     * @dev Check the balance of the specified address.
     * @param wallet The destination address of the funds.
     */
    function balanceOf(address wallet) external view onlyOwner returns (uint256) {
        return escrow.depositsOf(wallet);
    }
}