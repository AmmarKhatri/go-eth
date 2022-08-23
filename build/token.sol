/**
 *Submitted for verification at Etherscan.io on 2017-11-28
*/

pragma solidity ^0.5.16;


contract ERC20Basic {
    event Transfer(address indexed from, address indexed to, uint value);
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
    event Approval(address indexed owner, address indexed spender, uint value);
}

