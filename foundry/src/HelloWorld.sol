// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract HelloWorld {
    string public greeting = "Hello, World!";

    uint public iteration = 0;

    /// @notice Sets a new greeting message and increments the iteration counter
    /// @param _greeting The new greeting message to set
    function setGreeting(string memory _greeting) public {
        iteration += 1;
        greeting = _greeting;
    }
}