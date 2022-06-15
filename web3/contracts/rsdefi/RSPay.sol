// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './node_modules/@openzeppelin/contracts/utils/Address.sol';
import './node_modules/@openzeppelin/contracts/access/Ownable.sol';
import './node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol';
import 'https://github.com/refitor/selfAuth/blob/v1.0.0/web3/contracts/rscore/SelfAuth.sol';

// @title rs payment gateway inherited from OpenZeppelin Escrow
// @author refitor
contract RSPay is Ownable, SelfAuth {
    using SafeMath for uint256;
    using Address for address payable;

    event Deposited(address indexed payee, uint256 weiAmount);
    event Withdrawn(address indexed payee, uint256 weiAmount);

    mapping(address => uint256) private _deposits;

    /**
     * @dev Payment portal without lock-in period.
     */
    function deposit() external payable {
        uint256 amount = msg.value;
        _deposits[msg.sender] += amount;
        emit Deposited(msg.sender, amount);
    }

    /**
     * @dev Withdraw balance manually.
     * @param wallet The destination address of the funds.
     *
     * Emits a {authRequest} event.
     */
    function withdraw(address wallet) external onlyOwner {
        emit authRequest(wallet, "0x00");
    }

    /**
     * @dev Receive private authorization response.
     * @param authAddr Account address that requires private authorization.
     *
     * @dev encode example for params: abi.encode(1, 1)
     * @dev decode example for params: (uint256 methodID, uint256 value) = abi.decode(params, (uint256, uint256));
     *
     * WARNING: authResponse will only be called if private authorization is successful. It is necessary to perform caller verification in the 
     * implementation of this function. For details, please refer to the demo.
     */
    function authResponse(address authAddr, bytes memory /* params */) external {
        require(msg.sender == super.owner(), "permission denied");
        uint256 payment = address(this).balance;
        require(payment > 0, "insufficient balance");
        payable(authAddr).sendValue(payment);
        emit Withdrawn(authAddr, payment);
    }

    /**
     * @dev Check the balance of the specified address.
     * @param payee The destination address of the funds.
     */
    function depositsOf(address payee) public view returns (uint256) {
        return _deposits[payee];
    }
}