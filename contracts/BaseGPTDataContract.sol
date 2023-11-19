// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BaseGPTDataContract {
    string public protocolGPTName;
    string public data;

    constructor(string memory _protocolGPTName, string memory _data) {
        protocolGPTName = _protocolGPTName;
        data = _data;
    }

}