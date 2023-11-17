// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../Hyperlane-Mailbox/IMailBox.sol";

contract MessageSender{

    event MessageSent(string message);

    IMailbox public mailbox; 

    constructor() {
        //address of mailbox in scroll sepolia
        mailbox = IMailbox(0x3C5154a193D6e2955650f9305c8d80c18C814A68);
    }


    fallback() external payable {
        //for fallback 
    }

    receive () external payable  {

    }

    function addressToBytes32(address _addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(_addr)));
    }

    // here we call dispatch function of mailbox
    function sendMessage(uint32 _destinationDomain, string memory _message, address _recipientAddress) public payable  {
        
        bytes memory messageBody = bytes(_message);
        bytes32 recipientAddress = addressToBytes32(_recipientAddress);
        mailbox.dispatch{value: msg.value} (_destinationDomain, recipientAddress, messageBody);
        emit MessageSent(_message);
    }

}
