// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./IMailBox.sol";
import "./BaseGPTDataContract.sol";

contract ProtocolGPTFactory {
    IMailbox mailbox;
    mapping(string => address) public createdContracts;

    event ContractCreated(string data, address newContract);
    event PromptSent(string message);

    constructor() {
        //address of mailbox in  scroll sepolia
        mailbox = IMailbox(0x3C5154a193D6e2955650f9305c8d80c18C814A68);
    }

    event ReceivedMessage(uint32, address, uint256, string);

    function bytes32ToAddress(bytes32 _buf) internal pure returns (address) {
        return address(uint160(uint256(_buf)));
    }

     function createNewContract(string memory data) public {
        // Parse 'data' to determine contract type and parameters
        (string memory gptName, string memory gptData) = splitString(
            data,
            "?"
        );

        BaseGPTDataContract newContract = new BaseGPTDataContract(
            gptName,
            gptData
        );
        createdContracts[gptName] = (address(newContract));
        emit ContractCreated(gptData, address(newContract));
    }

    function splitString(string memory _base, string memory _delimiter)
        public
        pure
        returns (string memory, string memory)
    {
        bytes memory baseBytes = bytes(_base);
        bytes memory delimiterBytes = bytes(_delimiter);

        uint256 splitIndex;
        for (uint256 i = 0; i < baseBytes.length; i++) {
            if (baseBytes[i] == delimiterBytes[0]) {
                splitIndex = i;
                break;
            }
        }

        bytes memory firstPart = new bytes(splitIndex);
        for (uint256 i = 0; i < splitIndex; i++) {
            firstPart[i] = baseBytes[i];
        }

        bytes memory secondPart = new bytes(baseBytes.length - splitIndex - 1);
        for (uint256 i = 0; i < secondPart.length; i++) {
            secondPart[i] = baseBytes[i + splitIndex + 1];
        }

        return (string(firstPart), string(secondPart));
    }

    function startsWith(string memory _base, string memory _value)
        public
        pure
        returns (bool)
    {
        bytes memory baseBytes = bytes(_base);
        bytes memory valueBytes = bytes(_value);

        if (baseBytes.length < valueBytes.length) {
            return false;
        }

        for (uint256 i = 0; i < valueBytes.length; i++) {
            if (baseBytes[i] != valueBytes[i]) {
                return false;
            }
        }
        return true;
    }

    function handle(
        uint32 _origin,
        bytes32 _sender,
        bytes calldata _message
    ) external payable {
        address sender = bytes32ToAddress(_sender);
        emit ReceivedMessage(_origin, sender, msg.value, string(_message));

        string memory message = string(_message);
        //  Check message prefix
        if (startsWith(message, "msg?")) {
            // Split the message and emit the relevant part
            (, string memory relevantMessage) = splitString(message, "?");
            emit PromptSent(relevantMessage);
        }else{
            createNewContract(message);

        }
    }

    function test() public {
        emit PromptSent("test evet");
    }
}
