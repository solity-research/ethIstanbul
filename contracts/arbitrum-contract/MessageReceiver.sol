// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../Hyperlane-Mailbox/IMailBox.sol";

contract MessageReceiver {

    IMailbox mailbox; 
    

    constructor() {
        //address of mailbox in  Arbitrum-goerli
        mailbox = IMailbox(0x13dABc0351407d5aAa0A50003a166A73b4febfDc);
    }

    event ReceivedMessage(uint32,address,uint, string);

    function bytes32ToAddress(bytes32 _buf) internal pure returns (address) {
        return address(uint160(uint256(_buf)));
    }

    function handle(
    uint32 _origin,
    bytes32 _sender,
    bytes calldata _message
) external payable {
        address sender = bytes32ToAddress(_sender);
       emit ReceivedMessage(_origin, sender, msg.value, string(_message));
}
   

}
