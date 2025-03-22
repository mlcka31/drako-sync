// SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.19;

contract MoodTracker {
    mapping(address => string) private userMoods;

    event MoodSet(address user, string mood);

    function setMood(string memory _mood) public {
        userMoods[msg.sender] = _mood;
        emit MoodSet(msg.sender, _mood);
    }

    function getMood(address _user) public view returns (string memory) {
        return userMoods[_user];
    }

    function getMyMood() public view returns (string memory) {
        return userMoods[msg.sender];
    }
}
